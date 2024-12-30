package main

import (
	"log"
	"server/handlers"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.CORS())

	// routes
	api := r.Group("/api")
	{
		api.GET("/pokemon", handlers.GetAllPokemon)
	}

	log.Fatal(r.Run(":8080"))
}
