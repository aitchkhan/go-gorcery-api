package routes

import "github.com/aitchkhan/go-gorcery-api/controllers"

func RegisterUserRoutes()  {
	s := Routes.PathPrefix("/api/me").Subrouter()
	s.HandleFunc("/login", controllers.LoginHandler).Methods("POST")
	s.HandleFunc("/signup", controllers.SignUpHandler).Methods("POST")
	s.HandleFunc("/orders", controllers.CreateUserOrders).Methods("POST")
	s.HandleFunc("/orders", controllers.GetUserOrders).Methods("GET")
	s.HandleFunc("/cart", controllers.GetUserCart).Methods("GET")
}