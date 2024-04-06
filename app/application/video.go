package application

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"sort"

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
	id, err := middleware.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	videofile := domain.NewVideoFile(video.ID, video.Video)
	err = a.Video.videoRepository.ConvertVideoHLS(ctx, videofile)
	if err != nil {
		return nil, err
	}

	videoURL, err := a.Video.videoRepository.UploadVideoForStorage(ctx, videofile)
	if err != nil {
		return nil, err
	}

	imageBuffer, err := a.Image.imageRepository.ConvertThumbnailToWebp(ctx, video.ThumbnailImage, video.ImageContentType, video.ID)
	if err != nil {
		return nil, err
	}

	if imageBuffer == nil {
		err = a.Image.imageRepository.CreateThumbnail(ctx, video.ID, video.Video)
		if err != nil {
			return nil, err
		}
	}
	imagePath := video.ID + ".webp"
	image, err := os.Open(imagePath)
	if err != nil {
		return nil, err
	}
	defer func() error {
		err := os.Remove(imagePath)
		if err != nil {
			return err
		}
		return nil
	}()
	data, err := ioutil.ReadAll(image)
	if err != nil {
		return nil, err
	}

	thumbnail := domain.NewThumbnailImage(video.ID, data)
	err = a.Image.imageRepository.UploadThumbnailForStorage(ctx, *thumbnail)
	if err != nil {
		return nil, err
	}

	imageURL := fmt.Sprintf("%s/%s/%s", os.Getenv("AWS_S3_URL"), "thumbnail-image", imagePath)
	videoResponse, err := a.Video.videoRepository.InsertVideo(ctx, video.ID, videoURL, imageURL, video.Title, video.Description, id)
	if err != nil {
		return nil, err
	}

	return videoResponse, nil
}
