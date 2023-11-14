package repository

import (
	"context"

	"github.com/DanielVavgenczak/api-golang-mongo/internal/entity"
)


type ContractsMovie interface {
	Create(ctx context.Context, movie entity.Movie) (*entity.Movie, error)
	FindByName(ctx context.Context, name string)*entity.Movie

	//All() []entity.Movie
}