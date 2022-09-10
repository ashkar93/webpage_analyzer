package routers

import (
	"example/sample/controllers"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func InitRouter() *chi.Mux {

	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	r.Use(cors.Handler)

	//Health check
	r.Get("/", controllers.SayHelloworld)

	//Webpage Analyze
	r.Get("/api/v1/analyze-webpage", controllers.WebScraper)

	return r

}
