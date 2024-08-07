package port

import (
	"context"

	"github.com/yuorei/video-server/app/domain"
)

// adaputerがusecase層を呼び出されるメソッドのインターフェースを定義
type VideoInputPort interface {
	GetVideos(context.Context) ([]*domain.Video, error)
	GetVideosByUserID(context.Context, string) ([]*domain.Video, error)
	GetVideo(context.Context, string) (*domain.Video, error)
	UploadVideo(context.Context, *domain.UploadVideo) (*domain.UploadVideoResponse, error)
	GetWatchCount(context.Context, string) (int, error)
	IncrementWatchCount(context.Context, string, string) (int, error)
	CutVideo(context.Context, string, int, int) (string, error)
}

// ユースケースからインフラを呼び出されるメソッドのインターフェースを定義
type VideoRepository interface {
	GetVideosFromDB(context.Context) ([]*domain.Video, error)
	GetVideosByUserIDFromDB(context.Context, string) ([]*domain.Video, error)
	UploadVideoForStorage(context.Context, *domain.UploadVideo, string) (string, error)
	GetVideoFromDB(context.Context, string) (*domain.Video, error)
	GetWatchCount(context.Context, string) (int, error)
	IncrementWatchCount(context.Context, string, string) (int, error)
	CutVideo(context.Context, string, string, int, int) (string, error)
}
