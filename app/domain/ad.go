package domain

type Ad struct {
	ID           string `json:"id,omitempty"`
	AdURL        string `json:"ad_url,omitempty"`
	Title        string `json:"title,omitempty"`
	Description  string `json:"description,omitempty"`
	ThumbnailURL string `json:"thumbnail_url,omitempty"`
	VideoURL     string `json:"video_url,omitempty"`
}

type GetAdVideoRequest struct {
	// ブラウザ情報
	UserAgent            string `json:"user_agent,omitempty"`
	Platform             string `json:"platform,omitempty"`
	Language             string `json:"language,omitempty"`
	URL                  string `json:"url,omitempty"`
	PageTitle            string `json:"page_title,omitempty"`
	Referrer             string `json:"referrer,omitempty"`
	NetworkDownlink      string `json:"network_downlink,omitempty"`
	NetworkEffectiveType string `json:"network_effective_type,omitempty"`
	IPAddress            string `json:"ip_address,omitempty"`
	Location             string `json:"location,omitempty"`
	Hostname             string `json:"hostname,omitempty"`
	City                 string `json:"city,omitempty"`
	Region               string `json:"region,omitempty"`
	Country              string `json:"country,omitempty"`
	Org                  string `json:"org,omitempty"`
	Postal               string `json:"postal,omitempty"`
	Timezone             string `json:"timezone,omitempty"`
	// ビデオ情報
	VideoID     string   `json:"video_id,omitempty"`
	Title       string   `json:"title,omitempty"`
	Description *string  `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	// ユーザー情報
	UserID   string `json:"user_id,omitempty"`
	ClientID string `json:"client_id,omitempty"`
}

type WatchCountAdVideoRequest struct {
	// ブラウザ情報
	UserAgent            string `json:"user_agent,omitempty"`
	Platform             string `json:"platform,omitempty"`
	Language             string `json:"language,omitempty"`
	URL                  string `json:"url,omitempty"`
	PageTitle            string `json:"page_title,omitempty"`
	Referrer             string `json:"referrer,omitempty"`
	NetworkDownlink      string `json:"network_downlink,omitempty"`
	NetworkEffectiveType string `json:"network_effective_type,omitempty"`
	IPAddress            string `json:"ip_address,omitempty"`
	Location             string `json:"location,omitempty"`
	Hostname             string `json:"hostname,omitempty"`
	City                 string `json:"city,omitempty"`
	Region               string `json:"region,omitempty"`
	Country              string `json:"country,omitempty"`
	Org                  string `json:"org,omitempty"`
	Postal               string `json:"postal,omitempty"`
	Timezone             string `json:"timezone,omitempty"`
	// ビデオ情報
	VideoID     string   `json:"video_id,omitempty"`
	Title       string   `json:"title,omitempty"`
	Description *string  `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	// ユーザー情報
	UserID   string `json:"user_id,omitempty"`
	ClientID string `json:"client_id,omitempty"`

	// 広告情報
	AdID string `json:"ad_id,omitempty"`
}

func NewAd(ID, adURL, title, description, thumbnailURL, videoURL string) *Ad {
	return &Ad{
		ID:           ID,
		AdURL:        adURL,
		Title:        title,
		Description:  description,
		ThumbnailURL: thumbnailURL,
		VideoURL:     videoURL,
	}
}

func NewAdVideoRequest(userAgent, platform, language, url, pageTitle, referrer, networkDownlink, networkEffectiveType, ipAddress, location, hostname, city, region, country, org, postal, timezone, videoID, title, userID, clientID string, description *string, tags []string) *GetAdVideoRequest {
	return &GetAdVideoRequest{
		UserAgent:            userAgent,
		Platform:             platform,
		Language:             language,
		URL:                  url,
		PageTitle:            pageTitle,
		Referrer:             referrer,
		NetworkDownlink:      networkDownlink,
		NetworkEffectiveType: networkEffectiveType,
		IPAddress:            ipAddress,
		Location:             location,
		Hostname:             hostname,
		City:                 city,
		Region:               region,
		Country:              country,
		Org:                  org,
		Postal:               postal,
		Timezone:             timezone,
		VideoID:              videoID,
		Title:                title,
		Description:          description,
		Tags:                 tags,
		UserID:               userID,
		ClientID:             clientID,
	}
}
