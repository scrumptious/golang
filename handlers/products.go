package handlers

import (
	"log"
	"main/data"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	list := data.GetProducts()
	err := list.ToJSON(w)

	if err != nil {
		http.Error(w, "Unable to convert response to JSON", http.StatusInternalServerError)
	}
}
