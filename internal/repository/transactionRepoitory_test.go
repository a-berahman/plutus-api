package repository

import (
	"context"
	"testing"

	"github.com/a-berahman/plutus-api/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestTransactionRepository(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	db.AutoMigrate(&models.Transaction{})

	repo := newTransactionRepository(db)

	tests := []struct {
		name string
		run  func(t *testing.T)
	}{
		{
			name: "CreateTransaction",
			run: func(t *testing.T) {
				tx := models.Transaction{
					UserID: 1,
					Amount: 1000,
					Type:   models.Credit,
				}
				err := repo.CreateTransaction(context.Background(), &tx)
				assert.NoError(t, err)
				assert.NotZero(t, tx.ID)
			},
		},
		{
			name: "GetTransactionByID",
			run: func(t *testing.T) {
				tx := models.Transaction{
					UserID: 1,
					Amount: 1000,
					Type:   models.Credit,
				}
				db.Create(&tx)

				foundTx, err := repo.GetTransactionByID(context.Background(), tx.ID)
				assert.NoError(t, err)
				assert.NotNil(t, foundTx)
				assert.Equal(t, tx.Amount, foundTx.Amount)
				assert.Equal(t, tx.Type, foundTx.Type)
			},
		},
		// We can cover more scenarios but because it's an interview challenge keep it simple
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.run)
	}
}
