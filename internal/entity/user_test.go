package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewUser(t *testing.T) {
	user := NewUser("daniel", "vavgenczak", "daniel@gmail.com","1234546")	
	assert.NotNil(t, user)
	assert.NotNil(t, user.ID)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "daniel", user.Firstname)
	assert.Equal(t, "vavgenczak", user.Lastname)
	assert.Equal(t, "daniel@gmail.com", user.Email)
	assert.NotEmpty(t, user.CreatedAt)
	assert.NotEqual(t, "123456", user.Password)
	assert.NotEmpty(t,user.Password)
}

func Test_hashPassword(t *testing.T) {
		password := "123456"
		expectedPassword, err := hashPassword(password) 
		assert.Nil(t, err)
		assert.NotEqual(t, password, expectedPassword)
}