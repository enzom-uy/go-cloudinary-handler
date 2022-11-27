package main

import (
	"go_cloudinary/src/api/v1/upload"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	Router := chi.NewRouter()
	Router.Use(middleware.Logger)
	Router.Post("/", upload.Upload)
	http.ListenAndServe(":2707", Router)
}
