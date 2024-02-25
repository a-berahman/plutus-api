package repository

import (
	"context"

	"github.com/a-berahman/plutus-api/internal/models"
	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
	db *gorm.DB
}

func newTransactionRepository(db *gorm.DB) *TransactionRepositoryImpl {
	return &TransactionRepositoryImpl{db: db}
}

func (t *TransactionRepositoryImpl) CreateTransaction(ctx context.Context, tx *models.Transaction) error {
	return t.db.WithContext(ctx).Create(tx).Error
}

func (t *TransactionRepositoryImpl) GetTransactionByID(ctx context.Context, id uint) (*models.Transaction, error) {
	var transaction models.Transaction
	result := t.db.WithContext(ctx).First(&transaction, id)
	return &transaction, result.Error
}

func (t *TransactionRepositoryImpl) GetAllTransactions(ctx context.Context, page, pageSize int) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	offset := (page - 1) * pageSize
	result := t.db.WithContext(ctx).Offset(offset).Limit(pageSize).Find(&transactions)
	return transactions, result.Error
}

func (t *TransactionRepositoryImpl) UpdateTransaction(ctx context.Context, id uint, updates map[string]interface{}) error {
	return t.db.WithContext(ctx).Model(&models.Transaction{}).Where("id = ?", id).Updates(updates).Error
}
