package application

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/yuorei/video-server/app/application/port"
	"github.com/yuorei/video-server/app/domain"
	"github.com/yuorei/video-server/middleware"
)

type VideoUseCase struct {
	videoRepository port.VideoRepository
}

func NewVideoUseCase(videoRepository port.VideoRepository) *VideoUseCase {
	return &VideoUseCase{
		videoRepository: videoRepository,
	}
}

func (a *Application) GetVideos(ctx context.Context) ([]*domain.Video, error) {
	videos, err := a.Video.videoRepository.GetVideosFromDB(ctx)
	if err != nil {
		return nil, err
	}

	sort.Slice(videos, func(i, j int) bool {
		return videos[j].CreatedAt.Before(videos[i].CreatedAt)
	})

	return videos, nil
}

func (a *Application) GetVideosByUserID(ctx context.Context, userID string) ([]*domain.Video, error) {
	videos, err := a.Video.videoRepository.GetVideosByUserIDFromDB(ctx, userID)
	if err != nil {
		return nil, err
	}

	sort.Slice(videos, func(i, j int) bool {
		return videos[j].CreatedAt.Before(videos[i].CreatedAt)
	})

	return videos, nil
}

func (a *Application) GetVideo(ctx context.Context, videoID string) (*domain.Video, error) {
	return a.Video.videoRepository.GetVideoFromDB(ctx, videoID)
}

func (a *Application) UploadVideo(ctx context.Context, video *domain.UploadVideo) (*domain.UploadVideoResponse, error) {
	userID, err := middleware.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	videoURL, err := a.Video.videoRepository.UploadVideoForStorage(ctx, video, userID)
	if err != nil {
		return nil, err
	}

	var data []byte
	if video.ThumbnailImage != nil {
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(*video.ThumbnailImage)
		if err != nil {
			return nil, err
		}
		data = buf.Bytes()
	}

	contentType := video.ImageContentType
	thumbnail := domain.NewThumbnailImage(video.ID, contentType, data)
	err = a.Image.imageRepository.UploadThumbnailForStorage(ctx, *thumbnail)
	if err != nil {
		return nil, err
	}

	imagePath := video.ID + ".webp"
	const buckerName = "thumbnail-image"
	imageURL := fmt.Sprintf("%s/%s/%s", os.Getenv("AWS_S3_URL"), buckerName, imagePath)
	videoResponse := &domain.UploadVideoResponse{
		ID:                video.ID,
		VideoURL:          videoURL,
		ThumbnailImageURL: imageURL,
		Title:             video.Title,
		Description:       video.Description,
		IsPrivate:         video.IsPrivate,
		IsAdult:           video.IsAdult,
		IsExternalCutout:  video.IsExternalCutout,
		IsAd:              video.IsAd,
		UploaderID:        userID,
	}

	return videoResponse, nil
}

func (a *Application) GetWatchCount(ctx context.Context, videoID string) (int, error) {
	return a.Video.videoRepository.GetWatchCount(ctx, videoID)
}

func (a *Application) IncrementWatchCount(ctx context.Context, videoID, userID string) (int, error) {
	if !strings.Contains(userID, "client") {
		return 0, fmt.Errorf("not client")
	}

	return a.Video.videoRepository.IncrementWatchCount(ctx, videoID, userID)
}

func (a *Application) CutVideo(ctx context.Context, videoID string, start, end int) (string, error) {
	userID, err := middleware.GetUserIDFromContext(ctx)
	if err != nil {
		return "", err
	}

	return a.Video.videoRepository.CutVideo(ctx, videoID, userID, start, end)
}
