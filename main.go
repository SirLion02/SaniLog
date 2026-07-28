package main

import (
	"log"
	"net/http"
	"sanilog/handler"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	fileServer := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	r.Get("/login", handler.LoginPageHandler)
	r.Post("/login-submit", handler.LoginSubmitHandler)
	r.Get("/dashboard", handler.DashboardPageHandler)

	log.Println("Server jalan di :3000")
	http.ListenAndServe(":3000", r)
}
