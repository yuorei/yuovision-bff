package application

import (
	"github.com/yuorei/video-server/app/adapter/infrastructure"
)

type Application struct {
	Video   *VideoUseCase
	Image   *ImageUseCase
	User    *UserUseCase
	Comment *CommentUseCase
	Ad      *AdUseCase
}

func NewApplication(infra *infrastructure.Infrastructure) *Application {
	videoUseCase := NewVideoUseCase(infra)
	imageUseCase := NewImageUseCase(infra)
	userUseCase := NewUserUseCase(infra)
	CommentUseCase := NewCommentUseCase(infra)
	AdUseCase := NewAdUseCase(infra)

	return &Application{
		Video:   videoUseCase,
		Image:   imageUseCase,
		User:    userUseCase,
		Comment: CommentUseCase,
		Ad:      AdUseCase,
	}
}
