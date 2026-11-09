package link

import (
	"context"
	"fmt"
	"net/url"

	"github.com/google/uuid"
	"go.rtnl.ai/x/randstr"
)

type Repository interface {
	Create(ctx context.Context, l Link) error
	FindByCode(ctx context.Context, code string) (Link, error)
	FindByID(ctx context.Context, id string) (Link, error)
	List(ctx context.Context) ([]Link, error)
	Delete(ctx context.Context, id string) error
}

type Service struct {
	repository Repository
}

func NewService(r Repository) *Service {
	return &Service{repository: r}
}

func (s *Service) CreateLink(ctx context.Context, req LinkCreateRequest) (Link, error) {
	url, err := url.ParseRequestURI(req.Url)
	if err != nil {
		return Link{}, fmt.Errorf("%w: %s", ErrInvalidUrl, req.Url)
	}
	link := Link{
		ID:   uuid.New().String(),
		Url:  url.String(),
		Code: randstr.AlphaNumeric(16),
	}

	if err := s.repository.Create(ctx, link); err != nil {
		return Link{}, fmt.Errorf("failed to create link: %w", err)
	}

	return link, nil
}

func (s *Service) ListLinks(ctx context.Context) ([]Link, error) {
	return s.repository.List(ctx)
}

func (s *Service) FindByCode(ctx context.Context, code string) (Link, error) {
	return s.repository.FindByCode(ctx, code)
}
