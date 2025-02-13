package services_test

import (
	"controle/financeiro/domain/services"
	"controle/financeiro/infra/repositories/memory"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserService(t *testing.T) {
	userRepo := memory.NewInMemoryUserRepository()
	service := services.NewUserService(userRepo)

	t.Run("Create User", func(t *testing.T) {
		user, err := service.CreateUser("testuser", 5000.0)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, "testuser", user.Username)
	})

	t.Run("Get User by ID", func(t *testing.T) {
		user, _ := service.CreateUser("testuser2", 4000.0)
		fetchedUser, err := service.GetUserByID(user.ID)
		assert.NoError(t, err)
		assert.Equal(t, user.Username, fetchedUser.Username)
	})

	t.Run("Update User", func(t *testing.T) {
		user, _ := service.CreateUser("testuser3", 3000.0)
		err := service.UpdateUser(user.ID, "updateduser", 3500.0)
		assert.NoError(t, err)

		updatedUser, _ := service.GetUserByID(user.ID)
		assert.Equal(t, "updateduser", updatedUser.Username)
		assert.Equal(t, 3500.0, updatedUser.Income)
	})

	t.Run("Delete User", func(t *testing.T) {
		user, _ := service.CreateUser("testuser4", 2000.0)
		err := service.DeleteUser(user.ID)
		assert.NoError(t, err)

		deletedUser, err := service.GetUserByID(user.ID)
		assert.Error(t, err)
		assert.Nil(t, deletedUser)
	})
}
