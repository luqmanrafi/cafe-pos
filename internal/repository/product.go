package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
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

func (r *ProductRepository) ListAll(ctx context.Context)([]model.Product, error){
	rows, err := r.db.Query(ctx, `
	SELECT id, name, price, category, active 
	FROM products 
	ORDER BY active DESC, category, name`)
	
	if err != nil{
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

func(r *ProductRepository) SetActive(ctx context.Context, id int, active bool) error{
	tag, err := r.db.Exec(ctx, 
		`UPDATE products SET active = $1, updated_at = now() WHERE id = $2`, active, id)
	if err != nil{
		return err
	}

	if tag.RowsAffected() == 0 {
		return model.ErrNotFound
	}
	return nil
}

func(r *ProductRepository)GetByID(ctx context.Context, id int)(model.Product, error){
	var p model.Product
	err := r.db.QueryRow(ctx, `SELECT id, name,  price, category, active
	FROM products
	WHERE id = $1`, id).Scan(&p.ID, &p.Name, &p.Price, &p.Category, &p.Active)

	if errors.Is(err, pgx.ErrNoRows){
		return model.Product{}, model.ErrNotFound
	}

	if err != nil {
		return model.Product{}, err
	}
	return p, nil	
}

func(r *ProductRepository)Update(ctx context.Context, p model.Product)error{
	tag, err := r.db.Exec(ctx, `UPDATE products
	SET name = $1, price = $2, category = $3, updated_at = now()
	WHERE id = $4`, p.Name, p.Price, p.Category, p.ID)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0{
		return model.ErrNotFound
	}
	return nil
}