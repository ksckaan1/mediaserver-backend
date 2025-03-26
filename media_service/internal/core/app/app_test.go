package app

import (
	"common/pb/mediapb"
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestABC(t *testing.T) {
	client, err := grpc.NewClient("localhost:9191", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
		Title: "00e915c5e8b481c31b47e6a53c01251a.jpg",
	})
	require.NoError(t, err)

	// "C:\Users\kubra\OneDrive\Resimler\redmi12c\00e915c5e8b481c31b47e6a53c01251a.jpg"
	file, err := os.Open("C:\\Users\\kubra\\OneDrive\\Resimler\\redmi12c\\00e915c5e8b481c31b47e6a53c01251a.jpg")
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
