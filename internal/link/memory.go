package link

import (
	"context"
	"fmt"
	"sync"
)

type MemoryRepository struct {
	mu          sync.RWMutex
	links       map[string]Link
	linksByCode map[string]Link
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		links:       make(map[string]Link),
		linksByCode: make(map[string]Link),
	}
}

func (r *MemoryRepository) Create(ctx context.Context, req Link) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// TODO: probably validate that we somehow didnt end up with identical short code

	r.links[req.ID] = req
	r.linksByCode[req.Code] = req

	return nil
}

func (r *MemoryRepository) FindByCode(ctx context.Context, code string) (Link, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	val, ok := r.linksByCode[code]

	if !ok {
		return Link{}, fmt.Errorf("%w: %s", ErrLinkNotFound, code)
	}

	return val, nil
}

func (r *MemoryRepository) FindByID(ctx context.Context, id string) (Link, error) {
	return Link{}, nil
}

func (r *MemoryRepository) List(ctx context.Context) ([]Link, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	links := make([]Link, 0, len(r.links))
	for _, val := range r.links {
		links = append(links, val)
	}

	return links, nil
}

func (r *MemoryRepository) Delete(ctx context.Context, id string) error {
	return nil
}
