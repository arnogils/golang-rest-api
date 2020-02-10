package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Product Struct (Model)
type Product struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

var products []Product

// getProducts : return all the available products
func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	json.NewEncoder(w).Encode(products)
}

// getProduct : return a specific products
func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	params := mux.Vars(r)
	for _, item := range products {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Product{})
}

func main() {
	router := mux.NewRouter()

	products = append(products, Product{ID: "1234", Description: "iPhone"})
	products = append(products, Product{ID: "4567", Description: "T-shirt"})

	router.HandleFunc("/products", getProducts).Methods("GET")
	router.HandleFunc("/products/{id}", getProduct).Methods("GET")

	log.Fatal(http.ListenAndServe(":5000", router))
}
