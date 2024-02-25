package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/a-berahman/plutus-api/internal/handlers"
	"github.com/a-berahman/plutus-api/internal/models"
	"github.com/a-berahman/plutus-api/internal/repository"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	e := setupApiInstance()
	db := openDatabaseConnection(e)

	repo := repository.NewDB(db)
	handler := handlers.NewHandler(repo)

	setupRoutes(e, handler)

	go startServer(e)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func setupRoutes(e *echo.Echo, handler *handlers.Handler) {
	e.POST("/api/v1/users", handler.UserHandler.CreateUserHandler)
	e.GET("/api/v1/users/:id", handler.UserHandler.GetUserByIDHandler)
	e.PUT("/api/v1/users/:id", handler.UserHandler.UpdateUserHandler)
	e.DELETE("/api/v1/users/:id", handler.UserHandler.DeleteUserHandler)

	e.POST("/api/v1/transactions", handler.TransactionHandler.CreateTransactionHandler)
	e.GET("/api/v1/transactions/:id", handler.TransactionHandler.GetTransactionByIDHandler)
	e.GET("/api/v1/transactions", handler.TransactionHandler.GetAllTransactionsHandler)
	e.PUT("/api/v1/transactions/:id", handler.TransactionHandler.UpdateTransactionHandler)
}

func startServer(e *echo.Echo) {
	if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
		e.Logger.Fatal("shutting down the server")
	}
}

func openDatabaseConnection(e *echo.Echo) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("plutus.db"), &gorm.Config{})
	if err != nil {
		e.Logger.Fatal("Error connecting to database: ", err)
	}
	if err = db.AutoMigrate(&models.User{}, &models.Transaction{}); err != nil {
		e.Logger.Fatal("Error auto-migrating database: ", err)
	}
	return db
}

type Validator struct {
	validator *validator.Validate
}

func setupApiInstance() *echo.Echo {
	e := echo.New()
	e.Validator = &Validator{validator: validator.New()}

	e.Use(middleware.Logger(), middleware.Recover(), middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		},
	))
	return e
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
