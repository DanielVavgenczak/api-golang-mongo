package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewWriter(t *testing.T) {
	expected := "Anne Rice"
	writer := NewWriter("Anne Rice")
	assert.NotEmpty(t, writer.ID)
	assert.NotNil(t, writer.ID)
	assert.NotNil(t, writer.Name)
	assert.Equal(t, expected, writer.Name)
	assert.NotEmpty(t, writer.Name)
}