package main

import (
	"context"
	"time"

	"github.com/DanielVavgenczak/api-golang-mongo/internal/infra/database"
	"github.com/DanielVavgenczak/api-golang-mongo/internal/repository"
	"github.com/DanielVavgenczak/api-golang-mongo/internal/service"
	"github.com/DanielVavgenczak/api-golang-mongo/internal/web/handlers"
	"github.com/gin-gonic/gin"
)

var monddbIurl = "mongodb://mongoadmin:docker@localhost:27017/"
func main() {
	r := gin.Default()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db, err := database.NewMongoConnection(ctx, monddbIurl)
	
	if err != nil {
		panic(err)
	}
	movieRepository := repository.NewRepositoryDB(db.Database())
	movieService := service.NewService(*movieRepository)
	handlerMovie := handlers.NewMovieHandler(movieService)

	r.POST("/movie", handlerMovie.CreateHandler)

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message":"ok",
		})
	})

	
	r.Run(":3000")

}