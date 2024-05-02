package infrastructure

import (
	"context"

	"github.com/yuorei/video-server/app/domain"
	"github.com/yuorei/video-server/yuovision-proto/go/video/video_grpc"
)

func (i *Infrastructure) GetCommentsByVideoIDFromDB(ctx context.Context, videoID string) ([]*domain.Comment, error) {
	comment, err := i.gRPCClient.CommentClient.CommentsByVideo(ctx, &video_grpc.CommentsByVideoInput{VideoId: videoID})
	if err != nil {
		return nil, err
	}

	comments := make([]*domain.Comment, 0, len(comment.Comments))
	for _, c := range comment.Comments {
		comments = append(comments, domain.NewComment(c.Id, c.Video.Id, c.Text, c.CreatedAt.AsTime(), c.CreatedAt.AsTime(), &domain.User{ID: c.UserId, Name: c.Name}))
	}
	return comments, nil
}

func (i *Infrastructure) InsertComment(ctx context.Context, postComment *domain.Comment) (*domain.Comment, error) {
	commentInput := &video_grpc.PostCommentInput{
		VideoId: postComment.VideoID,
		UserId:  postComment.User.ID,
		Text:    postComment.Text,
		Name:    postComment.User.Name,
	}

	comment, err := i.gRPCClient.CommentClient.PostComment(ctx, commentInput)
	if err != nil {
		return nil, err
	}

	return domain.NewComment(comment.Id, comment.Video.Id, comment.Text, comment.CreatedAt.AsTime(), comment.CreatedAt.AsTime(), &domain.User{ID: comment.UserId, Name: comment.Name}), nil
}
