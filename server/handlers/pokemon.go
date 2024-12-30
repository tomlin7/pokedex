package handlers

import (
	"net/http"
	"server/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAllPokemon(c *gin.Context) {
	c.JSON(http.StatusOK, models.PokemonList)
}

func GetPokemonByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for _, pokemon := range models.PokemonList {
		if pokemon.ID == id {
			c.JSON(http.StatusOK, pokemon)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Pokemon not found"})
}

func SearchPokemon(c *gin.Context) {
	query := strings.ToLower(c.Query("q"))

	if query == "" {
		c.JSON(http.StatusOK, models.PokemonList)
		return
	}

	var results []models.Pokemon
	for _, pokemon := range models.PokemonList {
		if strings.Contains(strings.ToLower(pokemon.Name), query) ||
			strings.Contains(strconv.Itoa(pokemon.Number), query) {
			results = append(results, pokemon)
		}
	}

	c.JSON(http.StatusOK, results)
}
