package service

import (
	"context"
	"strings"

	"github.com/luqmanrafi/cafe-pos/internal/model"
)
type productsRepo interface{
	ListActive(ctx context.Context) ([]model.Product, error)
	ListAll(ctx context.Context)([]model.Product, error)
	GetByID(ctx context.Context, id int)(model.Product, error)
	Create(ctx context.Context, p model.Product) error
	Update(ctx context.Context, p model.Product) error
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

func validateProduct(name string, price int, category string)(model.Product, error){
	name = strings.TrimSpace(name)
	category = strings.TrimSpace(category)

	if name == ""{
		return model.Product{}, ValidationError{"Nama menu wajib diisi!"}
	}

	if category == ""{
		return model.Product{}, ValidationError{"Kategori wajib diisi!"}
	}
	if price <= 0 {
		return model.Product{}, ValidationError{"Harga harus lebih dari 0"}
	}

	return model.Product{Name: name, Price: price, Category: category}, nil
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

func(s *ProductService) GetByID(ctx context.Context, id int)(model.Product, error){
	return s.repo.GetByID(ctx, id)
}

func (s *ProductService) Create(ctx context.Context, name string, price int, category string) error{
	p, err := validateProduct(name, price, category)
	if err != nil{
		return err
	}

	return s.repo.Create(ctx, p)
}

func(s *ProductService) Update(ctx context.Context, id int, name string, price int, category string)error{
	p, err := validateProduct(name, price, category)
	if err != nil{
		return err
	}
	p.ID = id
	return s.repo.Update(ctx, p)
}

func(s *ProductService) SetActive(ctx context.Context, id int, active bool)error{
	return s.repo.SetActive(ctx, id, active)
}