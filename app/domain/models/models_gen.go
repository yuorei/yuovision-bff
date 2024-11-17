// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

type Node interface {
	IsNode()
	GetID() string
}

type AdVideoInput struct {
	// ブラウザ情報
	UserAgent            string  `json:"userAgent"`
	Platform             string  `json:"platform"`
	Language             string  `json:"language"`
	URL                  string  `json:"url"`
	PageTitle            string  `json:"pageTitle"`
	Referrer             *string `json:"referrer,omitempty"`
	NetworkDownlink      *string `json:"networkDownlink,omitempty"`
	NetworkEffectiveType *string `json:"networkEffectiveType,omitempty"`
	IPAddress            string  `json:"ipAddress"`
	Location             string  `json:"location"`
	Hostname             string  `json:"hostname"`
	City                 string  `json:"city"`
	Region               string  `json:"region"`
	Country              string  `json:"country"`
	Org                  string  `json:"org"`
	Postal               string  `json:"postal"`
	Timezone             string  `json:"timezone"`
	// ビデオ情報
	VideoID string `json:"videoID"`
	// ユーザー情報
	UserID   string `json:"userID"`
	ClientID string `json:"clientID"`
}

type AdVideoPayload struct {
	AdID         string `json:"adID"`
	AdURL        string `json:"adURL"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	ThumbnailURL string `json:"thumbnailURL"`
	VideoURL     string `json:"videoURL"`
}

type Comment struct {
	ID        string `json:"id"`
	Video     *Video `json:"video"`
	Text      string `json:"text"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	User      *User  `json:"user"`
}

func (Comment) IsNode()            {}
func (this Comment) GetID() string { return this.ID }

type CutVideoInput struct {
	VideoID   string `json:"VideoID"`
	StartTime int    `json:"StartTime"`
	EndTime   int    `json:"EndTime"`
}

type CutVideoPayload struct {
	CutVideoURL string `json:"cutVideoURL"`
}

type IncrementWatchCountInput struct {
	VideoID string `json:"VideoID"`
	UserID  string `json:"UserID"`
}

type IncrementWatchCountPayload struct {
	WatchCount int `json:"watchCount"`
}

type Mutation struct {
}

type PostCommentInput struct {
	VideoID string `json:"videoID"`
	Text    string `json:"text"`
}

type PostCommentPayload struct {
	ID        string `json:"id"`
	Video     *Video `json:"video"`
	Text      string `json:"text"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	User      *User  `json:"user"`
}

type Query struct {
}

type SubscriptionPayload struct {
	IsSuccess bool `json:"isSuccess"`
}

type UploadVideoInput struct {
	Video            graphql.Upload  `json:"video"`
	ThumbnailImage   *graphql.Upload `json:"thumbnailImage,omitempty"`
	Title            string          `json:"title"`
	Description      *string         `json:"description,omitempty"`
	Tags             []*string       `json:"tags,omitempty"`
	IsPrivate        bool            `json:"isPrivate"`
	IsAdult          bool            `json:"isAdult"`
	IsExternalCutout bool            `json:"isExternalCutout"`
	IsAds            bool            `json:"isAds"`
}

type User struct {
	ID                  string   `json:"id"`
	Name                string   `json:"name"`
	ProfileImageURL     string   `json:"profileImageURL"`
	IsSubscribed        bool     `json:"isSubscribed"`
	Role                Role     `json:"role"`
	Subscribechannelids []string `json:"subscribechannelids"`
	Videos              []*Video `json:"videos"`
}

func (User) IsNode()            {}
func (this User) GetID() string { return this.ID }

type UserInput struct {
	Name         string `json:"name"`
	IsSubscribed bool   `json:"isSubscribed"`
	Role         Role   `json:"role"`
}

type UserPayload struct {
	ID                  string   `json:"id"`
	Name                string   `json:"name"`
	ProfileImageURL     string   `json:"profileImageURL"`
	IsSubscribed        bool     `json:"isSubscribed"`
	Role                Role     `json:"role"`
	Subscribechannelids []string `json:"subscribechannelids"`
}

type Video struct {
	ID                string    `json:"id"`
	VideoURL          string    `json:"videoURL"`
	Title             string    `json:"title"`
	ThumbnailImageURL string    `json:"thumbnailImageURL"`
	Description       *string   `json:"description,omitempty"`
	Tags              []*string `json:"Tags,omitempty"`
	IsPrivate         bool      `json:"isPrivate"`
	IsAdult           bool      `json:"isAdult"`
	IsExternalCutout  bool      `json:"isExternalCutout"`
	IsAd              bool      `json:"isAd"`
	WatchCount        int       `json:"watchCount"`
	CreatedAt         string    `json:"createdAt"`
	UpdatedAt         string    `json:"updatedAt"`
	Uploader          *User     `json:"uploader"`
}

func (Video) IsNode()            {}
func (this Video) GetID() string { return this.ID }

type VideoPayload struct {
	ID                string    `json:"id"`
	VideoURL          string    `json:"videoURL"`
	Title             string    `json:"title"`
	ThumbnailImageURL string    `json:"thumbnailImageURL"`
	Description       *string   `json:"description,omitempty"`
	Tags              []*string `json:"tags,omitempty"`
	IsPrivate         bool      `json:"isPrivate"`
	IsAdult           bool      `json:"isAdult"`
	IsExternalCutout  bool      `json:"isExternalCutout"`
	IsAd              bool      `json:"isAd"`
	WatchCount        int       `json:"watchCount"`
	CreatedAt         string    `json:"createdAt"`
	UpdatedAt         string    `json:"updatedAt"`
	Uploader          *User     `json:"uploader"`
}

type WatchCountAdVideoInput struct {
	// ブラウザ情報
	UserAgent            string  `json:"userAgent"`
	Platform             string  `json:"platform"`
	Language             string  `json:"language"`
	URL                  string  `json:"url"`
	PageTitle            string  `json:"pageTitle"`
	Referrer             *string `json:"referrer,omitempty"`
	NetworkDownlink      *string `json:"networkDownlink,omitempty"`
	NetworkEffectiveType *string `json:"networkEffectiveType,omitempty"`
	IPAddress            string  `json:"ipAddress"`
	Location             string  `json:"location"`
	Hostname             string  `json:"hostname"`
	City                 string  `json:"city"`
	Region               string  `json:"region"`
	Country              string  `json:"country"`
	Org                  string  `json:"org"`
	Postal               string  `json:"postal"`
	Timezone             string  `json:"timezone"`
	// ビデオ情報
	VideoID     string   `json:"videoID"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Tags        []string `json:"tags,omitempty"`
	// ユーザー情報
	UserID   string `json:"userID"`
	ClientID string `json:"clientID"`
	// 広告情報
	AdID string `json:"adID"`
}

type WatchCountAdVideoPayload struct {
	Success bool `json:"success"`
}

type SubscribeChannelInput struct {
	ChannelID string `json:"channelID"`
}

type Role string

const (
	RoleAdmin  Role = "ADMIN"
	RoleNormal Role = "NORMAL"
	RoleAds    Role = "ADS"
)

var AllRole = []Role{
	RoleAdmin,
	RoleNormal,
	RoleAds,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleAdmin, RoleNormal, RoleAds:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
