package main

import (
	"context"
	"fmt"
	"time"

	"github.com/DanielVavgenczak/api-golang-mongo/internal/infra/database"
	"github.com/gin-gonic/gin"
)

var monddbIurl = "mongodb://localhost:27017"
func main() {
	r := gin.Default()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := database.NewMongoConnection(ctx, monddbIurl)
	
	if err != nil {
		panic(err)
	}
	
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message":"ok",
		})
	})
	fmt.Println("Build...")

	r.Run(":3000")

}