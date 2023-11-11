package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewMovie(t *testing.T) {
	director := NewDirector("George Lucas")
	writer := NewWriter("George Lucas")
	movie := NewMovie("Star Wars", "1977", *director,*writer)

	assert.Equal(t, movie.Director, director)
	assert.Equal(t, movie.Writer, writer)
	assert.NotEmpty(t, movie.ID)
	assert.NotEmpty(t, movie.Name)
	assert.Equal(t, "Star Wars", movie.Name)
}