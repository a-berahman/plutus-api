package repository

import (
	"context"

	"github.com/a-berahman/plutus-api/internal/models"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func newUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (u *UserRepositoryImpl) CreateUser(ctx context.Context, user *models.User) error {
	return u.db.WithContext(ctx).Create(user).Error
}

func (u *UserRepositoryImpl) GetActiveUserByID(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	result := u.db.WithContext(ctx).First(&user, id)
	return &user, result.Error
}

func (u *UserRepositoryImpl) UpdateUser(ctx context.Context, id uint, updates map[string]interface{}) error {
	return u.db.WithContext(ctx).Model(&models.User{}).Where("id = ?", id).Updates(updates).Error
}

func (u *UserRepositoryImpl) DeleteUser(ctx context.Context, id uint) error {
	return u.db.WithContext(ctx).Delete(&models.User{}, id).Error
}
