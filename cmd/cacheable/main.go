package main

import (
	"context"

	"github.com/schnell18/play-golang/cacheable/product"
)

func main() {
	svc := &Service{repo: NewRepo()}

	ctx := context.Background()

	// Read → cached
	p1, _ := product.CachedGetByID(svc, ctx, "123")

	// Update → evicts cache
	_ = product.CachedUpdateProduct(
		svc,
		ctx,
		&Product{ID: "123", Name: "Book", Price: 19.99},
	)

	// Next read → reloads fresh from DB and re-caches
	p2, _ := product.CachedGetByID(svc, ctx, "123")

	// Delete → evicts cache
	_ = product.CachedDeleteProduct(svc, ctx, "123")
}
