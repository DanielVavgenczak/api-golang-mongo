package entity

import (
	"time"

	"github.com/google/uuid"
)

type Movie struct {
	ID uuid.UUID `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
	Year string `json:"year" bson:"year"`
	Director *Director `json:"directors" bson:"directors"`
	Writer *Writer `json:"writers" bson:"writers"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

func NewMovie(name string, year string, director Director, writer Writer)*Movie{
	return &Movie{
		ID: uuid.New(),
		Name: name,
		Year: year,
		Director: &director,
		Writer: &writer,
		CreatedAt: time.Now(),
	}
}

