package infrastructure

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/yuorei/video-server/app/domain"
	"github.com/yuorei/video-server/yuovision-proto/go/video/video_grpc"
)

func (i *Infrastructure) GetUserFromDB(ctx context.Context, id string) (*domain.User, error) {
	userPayload, err := i.gRPCClient.UserClient.User(ctx, &video_grpc.UserID{Id: id})
	if err != nil {
		return nil, err
	}
	user := domain.NewUser(userPayload.Id, userPayload.Name, userPayload.ProfileImageUrl, userPayload.SubscribeChannelIds, userPayload.IsSubscribed, userPayload.Role.String())
	return user, nil
}

func (i *Infrastructure) InsertUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	userInput := &video_grpc.UserInput{
		Id:              user.ID,
		Name:            user.Name,
		ProfileImageUrl: user.ProfileImageURL,
		IsSubscribed:    false,
		Role:            video_grpc.Role_NORMAL,
	}

	userPayload, err := i.gRPCClient.UserClient.RegisterUser(ctx, userInput)
	if err != nil {
		return nil, err
	}

	user = domain.NewUser(userPayload.Id, userPayload.Name, userPayload.ProfileImageUrl, userPayload.SubscribeChannelIds, userPayload.IsSubscribed, userPayload.Role.String())
	return user, nil
}

func (i *Infrastructure) GetProfileImageURL(ctx context.Context, id string) (string, error) {
	resp, err := http.Get(os.Getenv("AUTH_URL") + "/profile-image/" + id)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// レスポンスの処理
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var profileImageURL domain.ProfileImageURL
	err = json.Unmarshal(body, &profileImageURL)
	if err != nil {
		return "", err
	}

	return profileImageURL.URL, nil
}

func (i *Infrastructure) AddSubscribeChannelForDB(ctx context.Context, subscribeChannel *domain.SubscribeChannel) (*domain.SubscribeChannel, error) {
	subscribeChannelInput := &video_grpc.SubscribeChannelInput{
		ChannelId: subscribeChannel.ChannelID,
		UserId:    subscribeChannel.UserID,
	}

	subscribeChannelResponse, err := i.gRPCClient.UserClient.SubscribeChannel(ctx, subscribeChannelInput)
	if err != nil {
		return nil, err
	}

	return &domain.SubscribeChannel{
		UserID:    subscribeChannelInput.UserId,
		ChannelID: subscribeChannelInput.ChannelId,
		IsSuccess: subscribeChannelResponse.IsSuccess,
	}, nil
}

func (i *Infrastructure) UnSubscribeChannelForDB(ctx context.Context, subscribeChannel *domain.SubscribeChannel) (*domain.SubscribeChannel, error) {
	subscribeChannelInput := &video_grpc.SubscribeChannelInput{
		ChannelId: subscribeChannel.ChannelID,
		UserId:    subscribeChannel.UserID,
	}

	unSubscribeChannelResponse, err := i.gRPCClient.UserClient.UnSubscribeChannel(ctx, subscribeChannelInput)
	if err != nil {
		return nil, err
	}

	return &domain.SubscribeChannel{
		UserID:    subscribeChannelInput.UserId,
		ChannelID: subscribeChannelInput.ChannelId,
		IsSuccess: unSubscribeChannelResponse.IsSuccess,
	}, nil
}
