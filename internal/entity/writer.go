package entity

import "github.com/google/uuid"

type Writer struct {
	ID uuid.UUID `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}


func NewWriter(name string) *Writer {
	return &Writer{
		ID: uuid.New(),
		Name: name,
	}
}

