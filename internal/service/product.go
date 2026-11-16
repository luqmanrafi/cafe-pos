package service

import (
	"context"

	"github.com/luqmanrafi/cafe-pos/internal/model"
)
type productsRepo interface{
	ListActive(ctx context.Context) ([]model.Product, error)
}

type ProductService struct{
	repo productsRepo
}

func NewProductService(repo productsRepo) *ProductService{
	return &ProductService{repo: repo}
}

func (s *ProductService) ListMenu(ctx context.Context) ([]model.Product, error){
	return s.repo.ListActive(ctx)
}