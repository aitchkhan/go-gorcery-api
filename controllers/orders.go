package controllers

import (
	"fmt"
	"net/http"
)

func GetUserOrders(w http.ResponseWriter, r *http.Request)  {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Get user orders")
}

func CreateUserOrders(w http.ResponseWriter, r *http.Request) {
	
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Create user orders")
}

func GetUserCart(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Get user cart")
}