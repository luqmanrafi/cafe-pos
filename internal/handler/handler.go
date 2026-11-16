package handler

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/luqmanrafi/cafe-pos/internal/service"
)

type Handler struct {
	tmpl *template.Template
	products *service.ProductService
}

type menuFormData struct{
	Error 		string
	Name 		string
	Price 		string
	Category 	string
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
	mux.HandleFunc("GET /menu/new", h.menuForm)
	mux.HandleFunc("POST /menu", h.menuCreate)
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

func (h *Handler) menuForm(w http.ResponseWriter, r *http.Request){
	if err := h.tmpl.ExecuteTemplate(w, "menu_form.html", menuFormData{}); err != nil {
		log.Println("gagal render form: ", err)
		http.Error(w, "terjadi kesalahan", http.StatusInternalServerError)
	}
}

func(h *Handler) menuCreate(w http.ResponseWriter, r *http.Request){
	name := r.PostFormValue("name")
	priceStr := r.PostFormValue("price")
	category := r.PostFormValue("category")

	data := menuFormData{
		Name: name,
		Price: priceStr,
		Category: category,
	}

	price, err := strconv.Atoi(strings.TrimSpace(priceStr))
	if err != nil{
		data.Error = "harga harus berupa angka!"
		h.renderForm(w, data)
		return 
	}

	if err := h.products.Create(r.Context(), name, price, category); err != nil{
		var ve service.ValidationError
		if errors.As(err, &ve){
			data.Error = ve.Message
			h.renderForm(w, data)
			return
		}
		log.Println("gagal simpan data: ", err)
		http.Error(w, "terjadi kesalahan", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/menu", http.StatusSeeOther)
}

func(h *Handler)renderForm(w http.ResponseWriter, data menuFormData){
	w.WriteHeader(http.StatusBadRequest)
	if err := h.tmpl.ExecuteTemplate(w, "menu_form.html", data); err != nil{
		log.Println("gagal render form: ", err)
	}
}
