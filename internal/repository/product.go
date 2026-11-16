package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/luqmanrafi/cafe-pos/internal/model"
)

type ProductRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) *ProductRepository{
	return &ProductRepository{db: db}
}

func(r *ProductRepository) ListActive(ctx context.Context) ([]model.Product, error){
	rows, err := r.db.Query(ctx, `
	SELECT id, name, price, category, active 
	FROM products 
	WHERE active = TRUE 
	ORDER BY category, name`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []model.Product
	for rows.Next(){
		var p model.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Category, &p.Active); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, rows.Err()
}

func(r *ProductRepository) Create(ctx context.Context, p model.Product) error{
	_, err := r.db.Exec(ctx, 
	`INSERT INTO products (name, price, category) 
	VALUES ($1, $2, $3)`,
	p.Name, p.Price, p.Category)
	return err
}