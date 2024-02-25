package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/a-berahman/plutus-api/internal/models"
	"github.com/a-berahman/plutus-api/internal/repository/mocks"
	"go.uber.org/mock/gomock"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestTransactionHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTransactionRepositoryInterface(ctrl)
	handler := newTransactionHandler(mockRepo)

	tests := []struct {
		name        string
		setupMock   func()
		request     func() *http.Request
		handlerFunc func(echo.Context) error
		statusCode  int
	}{
		{
			name: "CreateTransaction_Success",
			setupMock: func() {
				mockRepo.EXPECT().CreateTransaction(gomock.Any(), gomock.Any()).Return(nil)
			},
			request: func() *http.Request {
				transaction := models.TransactionCreateRequest{
					Amount:   "100.00",
					Type:     "credit",
					UserID:   1,
					Currency: "usd",
				}
				jsonBytes, _ := json.Marshal(transaction)
				req := httptest.NewRequest(http.MethodPost, "/transactions", bytes.NewBuffer(jsonBytes))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				return req
			},
			handlerFunc: handler.CreateTransactionHandler,
			statusCode:  http.StatusCreated,
		},
		// i keep this simple for the sake of the example for interview
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMock()

			req := tc.request()
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if assert.NoError(t, tc.handlerFunc(c)) {
				assert.Equal(t, tc.statusCode, rec.Code)
			}
		})
	}
}
