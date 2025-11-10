package testsvc

import (
	"testing"

	"github.com/stretchr/testify/assert"
) 

func TestCreateUser_Succeeds(t *testing.T) { 
	repo := NewInMemoryRepo()   
	service := NewUserService(repo) 

	u, ok := service.CreateUser("Sara", 28) 
	assert.True(t, ok)                   
	assert.NotZero(t, u.ID)            
	assert.Equal(t, "Sara", u.Name)     
	assert.Equal(t, 28, u.Age)         
}

func TestListUsers_ReturnsAll(t *testing.T) {
	repo := NewInMemoryRepo()   
	service := NewUserService(repo) 

	_, _ = service.CreateUser("Yahya", 33)
	_, _ = service.CreateUser("Rasmus", 25)

	users := service.ListUsers()

	assert.Len(t, users,2)
	
}
