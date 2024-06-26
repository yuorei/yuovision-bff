package domain

import (
	"fmt"
	"time"
)

func NewCommentID() string {
	return fmt.Sprintf("%s%s%s", "comment", IDSeparator, NewUUID())
}

type (
	Comment struct {
		ID        string
		VideoID   string
		Text      string
		CreatedAt time.Time
		UpdatedAt time.Time
		User      *User
	}
)

func NewComment(id, videoID, text string, createdAt, updatedAt time.Time, user *User) *Comment {
	return &Comment{
		ID:        id,
		VideoID:   videoID,
		Text:      text,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		User:      user,
	}
}

func NewPostComment(id, videoID, userID, name, text string) *Comment {
	return &Comment{
		ID:        id,
		VideoID:   videoID,
		Text:      text,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		User: &User{
			ID:   userID,
			Name: name,
		},
	}
}
