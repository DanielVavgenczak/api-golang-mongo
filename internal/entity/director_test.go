package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewDirector(t *testing.T) {
	director := NewDirector("Tim Burton")
	assert.Equal(t,"Tim Burton",director.Name)
	assert.NotNil(t, director.ID)
	assert.NotEmpty(t, director.ID)
	assert.NotNil(t, director.Name)
}