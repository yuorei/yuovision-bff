// ImageRepository
package port

import (
	"context"

	"github.com/yuorei/video-server/app/domain"
)

// adaputerがusecase層を呼び出されるメソッドのインターフェースを定義
type ImageInputPort interface {
}

// ユースケースからインフラを呼び出されるメソッドのインターフェースを定義
type ImageRepository interface {
	UploadThumbnailForStorage(context.Context, domain.ThumbnailImage) error
}
