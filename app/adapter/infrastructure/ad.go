package infrastructure

import (
	"context"

	"connectrpc.com/connect"
	"github.com/yuorei/video-server/app/domain"
	adsv1 "github.com/yuorei/yuorei-ads-proto/gen/rpc/ads/v1"
)

func (i *Infrastructure) GetAdsByVideoIDFromAdsServer(ctx context.Context, req *domain.GetAdVideoRequest) ([]*domain.Ad, error) {
	var description string
	if req.Description == nil {
		description = ""
	} else {
		description = *req.Description
	}

	ad, err := i.adsServer.ClientAds.GetAdVideo(ctx,
		connect.NewRequest(&adsv1.GetAdVideoRequest{
			// ブラウザ情報
			UserAgent:            req.UserAgent,
			Platform:             req.Platform,
			Language:             req.Language,
			Url:                  req.URL,
			PageTitle:            req.PageTitle,
			Referrer:             req.Referrer,
			NetworkDownlink:      req.NetworkDownlink,
			NetworkEffectiveType: req.NetworkEffectiveType,
			IpAddress:            req.IPAddress,
			Location:             req.Location,
			Hostname:             req.Hostname,
			City:                 req.City,
			Region:               req.Region,
			Country:              req.Country,
			Org:                  req.Org,
			Postal:               req.Postal,
			Timezone:             req.Timezone,
			// ビデオ情報
			VideoId:          req.VideoID,
			VideoTitle:       req.Title,
			VideoDescription: description,
			VideoTags:        req.Tags,
			// ユーザー情報
			UserId:   req.UserID,
			ClientId: req.ClientID,
		}))
	if err != nil {
		return nil, err
	}

	adsResponse := make([]*domain.Ad, 0)
	for _, ad := range ad.Msg.Responses {
		adResponse := domain.NewAd(ad.AdId, ad.AdUrl, ad.Title, ad.Description, ad.ThumbnailUrl, ad.VideoUrl)
		adsResponse = append(adsResponse, adResponse)
	}

	return adsResponse, nil
}

func (i *Infrastructure) WatchCountAdVideoFromAdsServer(ctx context.Context, postAd *domain.Ad) (*domain.Ad, error) {
	return nil, nil
}
