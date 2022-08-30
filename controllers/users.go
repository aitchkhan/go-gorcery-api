package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aitchkhan/go-gorcery-api/models"
)


func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		panic("Failed to read request body")
	}

	models.CreateUser(models.DB, &u)
	w.WriteHeader(http.StatusOK)
	
	fmt.Fprintf(w, "User Login: %v\n", "Haroon was here")
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User Signup: %v\n", "Haroon was here")
}