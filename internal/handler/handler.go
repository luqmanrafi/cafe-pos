package handler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/luqmanrafi/cafe-pos/internal/model"
)

type Handler struct {
	tmpl *template.Template
}

func New() (*Handler, error){
	funcs := template.FuncMap{
		"rupiah": formatRupiah,
	}
	tmpl, err := template.New("").Funcs(funcs).ParseGlob("templates/*.html")
	if err != nil{
		return nil, err
	}
	return &Handler{tmpl: tmpl}, nil
}

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", h.health)
	mux.HandleFunc("GET /menu", h.menu)
	return mux
}

func (h *Handler) health(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("ok"))
}

func (h *Handler) menu(w http.ResponseWriter, r *http.Request){
	products := []model.Product{
		{ID: 1, Name: "Latte", Price: 28000, Category: "Kopi", Active: true},
		{ID: 2, Name: "Espresso", Price: 20000, Category: "Kopi", Active: true},
		{ID: 3, Name: "Croissant", Price: 22000, Category: "Dessert", Active: true},
		{ID: 4, Name: "Cheesecake", Price: 24000, Category: "Dessert", Active: false},
	}

	if err := h.tmpl.ExecuteTemplate(w, "menu.html", products); err != nil{
		log.Println("gagal render menu:", err)
		http.Error(w, "terjadi kesalahan", http.StatusInternalServerError)
	}
}
