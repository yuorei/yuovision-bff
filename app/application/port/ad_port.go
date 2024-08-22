package port

import (
	"context"

	"github.com/yuorei/video-server/app/domain"
)

// adaputerがusecase層を呼び出されるメソッドのインターフェースを定義
type AdInputPort interface {
	GetAdsByVideoID(context.Context, *domain.GetAdVideoRequest) ([]*domain.Ad, error)
	WatchCountAdVideo(context.Context, *domain.WatchCountAdVideoRequest) error
}

// ユースケースからインフラを呼び出されるメソッドのインターフェースを定義
type AdRepository interface {
	GetAdsByVideoIDFromAdsServer(context.Context, *domain.GetAdVideoRequest) ([]*domain.Ad, error)
	WatchCountAdVideoFromAdsServer(context.Context, *domain.WatchCountAdVideoRequest) error
}
