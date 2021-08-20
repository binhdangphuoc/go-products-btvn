package main

import (
	"action"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	//"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func main(){
	router := mux.NewRouter().StrictSlash(true)
	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*",})

	notLogin := router.PathPrefix("/").Subrouter()
	mustAdmin := router.PathPrefix("/admin").Subrouter()

	notLogin.Methods(http.MethodGet).Path("/products").HandlerFunc(action.GetProducts)
	notLogin.Methods(http.MethodGet).Path("/products/{id}").HandlerFunc(action.GetDetailProduct)

	mustAdmin.Use(action.AminCheckingMiddleware)
	mustAdmin.Methods(http.MethodGet).Path("/products").HandlerFunc(action.GetProducts)
	mustAdmin.Methods(http.MethodGet).Path("/products/{id}").HandlerFunc(action.GetDetailProduct)
	mustAdmin.Methods(http.MethodPost).Path("/products").HandlerFunc(action.CreateNewProduct)
	mustAdmin.Methods(http.MethodPut).Path("/products/{id}").HandlerFunc(action.UpdateProduct)
	mustAdmin.Methods(http.MethodDelete).Path("/products/{id}").HandlerFunc(action.DeleteProduct)

	mustAdmin.Methods(http.MethodGet).Path("/orders").HandlerFunc(action.GetOrders)
	mustAdmin.Methods(http.MethodGet).Path("/orders/{id}").HandlerFunc(action.GetOrderDetail)

	err := http.ListenAndServe(":8888", handlers.CORS(headers, methods, origins)(router))
	if err != nil {
		panic(err)
	}

}
