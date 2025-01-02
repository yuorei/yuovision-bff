package application

import (
	"context"

	"math/rand"

	"github.com/yuorei/video-server/app/application/port"
	"github.com/yuorei/video-server/app/domain"
)

type AdUseCase struct {
	adRepository port.AdRepository
}

func NewAdUseCase(adRepository port.AdRepository) *AdUseCase {
	return &AdUseCase{
		adRepository: adRepository,
	}
}

func (a *Application) GetAdsByVideoID(ctx context.Context, req *domain.GetAdVideoRequest) ([]*domain.Ad, error) {
	ads, err := a.Ad.adRepository.GetAdsByVideoIDFromAdsServer(ctx, req)
	if err != nil {
		return nil, err
	}

	// ランダムにシャッフル
	rand.Shuffle(len(ads), func(i, j int) { ads[i], ads[j] = ads[j], ads[i] })
	return ads, nil
}

func (a *Application) WatchCountAdVideo(ctx context.Context, req *domain.WatchCountAdVideoRequest) error {
	return a.Ad.adRepository.WatchCountAdVideoFromAdsServer(ctx, req)
}
