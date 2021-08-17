package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/cryptus-neoxys/go-micro/prod-api/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(rw http.ResponseWriter, _ *http.Request) {
	p.l.Println("Handle GET products")
	lp := data.GetProducts()

	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshall JSON", http.StatusInternalServerError)
		return
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST products")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		p.l.Println("Unable to unmarshall JSON")
		http.Error(rw, "Unable to unmarshall JSON", http.StatusInternalServerError)
		return
	}

	p.l.Printf("Prod %#v: ", prod)

	data.AddProduct(prod)
}

func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		p.l.Println("Unable to parse ID")
		http.Error(rw, "Unable to parse ID", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT products", id)

	prod := &data.Product{}

	err = prod.FromJSON(r.Body)
	if err != nil {
		p.l.Println("Unable to unmarshall JSON")
		http.Error(rw, "Unable to unmarshall JSON", http.StatusInternalServerError)
		return
	}
	
	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}