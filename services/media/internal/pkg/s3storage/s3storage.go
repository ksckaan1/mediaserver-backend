package s3storage

import (
	"context"
	"fmt"
	"media_service/config"
	"media_service/internal/port"
	"shared/ports"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var _ port.Storage = (*S3Stroage)(nil)

type S3Stroage struct {
	minioClient *minio.Client
	bucketName  string
	idGenerator ports.IDGenerator
}

func New(cfg *config.Config, idGenerator ports.IDGenerator) (*S3Stroage, error) {
	minioClient, err := minio.New(cfg.S3Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.S3AccessKey, cfg.S3SecretKey, ""),
		Secure: cfg.S3UseSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("minio.New: %w", err)
	}
	return &S3Stroage{
		minioClient: minioClient,
		bucketName:  cfg.S3Bucket,
		idGenerator: idGenerator,
	}, nil
}

func (s *S3Stroage) Save(ctx context.Context, object *port.Object) (string, error) {
	id := s.idGenerator.NewID()
	filePath := fmt.Sprintf("%s.%s", id, object.Extension)
	_, err := s.minioClient.PutObject(ctx, s.bucketName, filePath, object.Content, object.Size, minio.PutObjectOptions{
		ContentType: object.MimeType,
	})
	if err != nil {
		return "", fmt.Errorf("minio.PutObject: %w", err)
	}
	return filePath, nil
}

func (s *S3Stroage) Delete(ctx context.Context, filePath string) error {
	err := s.minioClient.RemoveObject(ctx, s.bucketName, filePath, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("minio.RemoveObject: %w", err)
	}
	return nil
}

func (s *S3Stroage) CreatePresignedURL(ctx context.Context, id string) (string, map[string]string, error) {
	policy := minio.NewPostPolicy()
	policy.SetBucket(s.bucketName)
	policy.SetKey(fmt.Sprintf("%s-${filename}", id))
	policy.SetExpires(time.Now().Add(1 * time.Hour))
	policy.SetContentTypeStartsWith("video/")
	policy.SetContentLengthRange(0, 1024*1024*1024) // 1GB
	presignedURL, formData, err := s.minioClient.PresignedPostPolicy(ctx, policy)
	if err != nil {
		return "", nil, fmt.Errorf("minio.PresignedPostPolicy: %w", err)
	}
	fmt.Printf("curl ")
	for k, v := range formData {
		fmt.Printf("-F %s=%s ", k, v)
	}
	fmt.Printf("-F file=@/Users/ksckaan1/Movies/big-buck-bunny.mp4 ")
	fmt.Printf("%s\n", presignedURL.String())
	return presignedURL.String(), formData, nil
}
