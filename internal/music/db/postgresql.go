package db

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/flastors/songius/internal/music/model"
	"github.com/flastors/songius/internal/music/storage"
	"github.com/flastors/songius/pkg/client/postgresql"
	"github.com/flastors/songius/pkg/utils/logging"
	"github.com/jackc/pgconn"
)

func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

// Create implements music.Repository.
func (r *repository) Create(ctx context.Context, music *model.Music) error {
	q := `
		INSERT INTO music (song, artist, release_date, link, text)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
		`
	r.logger.Trace(formatQuery(q))
	if err := r.client.QueryRow(ctx, q, music.Song, music.Group, music.ReleaseDate, music.Link, music.Text).Scan(&music.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf("SQL Error: %s, Where: %s, Code: %s, Detail: %s, SQLState: %s", pgErr.Message, pgErr.Where, pgErr.Code, pgErr.Detail, pgErr.SQLState())
			return newErr
		}
		return err
	}
	return nil
}

// Delete implements music.Repository.
func (r *repository) Delete(ctx context.Context, id string) error {
	q := `
		DELETE FROM public.music
		WHERE id = $1
	`
	r.logger.Trace(formatQuery(q))

	_, err := r.client.Exec(ctx, q, id)
	if err != nil {
		return err
	}
	return nil
}

// FindAll implements music.Repository.
func (r *repository) FindAll(ctx context.Context, filterOptions storage.FilterOptions) ([]model.Music, error) {
	filterQuery, err := filterOptions.FilterQuery()
	if err != nil {
		return nil, err
	}
	paginationQuery := filterOptions.PaginationQuery()
	q := fmt.Sprintf(`
		SELECT id, song, artist, TO_CHAR(release_date, 'dd.mm.yyyy'), link, text
		FROM public.music
		%s
		%s
	`, filterQuery, paginationQuery)
	r.logger.Trace(formatQuery(q))

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	songs := make([]model.Music, 0)
	for rows.Next() {
		var m model.Music
		if err := rows.Scan(&m.ID, &m.Song, &m.Group, &m.ReleaseDate, &m.Link, &m.Text); err != nil {
			return nil, err
		}
		songs = append(songs, m)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return songs, nil
}

// FindById implements music.Repository.
func (r *repository) FindOne(ctx context.Context, id string) (model.Music, error) {
	q := `
		SELECT id, song, artist, TO_CHAR(release_date, 'dd.mm.yyyy'), link, text
		FROM public.music
		WHERE id = $1
	`
	var m model.Music
	r.logger.Trace(formatQuery(q))
	err := r.client.QueryRow(ctx, q, id).Scan(&m.ID, &m.Song, &m.Group, &m.ReleaseDate, &m.Link, &m.Text)
	if err != nil {
		return model.Music{}, err
	}
	return m, nil
}

// Update implements music.Repository.
func (r *repository) Update(ctx context.Context, music model.Music) error {
	q := `
		UPDATE public.music
		SET song = $2, artist = $3, release_date = $4, link = $5, text = $6
		WHERE id = $1
	`
	r.logger.Trace(formatQuery(q))
	_, err := r.client.Exec(ctx, q, music.ID, music.Song, music.Group, music.ReleaseDate, music.Link, music.Text)
	if err != nil {
		return err
	}
	return nil
}

func NewRepository(client postgresql.Client, logger *logging.Logger) storage.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}
