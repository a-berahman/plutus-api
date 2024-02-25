package repository

import (
	"context"

	"github.com/a-berahman/plutus-api/internal/models"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetActiveUserByID(ctx context.Context, id uint) (*models.User, error)
	UpdateUser(ctx context.Context, id uint, updates map[string]interface{}) error
	DeleteUser(ctx context.Context, id uint) error
}

type TransactionRepositoryInterface interface {
	CreateTransaction(ctx context.Context, tx *models.Transaction) error
	GetTransactionByID(ctx context.Context, id uint) (*models.Transaction, error)
	GetAllTransactions(ctx context.Context, page, pageSize int) ([]*models.Transaction, error)
	UpdateTransaction(ctx context.Context, id uint, updates map[string]interface{}) error
}

type Reoository struct {
	UserRepository        UserRepositoryInterface
	TransactionRepository TransactionRepositoryInterface
}

// NewDB creates and returns instances of all database related repositories
func NewDB(db *gorm.DB) *Reoository {
	return &Reoository{
		UserRepository:        newUserRepository(db),
		TransactionRepository: newTransactionRepository(db),
	}
}
