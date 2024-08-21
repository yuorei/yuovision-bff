package infrastructure

import (
	"github.com/redis/go-redis/v9"
	"github.com/yuorei/video-server/app/driver/client"
	r "github.com/yuorei/video-server/app/driver/redis"
)

type Infrastructure struct {
	redis      *redis.Client
	gRPCClient *client.Client
	adsServer  *client.AdsServer
}

func NewInfrastructure() *Infrastructure {
	return &Infrastructure{
		redis:      r.ConnectRedis(),
		gRPCClient: client.NewClient(),
		adsServer:  client.NewClientAdServer(),
	}
}
