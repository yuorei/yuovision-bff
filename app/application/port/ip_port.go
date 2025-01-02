package port

import (
	"context"

	"github.com/yuorei/video-server/app/domain"
)

type IPInputPort interface {
	IPInfomation(ctx context.Context, ip string) (*domain.IPResponse, error)
}

type IPRepository interface {
	GetIPInfomation(ctx context.Context, ip string) (*domain.IPResponse, error)
}
