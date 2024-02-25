package handlers

import (
	"fmt"
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

	amountInCents, err := convertToCents(transactionRQ.Amount)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid amount format")
	}

	newTransaction := &models.Transaction{
		UserID:   transactionRQ.UserID,
		Amount:   amountInCents,
		Currency: models.TransactionCurrency(transactionRQ.Currency),
		Type:     models.TransactionType(transactionRQ.Type),
	}

	if err := t.Repo.CreateTransaction(c.Request().Context(), newTransaction); err != nil {
		return c.JSON(http.StatusInternalServerError, "could not create transaction")
	}

	transactionResp := models.TransactionResponse{
		ID:       newTransaction.ID,
		UserID:   newTransaction.UserID,
		Amount:   transactionRQ.Amount,
		Currency: string(newTransaction.Currency),
		Type:     string(newTransaction.Type),
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

	return c.JSON(http.StatusOK, models.TransactionResponse{
		ID:       transaction.ID,
		UserID:   transaction.UserID,
		Amount:   convertCentsToString(transaction.Amount),
		Currency: string(transaction.Currency),
		Type:     string(transaction.Type),
	})
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
		amountStr := convertCentsToString(repoTransaction.Amount)
		transactions = append(transactions, models.TransactionResponse{
			ID:       repoTransaction.ID,
			UserID:   repoTransaction.UserID,
			Amount:   amountStr,
			Currency: string(repoTransaction.Currency),
			Type:     string(repoTransaction.Type),
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
	if transactionRQ.Amount != "" {
		amountInCents, err := convertToCents(transactionRQ.Amount)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "invalid amount format")
		}
		updates["amount"] = amountInCents
	}
	if strings.TrimSpace(transactionRQ.Type) != "" {
		updates["type"] = models.TransactionType(transactionRQ.Type)
	}
	if strings.TrimSpace(transactionRQ.Currency) != "" {
		updates["currency"] = models.TransactionCurrency(transactionRQ.Currency)
	}

	if err := t.Repo.UpdateTransaction(c.Request().Context(), uint(id), updates); err != nil {
		return c.JSON(http.StatusInternalServerError, "could not update transaction")
	}

	return c.JSON(http.StatusOK, "transaction updated successfully")
}

func convertToCents(amountStr string) (int64, error) {
	parts := strings.Split(amountStr, ".")
	var majorUnits, minorUnits int64
	var err error
	if majorUnits, err = strconv.ParseInt(parts[0], 10, 64); err != nil {
		return 0, err
	}
	majorUnits *= 100

	if len(parts) == 2 {
		if len(parts[1]) > 2 {
			parts[1] = parts[1][:2]
		}
		if minorUnits, err = strconv.ParseInt(parts[1], 10, 64); err != nil {
			return 0, err
		}
		if len(parts[1]) == 1 {
			minorUnits *= 10
		}
	}

	return majorUnits + minorUnits, nil
}

func convertCentsToString(amount int64) string {
	majorUnits := amount / 100
	minorUnits := amount % 100
	return fmt.Sprintf("%d.%02d", majorUnits, minorUnits)
}
