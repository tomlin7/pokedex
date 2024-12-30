package main

import (
	"log"
	"server/handlers"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

// ${BASE_URL}/api/pokemon
// ${BASE_URL}/api/pokemon/1
// ${BASE_URL}/api/pokemon/search?q=bulbasaur

func main() {
	r := gin.Default()

	r.Use(middleware.CORS())

	// routes
	api := r.Group("/api")
	{
		api.GET("/pokemon", handlers.GetAllPokemon)
		api.GET("/pokemon/:id", handlers.GetPokemonByID)
		api.GET("/pokemon/search", handlers.SearchPokemon)
	}

	log.Fatal(r.Run(":8080"))
}
