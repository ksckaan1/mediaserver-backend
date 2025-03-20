package app

import (
	"common/pb/mediapb"
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestABC(t *testing.T) {
	client, err := grpc.NewClient("localhost:9191", grpc.WithInsecure())
	require.NoError(t, err)

	mediaClient := mediapb.NewMediaServiceClient(client)

	ctx := context.Background()

	stream, err := mediaClient.UploadMedia(ctx)
	require.NoError(t, err)

	defer func() {
		msg, err := stream.CloseAndRecv()
		require.NoError(t, err)
		t.Log(msg)
	}()

	defer stream.CloseSend()

	err = stream.Send(&mediapb.UploadMediaRequest{
		Title: "gopher.png",
	})
	require.NoError(t, err)

	file, err := os.Open("/Users/ksckaan1/Pictures/gopher.png")
	require.NoError(t, err)
	defer file.Close()

	buff := make([]byte, 1024)
	for {
		n, err := file.Read(buff)
		if err != nil {
			break
		}
		err = stream.Send(&mediapb.UploadMediaRequest{
			Content: buff[:n],
		})
		require.NoError(t, err)
		t.Log("ok")
	}

	t.Log("bitti")
}
