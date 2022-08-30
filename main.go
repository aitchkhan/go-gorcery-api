package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aitchkhan/go-gorcery-api/models"
	"github.com/aitchkhan/go-gorcery-api/routes"
)

func main() {

	models.OpenDBConn()
	err := models.DB.AutoMigrate(&models.User{})
	if err !=nil {
		panic("failed on migrating")
	}

	routes.RegisterUserRoutes()
	http.Handle("/", routes.Routes)

	srv := &http.Server{
		Handler:      routes.Routes,
		Addr:         "0.0.0.0:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Home: %v\n", "Haroon was here")
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Product Handle:\n")
}

func OrdersHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Product Handle:\n")
}
