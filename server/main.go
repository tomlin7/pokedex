package main

import (
	"log"
	"server/config"
	"server/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// ${BASE_URL}/api/pokemon
// ${BASE_URL}/api/pokemon/1
// ${BASE_URL}/api/pokemon/search?q=bulbasaur

func main() {
	config.ConnectDB()

	r := gin.Default()
	r.Use(cors.Default())

	// routes
	api := r.Group("/api")
	api.Use(cors.Default())
	{
		api.GET("/pokemon", handlers.GetAllPokemon)
		api.GET("/pokemon/:id", handlers.GetPokemonByID)
		api.GET("/pokemon/search", handlers.SearchPokemon)
	}

	log.Fatal(r.Run(":8080"))
}
