package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/a-berahman/plutus-api/internal/models"
	"github.com/a-berahman/plutus-api/internal/repository"
	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	Repo repository.TransactionRepositoryInterface
}

func newTransactionHandler(repo repository.TransactionRepositoryInterface) *TransactionHandler {
	return &TransactionHandler{Repo: repo}
}
func (t *TransactionHandler) CreateTransactionHandler(c echo.Context) error {
	var transactionRQ models.TransactionCreateRequest
	if err := c.Bind(&transactionRQ); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}
	if err := c.Validate(transactionRQ); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	newTransaction := &models.Transaction{
		UserID: transactionRQ.UserID,
		Amount: transactionRQ.Amount,
		Type:   models.TransactionType(transactionRQ.Type),
	}

	if err := t.Repo.CreateTransaction(c.Request().Context(), newTransaction); err != nil {
		return c.JSON(http.StatusInternalServerError, "could not create transaction")
	}

	transactionResp := models.TransactionResponse{
		ID:     newTransaction.ID,
		UserID: newTransaction.UserID,
		Amount: newTransaction.Amount,
		Type:   string(newTransaction.Type),
	}

	return c.JSON(http.StatusCreated, transactionResp)
}
func (t *TransactionHandler) GetTransactionByIDHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid transaction ID")
	}

	transaction, err := t.Repo.GetTransactionByID(c.Request().Context(), uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "could not retrieve transaction")
	}
	if transaction == nil {
		return c.JSON(http.StatusNotFound, "transaction not found")
	}

	transactionResp := models.TransactionResponse{
		ID:     transaction.ID,
		UserID: transaction.UserID,
		Amount: transaction.Amount,
		Type:   string(transaction.Type),
	}

	return c.JSON(http.StatusOK, transactionResp)
}

func (t *TransactionHandler) GetAllTransactionsHandler(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page < 1 {
		page = 1
	}
	pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))
	if pageSize <= 0 {
		pageSize = 10
	}

	repoTransactions, err := t.Repo.GetAllTransactions(c.Request().Context(), page, pageSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "could not retrieve transactions")
	}

	transactions := make([]models.TransactionResponse, 0, len(repoTransactions))
	for _, repoTransaction := range repoTransactions {
		transactions = append(transactions, models.TransactionResponse{
			ID:     repoTransaction.ID,
			UserID: repoTransaction.UserID,
			Amount: repoTransaction.Amount,
			Type:   string(repoTransaction.Type),
		})
	}

	return c.JSON(http.StatusOK, transactions)
}

func (t *TransactionHandler) UpdateTransactionHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var transactionRQ models.TransactionUpdateRequest
	if err := c.Bind(&transactionRQ); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}
	if err := c.Validate(transactionRQ); err != nil {
		return c.JSON(http.StatusBadRequest, "validation error: "+err.Error())
	}

	updates := make(map[string]interface{})
	if transactionRQ.Amount != 0 {
		updates["amount"] = transactionRQ.Amount
	}
	if strings.TrimSpace(transactionRQ.Type) != "" {
		updates["type"] = models.TransactionType(transactionRQ.Type)
	}

	if err := t.Repo.UpdateTransaction(c.Request().Context(), uint(id), updates); err != nil {
		return c.JSON(http.StatusInternalServerError, "could not update transaction")
	}

	return c.JSON(http.StatusOK, "transaction updated successfully")
}
