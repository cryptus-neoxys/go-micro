package handlers

import (
	"log"
	"net/http"

	"github.com/cryptus-neoxys/go-micro/prod-api/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	} else if r.Method == http.MethodPost {
		p.addProduct(rw, r)
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, _ *http.Request) {
	p.l.Println("Handle GET products")
	lp := data.GetProducts()

	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshall JSON", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST products")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshall JSON", http.StatusInternalServerError)
	}

	p.l.Printf("Prod %#v: ", prod)

	data.AddProduct(prod)
}
