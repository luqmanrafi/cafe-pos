package handler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/luqmanrafi/cafe-pos/internal/service"
)

type Handler struct {
	tmpl *template.Template
	products *service.ProductService
}

func New(products *service.ProductService) (*Handler, error){
	funcs := template.FuncMap{
		"rupiah": formatRupiah,
	}
	tmpl, err := template.New("").Funcs(funcs).ParseGlob("templates/*.html")
	if err != nil{
		return nil, err
	}
	return &Handler{tmpl: tmpl, products: products}, nil
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
	products, err := h.products.ListMenu(r.Context())
	if err != nil{
		log.Println("gagal mengambil menu: ", err)
		http.Error(w, "terjadi kesalahan", http.StatusInternalServerError)
		return
	}

	if err := h.tmpl.ExecuteTemplate(w, "menu.html", products); err != nil{
		log.Println("gagal render menu: ", err )
		http.Error(w, "terjadi kesalahan", http.StatusInternalServerError)
	}
}
