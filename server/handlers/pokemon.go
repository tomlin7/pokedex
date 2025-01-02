package handlers

import (
	"context"
	"net/http"
	"server/config"
	"server/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetAllPokemon(c *gin.Context) {
	collection := config.DB.Collection("pokemon")
	ctx := context.Background()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch pokemon"})
		return
	}
	defer cursor.Close(ctx)

	var pokemons []models.Pokemon
	if err = cursor.All(ctx, &pokemons); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode pokemon"})
		return
	}

	c.JSON(http.StatusOK, pokemons)
}

func GetPokemonByID(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	collection := config.DB.Collection("pokemon")
	ctx := context.Background()

	var pokemon models.Pokemon
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&pokemon)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pokemon not found"})
		return
	}

	c.JSON(http.StatusOK, pokemon)
}

func SearchPokemon(c *gin.Context) {
	query := strings.ToLower(c.Query("q"))
	collection := config.DB.Collection("pokemon")
	ctx := context.Background()

	filter := bson.M{
		"$or": []bson.M{
			{"name": bson.M{"$regex": query, "$options": "i"}},
			{"number": bson.M{"$eq": parseNumber(query)}},
		},
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Search failed"})
		return
	}
	defer cursor.Close(ctx)

	var pokemons []models.Pokemon
	if err = cursor.All(ctx, &pokemons); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode results"})
		return
	}

	c.JSON(http.StatusOK, pokemons)
}

func parseNumber(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return num
}
