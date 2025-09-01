// Package product defines product APIs.
//
//go:generate go run ../../cmd/cachegen
package product

import "context"

type Product struct {
	ID    string
	Name  string
	Price float64
}

type Service struct {
	repo Repository
}

type Repository interface {
	Update(ctx context.Context, p *Product) error
	Delete(ctx context.Context, id string) error
	FindByID(ctx context.Context, id string) (*Product, error)
}

// GetByID retrieves product by id.
// cache:cacheable name=product GetByID ttl=10m key=product:{id}
func (s *Service) GetByID(ctx context.Context, id string) (*Product, error) {
	return s.repo.FindByID(ctx, id)
}

// UpdateProduct updates product by id.
// cache:evict name=product key=product:{p.ID}
func (s *Service) UpdateProduct(ctx context.Context, p *Product) error {
	return s.repo.Update(ctx, p)
}

// DeleteProduct deletes product by id.
// cache:evict name=product key=product:{id}
func (s *Service) DeleteProduct(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
