package client

import (
	"net/http"

	adsv1 "github.com/yuorei/yuorei-ads-proto/gen/rpc/ads/v1/adsv1connect"
)

type AdsServer struct {
	ClientAds adsv1.AdManagementServiceClient
}

func NewClientAdServer() *AdsServer {
	clientAds := adsv1.NewAdManagementServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
	)

	return &AdsServer{
		ClientAds: clientAds,
	}
}
