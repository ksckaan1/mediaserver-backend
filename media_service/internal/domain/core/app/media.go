package app

import (
	"common/enums/mediatype"
	"common/pb/mediapb"
	"context"
	"fmt"
	"io"
	"media_service/internal/domain/core/model"
	"media_service/internal/port"
	"os"

	"github.com/h2non/filetype"
	"github.com/samber/lo"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ mediapb.MediaServiceServer = (*Media)(nil)

type Media struct {
	mediapb.UnimplementedMediaServiceServer
	repo    Repository
	storage port.Storage
	idGen   port.IDGenerator
}

func New(repo Repository, storage port.Storage, idGen port.IDGenerator) (*Media, error) {
	return &Media{
		repo:    repo,
		storage: storage,
		idGen:   idGen,
	}, nil
}

func (m *Media) UploadMedia(stream grpc.ClientStreamingServer[mediapb.UploadMediaRequest, mediapb.UploadMediaResponse]) error {
	resp, err := stream.Recv()
	if err != nil {
		return fmt.Errorf("stream.Recv: %w", err)
	}
	object, err := os.CreateTemp("", "media_file")
	if err != nil {
		return fmt.Errorf("os.CreateTemp: %w", err)
	}
	defer os.Remove(object.Name())
	defer object.Close()
	fileSize := 0

	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			break
		}
		n, err := object.Write(chunk.Content)
		if err != nil {
			return fmt.Errorf("object.Write: %w", err)
		}
		fileSize += n
	}

	id := m.idGen.NewID()
	ctx := stream.Context()

	_, err = object.Seek(0, 0)
	if err != nil {
		return fmt.Errorf("object.Seek: %w", err)
	}

	fileType, err := filetype.MatchReader(object)
	if err != nil {
		return fmt.Errorf("filetype.MatchReader: %w", err)
	}

	_, err = object.Seek(0, 0)
	if err != nil {
		return fmt.Errorf("object.Seek: %w", err)
	}

	filePath, err := m.storage.Save(ctx, &port.Object{
		Content:   object,
		Size:      int64(fileSize),
		MimeType:  fileType.MIME.Value,
		Extension: fileType.Extension,
	})
	if err != nil {
		return fmt.Errorf("storage.Save: %w", err)
	}

	err = m.repo.CreateMedia(ctx, &model.Media{
		ID:       id,
		Title:    resp.Title,
		Path:     filePath,
		Type:     mediatype.Image,
		MimeType: fileType.MIME.Value,
		Size:     int64(fileSize),
	})
	if err != nil {
		return fmt.Errorf("repo.CreateMedia: %w", err)
	}

	return stream.SendAndClose(&mediapb.UploadMediaResponse{
		MediaId: id,
	})
}

func (m *Media) GetMediaByID(ctx context.Context, req *mediapb.GetMediaByIDRequest) (*mediapb.Media, error) {
	media, err := m.repo.GetMediaByID(ctx, req.GetMediaId())
	if err != nil {
		return nil, fmt.Errorf("repo.GetMediaByID: %w", err)
	}
	return &mediapb.Media{
		Id:        media.ID,
		CreatedAt: timestamppb.New(media.CreatedAt),
		UpdatedAt: timestamppb.New(media.UpdatedAt),
		Title:     media.Title,
		Path:      media.Path,
		Type:      mediapb.MediaType(media.Type.Number()),
		MimeType:  media.MimeType,
		Size:      media.Size,
	}, nil
}

func (m *Media) ListMedias(ctx context.Context, req *mediapb.ListMediasRequest) (*mediapb.MediaList, error) {
	medias, err := m.repo.ListMedias(ctx, req.Limit, req.Offset)
	if err != nil {
		return nil, fmt.Errorf("repo.ListMedias: %w", err)
	}
	return &mediapb.MediaList{
		List: lo.Map(medias.List, func(m *model.Media, _ int) *mediapb.Media {
			return &mediapb.Media{
				Id:        m.ID,
				CreatedAt: timestamppb.New(m.CreatedAt),
				UpdatedAt: timestamppb.New(m.UpdatedAt),
				Title:     m.Title,
				Path:      m.Path,
				Type:      mediapb.MediaType(m.Type.Number()),
				MimeType:  m.MimeType,
				Size:      m.Size,
			}
		}),
		Count:  medias.Count,
		Limit:  medias.Limit,
		Offset: medias.Offset,
	}, nil
}

func (m *Media) UpdateMediaByID(ctx context.Context, req *mediapb.UpdateMediaByIDRequest) (*emptypb.Empty, error) {
	err := m.repo.UpdateMediaByID(ctx, &model.Media{
		ID:    req.MediaId,
		Title: req.Title,
	})
	if err != nil {
		return nil, fmt.Errorf("repo.UpdateMediaByID: %w", err)
	}
	return nil, nil
}

func (m *Media) DeleteMediaByID(ctx context.Context, req *mediapb.DeleteMediaByIDRequest) (*emptypb.Empty, error) {
	media, err := m.repo.GetMediaByID(ctx, req.MediaId)
	if err != nil {
		return nil, fmt.Errorf("repo.GetMediaByID: %w", err)
	}
	err = m.repo.DeleteMediaByID(ctx, req.MediaId)
	if err != nil {
		return nil, fmt.Errorf("repo.DeleteMediaByID: %w", err)
	}
	err = m.storage.Delete(ctx, media.Path)
	if err != nil {
		return nil, fmt.Errorf("storage.Delete: %w", err)
	}
	return nil, nil
}
