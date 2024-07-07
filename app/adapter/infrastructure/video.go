package infrastructure

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/yuorei/video-server/app/domain"
	"github.com/yuorei/video-server/lib"
	"github.com/yuorei/video-server/yuovision-proto/go/video/video_grpc"
)

func (i *Infrastructure) GetVideosFromDB(ctx context.Context) ([]*domain.Video, error) {
	var videos []*domain.Video
	key := fmt.Sprintf("%T", videos)
	hit, err := getFromRedis(ctx, i.redis, key, &videos)
	if err != nil {
		return nil, err
	}
	if hit {
		return videos, nil
	}

	videosResponse, err := i.gRPCClient.VideoClient.Videos(ctx, &empty.Empty{})
	if err != nil {
		return nil, err
	}
	videos = make([]*domain.Video, 0, len(videosResponse.Videos))
	for _, video := range videosResponse.Videos {
		videos = append(videos, domain.NewVideo(video.Id, video.VideoUrl, video.ThumbnailImageUrl, video.Title, &video.Description, video.Tags, video.Private, video.Adult, video.ExternalCutout, video.IsAd, video.UserId, video.CreatedAt.AsTime(), video.UpdatedAt.AsTime()))
	}
	err = setToRedis(ctx, i.redis, key, 1*time.Minute, videos)
	if err != nil {
		return nil, err
	}

	return videos, nil
}

func (i *Infrastructure) GetVideosByUserIDFromDB(ctx context.Context, userID string) ([]*domain.Video, error) {
	videoResponse, err := i.gRPCClient.VideoClient.VideosByUserID(ctx, &video_grpc.VideoUserID{Id: userID})
	if err != nil {
		return nil, err
	}
	videos := make([]*domain.Video, 0, len(videoResponse.Videos))
	for _, video := range videoResponse.Videos {
		videos = append(videos, domain.NewVideo(video.Id, video.VideoUrl, video.ThumbnailImageUrl, video.Title, &video.Description, video.Tags, video.Private, video.Adult, video.ExternalCutout, video.IsAd, video.UserId, video.CreatedAt.AsTime(), video.UpdatedAt.AsTime()))
	}
	return videos, nil
}

func (i *Infrastructure) GetVideoFromDB(ctx context.Context, id string) (*domain.Video, error) {
	videoPayload, err := i.gRPCClient.VideoClient.Video(ctx, &video_grpc.VideoID{Id: id})
	if err != nil {
		return nil, err
	}

	video := domain.NewVideo(videoPayload.Id, videoPayload.VideoUrl, videoPayload.ThumbnailImageUrl, videoPayload.Title, &videoPayload.Description, videoPayload.Tags, videoPayload.Private, videoPayload.Adult, videoPayload.ExternalCutout, videoPayload.IsAd, videoPayload.UserId, videoPayload.CreatedAt.AsTime(), videoPayload.UpdatedAt.AsTime())
	return video, nil
}

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
			ThumbnailImageUrl: fmt.Sprintf("%s/%s/%s.webp", os.Getenv("AWS_S3_URL"), "thumbnail-image", video.ID),
			UserId:            userID,
			Tags:              video.Tags,
			Private:           video.IsPrivate,
			Adult:             video.IsAdult,
			ExternalCutout:    video.IsExternalCutout,
			IsAd:              video.IsAd,
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

	return videoPayload.VideoUrl, nil
}

func (i *Infrastructure) GetWatchCount(ctx context.Context, videoID string) (int, error) {
	videoWatchCount, err := i.gRPCClient.VideoClient.WatchCount(ctx, &video_grpc.WatchCountInput{VideoId: videoID})
	if err != nil {
		return 0, err
	}

	return int(videoWatchCount.Count), nil
}

func (i *Infrastructure) IncrementWatchCount(ctx context.Context, videoID, userID string) (int, error) {
	videoWatchCount, err := i.gRPCClient.VideoClient.IncrementWatchCount(ctx, &video_grpc.IncrementWatchCountInput{VideoId: videoID, UserId: userID})
	if err != nil {
		return 0, err
	}

	return int(videoWatchCount.Count), nil
}

func (i *Infrastructure) CutVideo(ctx context.Context, videoID, userID string, start, end int) (string, error) {
	cutVideoPayload, err := i.gRPCClient.VideoClient.CutVideo(ctx, &video_grpc.CutVideoInput{
		VideoId: videoID,
		UserId:  userID,
		Start:   int32(start),
		End:     int32(end),
	})
	if err != nil {
		return "", err
	}

	return cutVideoPayload.VideoUrl, nil
}
