package main

import (
	"log"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.CORS())

	log.Fatal(r.Run(":8080"))
}
