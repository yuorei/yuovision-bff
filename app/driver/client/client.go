package client

import (
	"log"

	"github.com/yuorei/video-server/yuovision-proto/go/video/video_grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	VideoClient   video_grpc.VideoServiceClient
	UserClient    video_grpc.UserServiceClient
	CommentClient video_grpc.CommentServiceClient
}

func NewClient() *Client {
	client := &Client{}
	client.NewConnect()
	return client
}

func (c *Client) NewConnect() {
	// TCPサーバーのアドレスを指定
	userAddress := "localhost:50051"
	// サーバーに接続する
	conn, err := grpc.Dial(
		userAddress,
		// コネクションでSSL/TLSを使用しない
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// コネクションが確立されるまで待機する(同期処理をする)
		grpc.WithBlock(),
	)

	if err != nil {
		log.Fatal("Connection failed. err: ", err)
		return
	}

	c.VideoClient = video_grpc.NewVideoServiceClient(conn)
	c.UserClient = video_grpc.NewUserServiceClient(conn)
	c.CommentClient = video_grpc.NewCommentServiceClient(conn)
}
