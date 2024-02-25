package handlers

import (
	"github.com/a-berahman/plutus-api/internal/repository"
	"github.com/labstack/echo/v4"
)

type UserHandlerInterface interface {
	CreateUserHandler(c echo.Context) error
	GetUserByIDHandler(c echo.Context) error
	UpdateUserHandler(c echo.Context) error
	DeleteUserHandler(c echo.Context) error
}
type TransactionHandlerInterface interface {
	CreateTransactionHandler(c echo.Context) error
	GetTransactionByIDHandler(c echo.Context) error
	GetAllTransactionsHandler(c echo.Context) error
	UpdateTransactionHandler(c echo.Context) error
}
type Handler struct {
	UserHandler        UserHandlerInterface
	TransactionHandler TransactionHandlerInterface
}

// NewHandler creates a new instance of the handler
func NewHandler(repo *repository.Reoository) *Handler {
	return &Handler{
		UserHandler:        newUserHandler(repo.UserRepository),
		TransactionHandler: newTransactionHandler(repo.TransactionRepository),
	}
}
