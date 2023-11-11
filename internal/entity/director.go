package entity

import (
	"github.com/google/uuid"
)

type Director struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

func NewDirector(name string )*Director {
	return &Director{
		ID: uuid.NewString(),
		Name: name,
	}
}


