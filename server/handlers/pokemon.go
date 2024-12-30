package handlers

import (
	"net/http"
	"server/models"

	"github.com/gin-gonic/gin"
)

func GetAllPokemon(c *gin.Context) {
	c.JSON(http.StatusOK, models.PokemonList)
}
