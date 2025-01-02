package application

import (
	"context"

	"github.com/yuorei/video-server/app/application/port"
	"github.com/yuorei/video-server/app/domain"
)

type IPUseCase struct {
	userRepository port.IPRepository
}

func NewIPUseCase(userRepository port.IPRepository) *IPUseCase {
	return &IPUseCase{
		userRepository: userRepository,
	}
}

func (a *Application) IPInfomation(ctx context.Context, id string) (*domain.IPResponse, error) {
	return a.IP.userRepository.GetIPInfomation(ctx, id)
}
