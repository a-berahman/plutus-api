package repository

import (
	"context"
	"testing"

	"github.com/a-berahman/plutus-api/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestUserRepository(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	db.AutoMigrate(&models.User{})

	repo := newUserRepository(db)

	tests := []struct {
		name string
		run  func(t *testing.T)
	}{
		{
			name: "CreateUser",
			run: func(t *testing.T) {
				user := models.User{Name: "Ahmad", Email: "ahmad.berahman@hotmail.com"}
				err := repo.CreateUser(context.Background(), &user)
				assert.NoError(t, err)
				assert.NotZero(t, user.ID)
			},
		},
		{
			name: "GetUserByID",
			run: func(t *testing.T) {
				user := models.User{Name: "Ahmad", Email: "ahmad.berahman@hotmail.com"}
				db.Create(&user)

				foundUser, err := repo.GetActiveUserByID(context.Background(), user.ID)
				assert.NoError(t, err)
				assert.NotNil(t, foundUser)
				assert.Equal(t, user.Name, foundUser.Name)
				assert.Equal(t, user.Email, foundUser.Email)
			},
		},
		// We can cover more scenarios here but because it's an interview problem we keep it simple :)
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.run)
	}
}
