package handlers

import (
	"net/http"
	"server/models"
	"strconv"

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
