package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/a-berahman/plutus-api/internal/models"
	"github.com/a-berahman/plutus-api/internal/repository/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUserHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoryInterface(ctrl)
	handler := newUserHandler(mockRepo)

	tests := []struct {
		name        string
		setupMock   func()
		request     func() *http.Request
		handlerFunc func(echo.Context) error
		statusCode  int
	}{
		{
			name: "CreateUser_Success",
			setupMock: func() {
				mockRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(nil)
			},
			request: func() *http.Request {
				user := models.UserCreateRequest{Name: "ahmadb", Email: "ahmad.berahman@hotmail.com"}
				jsonBytes, _ := json.Marshal(user)
				req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonBytes))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				return req
			},
			handlerFunc: handler.CreateUserHandler,
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
