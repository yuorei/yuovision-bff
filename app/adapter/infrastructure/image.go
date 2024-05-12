package infrastructure

import (
	"context"

	"github.com/yuorei/video-server/app/domain"
	"github.com/yuorei/video-server/yuovision-proto/go/video/video_grpc"
)

func (i *Infrastructure) UploadThumbnailForStorage(ctx context.Context, thumbnail domain.ThumbnailImage) error {
	stream, err := i.gRPCClient.VideoClient.UploadThumbnail(ctx)
	if err != nil {
		return err
	}

	meta := &video_grpc.UploadThumbnailInput_Meta{
		Meta: &video_grpc.ThumbnailMeta{
			Id:          thumbnail.ID,
			ContentType: thumbnail.ContentType,
		},
	}

	request := &video_grpc.UploadThumbnailInput{
		Value: meta,
	}

	err = stream.Send(request)
	if err != nil {
		return err
	}

	data := thumbnail.ThumbnailImage
	chunkSize := 3 * 1024 * 1024 // チャンクサイズ（3MB）
	for offset := 0; offset < len(data); offset += chunkSize {
		end := offset + chunkSize
		if end > len(data) {
			end = len(data)
		}
		chunk := data[offset:end]

		// Create a request containing the chunk of thumbnail data
		request := &video_grpc.UploadThumbnailInput{
			Value: &video_grpc.UploadThumbnailInput_ThumbnailImage{
				ThumbnailImage: chunk,
			},
		}
		// Send the chunk data
		err := stream.Send(request)
		if err != nil {
			return err
		}
	}

	// Receive response from the server
	_, err = stream.CloseAndRecv()
	if err != nil {
		return err
	}

	return nil
}
