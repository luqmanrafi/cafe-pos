package main

import (
	"log"
	"net/http"

	"github.com/luqmanrafi/cafe-pos/internal/handler"
)

func main() {
	h, err := handler.New()
	if err != nil{
		log.Fatal(err)
	}

	log.Println("Server berjalan pada port :8080")
	if err := http.ListenAndServe(":8080", h.Routes()); err != nil{
		log.Fatal(err)
	}
}