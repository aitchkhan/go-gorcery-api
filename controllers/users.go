package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aitchkhan/go-gorcery-api/models"
	"github.com/aitchkhan/go-gorcery-api/types"
	"github.com/davecgh/go-spew/spew"
	"golang.org/x/crypto/bcrypt"
)
type Response struct{
	Message string `json:"message"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var u types.User
	var dbUser models.User
	spew.Dump("in controller", u.Email)
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		panic("Failed to read request body")
	}

	dbUser.Email = u.Email
	models.GetUserByEmail(models.DB, &dbUser)
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.HashedPassword), []byte(u.Password))
  if err != nil {
	w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		r := Response{
			Message: "Password didnt match",
		}
		rJSON, err := json.Marshal(r)
		if err != nil {
			log.Fatal("error marshalling")
		}
		w.Write(rJSON)
		return
}
	if dbUser.ID == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		r := Response{
			Message: "failed to find user",
		}
		rJSON, err := json.Marshal(r)
		if err != nil {
			log.Fatal("error marshalling")
		}
		w.Write(rJSON)
		return
	}

	userJSON, err := json.Marshal(dbUser)
	if err != nil {
		panic("Failed to find user by email")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var u types.User
	var dbUser models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		panic("Failed to read request body")
	}
	spew.Dump(u)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Failed to bcrypt")
	}

	dbUser = models.User{
		Name: u.Name,
		Email: u.Email,
	  HashedPassword: string(hashedPassword),
	}
	models.CreateUser(models.DB, &dbUser)
	s := Response{
		Message: "User logged in successfully",
	}
	successJSON, err := json.Marshal(s)
	if err != nil {
		log.Fatal("Failed to marshal s")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(successJSON)
}