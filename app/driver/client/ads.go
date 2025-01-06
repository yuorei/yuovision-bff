package client

import (
	"net/http"
	"os"

	adsv1 "github.com/yuorei/yuorei-ads-proto/gen/rpc/ads/v1/adsv1connect"
)

type AdsServer struct {
	ClientAds adsv1.AdManagementServiceClient
}

func NewClientAdServer() *AdsServer {
	clientAds := adsv1.NewAdManagementServiceClient(
		http.DefaultClient,
		os.Getenv("AD_SERVER_API"),
	)

	return &AdsServer{
		ClientAds: clientAds,
	}
}
