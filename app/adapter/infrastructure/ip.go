package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/yuorei/video-server/app/domain"
)

func (i *Infrastructure) GetIPInfomation(ctx context.Context, ip string) (*domain.IPResponse, error) {
	resp, err := http.Get("https://ipinfo.io/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// レスポンスのボディを読み込む
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// レスポンスのボディを構造体にパース
	var ipResp domain.IPResponse
	err = json.Unmarshal(body, &ipResp)
	if err != nil {
		return nil, err
	}
	fmt.Println(ipResp)
	return &ipResp, nil
}
