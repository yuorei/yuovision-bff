package domain

import (
	"fmt"
	"io"
	"time"
)

type (
	Video struct {
		ID                string
		VideoURL          string
		ThumbnailImageURL string
		Title             string
		Description       *string
		Tags              []string
		WatchCount        int
		IsPrivate         bool
		IsAdult           bool
		IsExternalCutout  bool
		IsAd              bool
		UploaderID        string
		CreatedAt         time.Time
		UpdatedAt         time.Time
	}

	UploadVideo struct {
		ID               string
		Video            io.ReadSeeker
		VideoContentType string
		ThumbnailImage   *io.ReadSeeker
		ImageContentType string
		Title            string
		Description      *string
		Tags             []string
		IsPrivate        bool
		IsAdult          bool
		IsExternalCutout bool
		IsAd             bool
		ExternalCutout   bool
	}

	UploadVideoResponse struct {
		ID                string
		VideoURL          string
		ThumbnailImageURL string
		Title             string
		Description       *string
		Tags              []string
		IsPrivate         bool
		IsAdult           bool
		IsExternalCutout  bool
		IsAd              bool
		UploaderID        string
		CreatedAt         time.Time
	}

	VideoFile struct {
		ID    string
		Video io.ReadSeeker
	}

	ThumbnailImage struct {
		ID             string
		ContentType    string
		ThumbnailImage []byte
	}
)

func NewVideoID() string {
	return fmt.Sprintf("%s%s%s", "video", IDSeparator, NewUUID())
}

func NewVideo(id string, videoURL string, thumbnailImageURL string, title string, description *string, tags []string, watchCount int, private bool, adult bool, externalCutout bool, isAd bool, uploaderID string, createdAt time.Time, updatedAt time.Time) *Video {
	return &Video{
		ID:                id,
		VideoURL:          videoURL,
		ThumbnailImageURL: thumbnailImageURL,
		Title:             title,
		Description:       description,
		Tags:              tags,
		IsPrivate:         private,
		IsAdult:           adult,
		IsExternalCutout:  externalCutout,
		IsAd:              isAd,
		UploaderID:        uploaderID,
		CreatedAt:         createdAt,
		UpdatedAt:         updatedAt,
		WatchCount:        watchCount,
	}
}

func NewUploadVideo(id string, video io.ReadSeeker, videoContentType string, thumbnailImage *io.ReadSeeker, imageContentType string, title string, description *string, tags []string, private bool, adult bool, externalCutout bool, isAd bool) *UploadVideo {
	return &UploadVideo{
		ID:               id,
		Video:            video,
		VideoContentType: videoContentType,
		ThumbnailImage:   thumbnailImage,
		ImageContentType: imageContentType,
		Title:            title,
		Description:      description,
		Tags:             tags,
		IsPrivate:        private,
		IsAdult:          adult,
		IsExternalCutout: externalCutout,
		IsAd:             isAd,
		ExternalCutout:   externalCutout,
	}
}

func NewVideoFile(id string, video io.ReadSeeker) *VideoFile {
	return &VideoFile{
		ID:    id,
		Video: video,
	}
}

func NewThumbnailImage(id, contentType string, thumbnailImage []byte) *ThumbnailImage {
	return &ThumbnailImage{
		ID:             id,
		ContentType:    contentType,
		ThumbnailImage: thumbnailImage,
	}
}
