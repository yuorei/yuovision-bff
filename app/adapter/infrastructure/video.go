package infrastructure

import (
	"context"
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/yuorei/video-server/app/domain"
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
		videos = append(videos, domain.NewVideo(video.Id, video.VideoUrl, video.ThumbnailImageUrl, video.Title, &video.Description, video.UserId, video.CreatedAt.AsTime()))
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
		videos = append(videos, domain.NewVideo(video.Id, video.VideoUrl, video.ThumbnailImageUrl, video.Title, &video.Description, video.UserId, video.CreatedAt.AsTime()))
	}
	return videos, nil
}

func (i *Infrastructure) GetVideoFromDB(ctx context.Context, id string) (*domain.Video, error) {
	videoPayload, err := i.gRPCClient.VideoClient.Video(ctx, &video_grpc.VideoID{Id: id})
	if err != nil {
		return nil, err
	}

	video := domain.NewVideo(videoPayload.Id, videoPayload.VideoUrl, videoPayload.ThumbnailImageUrl, videoPayload.Title, &videoPayload.Description, videoPayload.UserId, videoPayload.CreatedAt.AsTime())
	return video, nil
}
