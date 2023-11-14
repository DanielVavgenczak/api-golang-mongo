package handlers

import (
	"context"
	"net/http"
	"time"

	dto "github.com/DanielVavgenczak/api-golang-mongo/internal/DTO"
	"github.com/DanielVavgenczak/api-golang-mongo/internal/service"
	"github.com/gin-gonic/gin"
)

type MovieHandler struct { 
	MovieHandler *service.ServiceMovie
}

func NewMovieHandler(service *service.ServiceMovie)*MovieHandler {
	return &MovieHandler{
		MovieHandler: service,
	}
}

func(mh *MovieHandler)CreateHandler(ctx *gin.Context) {
	
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var movieInput dto.MovieInput
	
	if err :=  ctx.BindJSON(&movieInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status-code": http.StatusBadRequest,
		})
		return
	}

	created, err := mh.MovieHandler.CreateMovie(context, movieInput)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status-code": http.StatusBadRequest,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": created,
	})
	


}

