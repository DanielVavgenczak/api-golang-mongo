package service

import (
	"context"
	"errors"
	"fmt"
	"log"

	dto "github.com/DanielVavgenczak/api-golang-mongo/internal/DTO"
	"github.com/DanielVavgenczak/api-golang-mongo/internal/entity"
	"github.com/DanielVavgenczak/api-golang-mongo/internal/repository"
)

var (
	ErrMovieIsAlreadExists = errors.New("movie is alread exists")
)

type ServiceMovie struct {
	Repository repository.ContractsMovie
}

func NewService(rb repository.RepositoryDB) *ServiceMovie{
	return &ServiceMovie{
		Repository: &rb,
	}
}
func (s *ServiceMovie) CreateMovie(ctx context.Context, movie dto.MovieInput)(*entity.Movie, error) {

	director := entity.NewDirector(
		movie.Director,
	)
	writer := entity.NewWriter(movie.Writer)

	create := entity.NewMovie(
		movie.Name,
		movie.Year,
		*director,
		*writer,
	)

	movieExists := s.Repository.FindByName(ctx, create.Name)
	fmt.Println("Encontrado", movieExists.Name)
	if movieExists.Name != ""  {
		log.Println("Tem como", movieExists, len(movieExists.Name))
		return nil, ErrMovieIsAlreadExists
	}

	movieInsert, err := s.Repository.Create(ctx, *create)
	if err != nil {
		return nil, err
	}

	return movieInsert, nil
}
