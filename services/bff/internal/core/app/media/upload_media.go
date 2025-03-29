package media

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"shared/pb/mediapb"
)

type UploadMedia struct {
	mediaClient mediapb.MediaServiceClient
}

func NewUploadMedia(mediaClient mediapb.MediaServiceClient) *UploadMedia {
	return &UploadMedia{
		mediaClient: mediaClient,
	}
}

type UploadMediaRequest struct {
	Media *multipart.FileHeader `form:"media"`
	Title string                `form:"title"`
}

type UploadMediaResponse struct {
	MediaID string `json:"media_id"`
}

func (h *UploadMedia) Handle(ctx context.Context, req *UploadMediaRequest) (*UploadMediaResponse, int, error) {
	err := h.validateForm(req)
	if err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("validateForm: %w", err)
	}

	stream, err := h.mediaClient.UploadMedia(ctx)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("mediaClient.UploadMedia: %w", err)
	}
	defer stream.CloseSend()

	err = stream.Send(&mediapb.UploadMediaRequest{
		Title: req.Title,
	})
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("mediaClient.UploadMedia: %w", err)
	}

	file, err := req.Media.Open()
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("file.Open: %w", err)
	}
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
		if err != nil {
			return nil, http.StatusInternalServerError, fmt.Errorf("mediaClient.UploadMedia: %w", err)
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("mediaClient.UploadMedia: %w", err)
	}

	return &UploadMediaResponse{
		MediaID: resp.MediaId,
	}, http.StatusOK, nil
}

func (h *UploadMedia) validateForm(req *UploadMediaRequest) error {
	if req.Title == "" {
		return fmt.Errorf("title is required")
	}
	if req.Media == nil {
		return fmt.Errorf("media is required")
	}
	return nil
}
