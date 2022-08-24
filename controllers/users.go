package controllers

import (
	"fmt"
	"net/http"

	"github.com/aitchkhan/go-gorcery-api/models"
)


func LoginHandler(w http.ResponseWriter, r *http.Request) {
	user := models.User{
		Name: "Haroon Khan",
		Email: "aitchkhan@gmail.com",
	}

	models.CreateUser(models.DB, &user)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User Login: %v\n", "Haroon was here")
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User Signup: %v\n", "Haroon was here")
}