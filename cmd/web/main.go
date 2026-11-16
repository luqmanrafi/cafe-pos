package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/luqmanrafi/cafe-pos/internal/handler"
	"github.com/luqmanrafi/cafe-pos/internal/repository"
)

func main() {
	_ = godotenv.Load()
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == ""{
		log.Fatal("DATABASE_URL belum diisi")
	}
	ctx := context.Background()

	pool, err := repository.NewPool(ctx, dbUrl)
	if err != nil{
		log.Fatal("gagal koneksi ke database: ", err)
	}
	defer pool.Close()

	productsRepo := repository.NewProductRepository(pool)
	products, err := productsRepo.ListActive(ctx)
	if err != nil {
		log.Fatal("gagal mengambil menu: ", err)
	}
	for _, p := range products{
		log.Println(p.Name, p.Price, p.Category)
	}
	
	h, err := handler.New()
	if err != nil{
		log.Fatal(err)
	}

	log.Println("Server berjalan pada port :8080")
	if err := http.ListenAndServe(":8080", h.Routes()); err != nil{
		log.Fatal(err)
	}
}