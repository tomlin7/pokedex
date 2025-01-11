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
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*", "https://pokedex-rho-hazel.vercel.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "X-Requested-With"},
		AllowCredentials: true,
	}))

	// routes
	api := r.Group("/api")
	{
		api.GET("/pokemon", handlers.GetAllPokemon)
		api.GET("/pokemon/:id", handlers.GetPokemonByID)
		api.GET("/pokemon/search", handlers.SearchPokemon)
		api.POST("/pokemon", handlers.CreatePokemon)
		api.PUT("/pokemon/:id", handlers.UpdatePokemon)
		api.DELETE("/pokemon/:id", handlers.DeletePokemon)
	}

	log.Fatal(r.Run(":8080"))
}
