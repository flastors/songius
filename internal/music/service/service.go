package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/flastors/songius/internal/api"
	"github.com/flastors/songius/internal/music/model"
	"github.com/flastors/songius/internal/music/storage"
	storageModel "github.com/flastors/songius/internal/music/storage/model"
	"github.com/flastors/songius/pkg/api/filter"
	"github.com/flastors/songius/pkg/utils/logging"
)

type Service struct {
	repo      storage.Repository
	apiClient *api.APIClient
	logger    *logging.Logger
}

func NewService(repo storage.Repository, apiClient *api.APIClient, logger *logging.Logger) *Service {
	return &Service{
		repo:      repo,
		apiClient: apiClient,
		logger:    logger,
	}
}

func (s *Service) GetAll(ctx context.Context, filterOptions filter.Options) ([]model.Music, error) {
	storageFilterOptions := storageModel.NewFilterOptions(filterOptions)
	all, err := s.repo.FindAll(ctx, storageFilterOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to get all songs: %v", err)
	}
	return all, nil
}

func (s *Service) GetOne(ctx context.Context, id string) (model.Music, error) {
	m, err := s.repo.FindOne(ctx, id)
	if err != nil {
		return model.Music{}, fmt.Errorf("failed to get song: %v", err)
	}
	return m, nil
}

func (s *Service) Create(ctx context.Context, mdto *model.CreateMusicDTO) (*model.Music, error) {
	// Request to external API
	mApi, err := s.apiClient.GetSongInfo(mdto.Song, mdto.Group)
	if err != nil {
		return nil, fmt.Errorf("failed request to API: %v", err)
	}
	m := model.NewMusicModel(mApi.Song, mApi.Group, mApi.ReleaseDate, mApi.Link, strings.ReplaceAll(mApi.Text, "'", "''"))
	err = s.repo.Create(ctx, m)
	if err != nil {
		return nil, fmt.Errorf("failed to create song: %v", err)
	}
	return m, nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete song: %v", err)
	}
	return nil
}

func (s *Service) Update(ctx context.Context, m *model.Music) error {
	err := s.repo.Update(ctx, *m)
	if err != nil {
		return fmt.Errorf("failed to update song: %v", err)
	}
	return nil
}
