package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/flastors/songius/internal/config"
	musicModel "github.com/flastors/songius/internal/music/model"
)

type APIClient struct {
	config *config.ExternalAPIConfig
}

func NewAPIClient(config *config.ExternalAPIConfig) *APIClient {
	return &APIClient{
		config: config,
	}
}

func (ac *APIClient) GetSongInfo(song, group string) (*musicModel.APIMusicDTO, error) {
	resp, err := http.Get(fmt.Sprintf("%s/info?song=%s&group=%s", ac.config.Url, song, group))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusBadRequest {
			return nil, fmt.Errorf("bad request")
		} else {
			return nil, fmt.Errorf("no responce from API")
		}
	}

	var m musicModel.APIMusicDTO
	if err := json.NewDecoder(resp.Body).Decode(&m); err != nil {
		return nil, err
	}
	m.Song = song
	m.Group = group

	return &m, nil
}
