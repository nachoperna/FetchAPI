package main

import (
	"FetchAPI/datos"
	"FetchAPI/views"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", ShoeHome)
	http.HandleFunc("/users", ListUsers)
	http.HandleFunc("/products", ListProducts)

	http.ListenAndServe(":8080", nil)
}

func ShoeHome(w http.ResponseWriter, r *http.Request) {
	views.Home().Render(r.Context(), w)
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	url_api := "https://jsonplaceholder.typicode.com/users"
	respuesta, err := http.DefaultClient.Get(url_api)
	if err != nil {
		fmt.Println("Error al obtener respuesta")
		return
	}
	if respuesta.StatusCode != 200 {
		fmt.Println("Estado incorrecto de respuesta")
		return
	}
	defer respuesta.Body.Close()

	var users []datos.User
	json.NewDecoder(respuesta.Body).Decode(&users)

	views.ListUsers(users).Render(r.Context(), w)
}

func ListProducts(w http.ResponseWriter, r *http.Request) {
	url_api := "https://fakestoreapi.com/products"
	respuesta, err := http.DefaultClient.Get(url_api)
	if err != nil {
		fmt.Println("Error al obtener respuesta")
		return
	}
	if respuesta.StatusCode != 200 {
		fmt.Println("Estado incorrecto de respuesta")
		return
	}
	defer respuesta.Body.Close()

	var products []datos.Product
	json.NewDecoder(respuesta.Body).Decode(&products)

	views.ListProducts(products).Render(r.Context(), w)
}
