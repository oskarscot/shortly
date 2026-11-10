package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/oskarscot/shortly/internal/link"
	repo "github.com/oskarscot/shortly/internal/postgres/sqlc"
)

type LinkRepository struct {
	queries *repo.Queries
}

var _ link.Repository = (*LinkRepository)(nil)

func NewLinkRepository(db repo.DBTX) *LinkRepository {
	return &LinkRepository{queries: repo.New(db)}
}

func (r *LinkRepository) Create(ctx context.Context, l link.Link) error {
	return r.queries.Create(ctx, repo.CreateParams{
		ID:   l.ID,
		Code: l.Code,
		Url:  l.Url,
	})
}

func (r *LinkRepository) FindByID(ctx context.Context, id string) (link.Link, error) {
	row, err := r.queries.FindByID(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return link.Link{}, fmt.Errorf("%w: %s", link.ErrLinkNotFound, id)
	}
	if err != nil {
		return link.Link{}, fmt.Errorf("find link by id: %w", err)
	}
	return link.Link{
		ID:   row.ID,
		Code: row.Code,
		Url:  row.Url,
	}, nil
}

func (r *LinkRepository) List(ctx context.Context) ([]link.Link, error) {
	rows, err := r.queries.List(ctx)
	links := make([]link.Link, 0, len(rows))

	if err != nil {
		return []link.Link{}, err
	}

	for _, val := range rows {
		links = append(links, link.Link{
			ID:   val.ID,
			Code: val.Code,
			Url:  val.Url,
		})
	}

	return links, nil
}

func (r *LinkRepository) Delete(ctx context.Context, id string) error {
	return r.queries.Delete(ctx, id)
}

func (r *LinkRepository) FindByCode(ctx context.Context, code string) (link.Link, error) {
	row, err := r.queries.FindByCode(ctx, code)
	if errors.Is(err, pgx.ErrNoRows) {
		return link.Link{}, fmt.Errorf("%w: %s", link.ErrLinkNotFound, code)
	}
	if err != nil {
		return link.Link{}, fmt.Errorf("find link by code: %w", err)
	}

	return link.Link{
		ID:   row.ID,
		Code: row.Code,
		Url:  row.Url,
	}, nil
}
