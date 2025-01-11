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

func CreatePokemon(c *gin.Context) {
	var pokemon models.Pokemon
	if err := c.ShouldBindJSON(&pokemon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := config.DB.Collection("pokemon")
	ctx := context.Background()

	result, err := collection.InsertOne(ctx, pokemon)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create pokemon"})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func UpdatePokemon(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var pokemon models.Pokemon
	if err := c.ShouldBindJSON(&pokemon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := config.DB.Collection("pokemon")
	ctx := context.Background()

	update := bson.M{
		"$set": bson.M{
			"name":   pokemon.Name,
			"number": pokemon.Number,
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update pokemon"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pokemon updated"})
}

func DeletePokemon(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	collection := config.DB.Collection("pokemon")
	ctx := context.Background()

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete pokemon"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pokemon deleted"})
}
