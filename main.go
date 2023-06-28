package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {

		portString := os.Getenv("PORT")

		if portString == "" {
			log.Fatal("PORT is not found in the environment")
		}

		router := chi.NewRouter()
		router.Use(middleware.Logger)
		router.Use(cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300,
		}))

		v1Router := chi.NewRouter()
		v1Router.Get("/healthz", handlerReadiness)
		v1Router.Get("/err", handlerErr)

		router.Mount("/v1", v1Router)

		server := &http.Server{
			Handler: router,
			Addr:    ":" + portString,
		}

		err := server.ListenAndServe()

		if err != nil {
			log.Fatal(err)
		}
	}

}
