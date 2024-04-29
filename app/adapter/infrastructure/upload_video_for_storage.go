package infrastructure

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/yuorei/video-server/app/domain"
	"github.com/yuorei/video-server/lib"
	"github.com/yuorei/video-server/yuovision-proto/go/video/video_grpc"
)

func (i *Infrastructure) UploadVideoForStorage(ctx context.Context, video *domain.UploadVideo, userID string) (string, error) {
	stream, err := i.gRPCClient.VideoClient.UploadVideo(ctx)
	if err != nil {
		return "", err
	}

	meta := &video_grpc.UploadVideoInput_Meta{
		Meta: &video_grpc.VideoMeta{
			Id:                video.ID,
			Title:             video.Title,
			Description:       *video.Description,
			UserId:            userID,
			ThumbnailImageUrl: fmt.Sprintf("%s/%s/%s.webp", os.Getenv("AWS_S3_URL"), "thumbnail-image", video.ID),
		},
	}

	request := &video_grpc.UploadVideoInput{
		Value: meta,
	}

	err = stream.Send(request)
	if err != nil {
		return "", err
	}

	data, err := lib.ReadSeekerToBytes(video.Video)
	if err != nil {
		return "", err
	}
	chunkSize := 3 * 1024 * 1024 // チャンクサイズ（3MB）
	for offset := 0; offset < len(data); offset += chunkSize {
		end := offset + chunkSize
		if end > len(data) {
			end = len(data)
		}
		chunk := data[offset:end]

		// Create a request containing the chunk of thumbnail data
		request := &video_grpc.UploadVideoInput{
			Value: &video_grpc.UploadVideoInput_Video{
				Video: chunk,
			},
		}
		// Send the chunk data
		err := stream.Send(request)
		if err != nil {
			return "", err
		}
	}

	// Receive response from the server
	videoPayload, err := stream.CloseAndRecv()
	if err != nil {
		return "", err
	}

	// bucketName := "video"
	// url := fmt.Sprintf("%s/%s/output_%s.m3u8", os.Getenv("AWS_S3_URL"), bucketName, video.ID)
	return videoPayload.VideoUrl, nil
}

func uploadVideoForS3(path string) error {
	ctx := context.Background()
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	cred := credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")
	cfg, err := config.LoadDefaultConfig(ctx, config.WithCredentialsProvider(cred))
	if err != nil {
		return err
	}

	// change object address style
	client := s3.NewFromConfig(cfg, func(options *s3.Options) {
		options.UsePathStyle = true
		options.BaseEndpoint = aws.String(os.Getenv("AWS_S3_ENDPOINT"))
		options.Region = "ap-northeast-1"
	})

	// get buckets
	lbo, err := client.ListBuckets(ctx, nil)
	if err != nil {
		return err
	}
	buckets := make(map[string]struct{}, len(lbo.Buckets))
	for _, b := range lbo.Buckets {
		buckets[*b.Name] = struct{}{}
	}

	// create 'video' bucket if not exist
	bucketName := "video"
	if _, ok := buckets[bucketName]; !ok {
		_, err = client.CreateBucket(ctx, &s3.CreateBucketInput{
			Bucket: &bucketName,
			ACL:    types.BucketCannedACLPublicRead,
		})
		if err != nil {
			return err
		}
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	// put object
	_, err = client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(strings.Split(path, "/")[1]),
		Body:   file,
		ACL:    types.ObjectCannedACLPublicRead,
	})

	if err != nil {
		return fmt.Errorf("failed to upload file: %w", err)
	}
	log.Println("Successful upload: ", path)

	return nil
}
