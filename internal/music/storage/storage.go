package storage

import (
	"context"

	music "github.com/flastors/songius/internal/music/model"
)

type Repository interface {
	Create(ctx context.Context, music *music.Music) error
	FindAll(ctx context.Context, filterOptions FilterOptions) ([]music.Music, error)
	FindOne(ctx context.Context, id string) (music.Music, error)
	Update(ctx context.Context, music music.Music) error
	Delete(ctx context.Context, id string) error
}

type FilterOptions interface {
	PaginationQuery() string
	FilterQuery() (string, error)
}
