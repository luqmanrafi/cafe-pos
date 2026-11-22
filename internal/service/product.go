package service

import (
	"context"
	"strings"

	"github.com/luqmanrafi/cafe-pos/internal/model"
)
type productsRepo interface{
	ListActive(ctx context.Context) ([]model.Product, error)
	ListAll(ctx context.Context)([]model.Product, error)
	Create(ctx context.Context, p model.Product) error
	SetActive(ctx context.Context, id int, active bool) error
}

type ProductService struct{
	repo productsRepo
}

type ValidationError struct{
	Message string
}

func (e ValidationError) Error() string{
	return e.Message
}

func NewProductService(repo productsRepo) *ProductService{
	return &ProductService{repo: repo}
}

func (s *ProductService) ListMenu(ctx context.Context) ([]model.Product, error){
	return s.repo.ListActive(ctx)
}

func (s *ProductService) ListAll(ctx context.Context)([]model.Product, error){
	return s.repo.ListAll(ctx)
}

func (s *ProductService) Create(ctx context.Context, name string, price int, category string) error{
	name = strings.TrimSpace(name)
	category = strings.TrimSpace(category)

	if name == ""{
		return ValidationError{"Nama menu wajib diisi"}
	}
	if category == ""{
		return ValidationError{"Kategori menu wajib diisi"}
	}
	if price <= 0{
		return ValidationError{"Harga tidak boleh kurang dari 0"}
	}

	return s.repo.Create(ctx, model.Product{
		Name: name,
		Price: price,
		Category: category,
	})
}

func(s *ProductService) SetActive(ctx context.Context, id int, active bool)error{
	return s.repo.SetActive(ctx, id, active)
}