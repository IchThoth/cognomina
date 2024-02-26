package tube

import (
	"net/http"
	"time"
)

type YTClient struct {
	hc http.Client
}

type TubeSearchResult struct {
	name     string
	link     string
	duration time.Time
}

func (c *YTClient) GetDownloadLinks() (*TubeSearchResult, error) {

}
