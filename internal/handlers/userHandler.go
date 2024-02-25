package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/a-berahman/plutus-api/internal/models"
	"github.com/a-berahman/plutus-api/internal/repository"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	Repo repository.UserRepositoryInterface
}

func newUserHandler(repo repository.UserRepositoryInterface) *UserHandler {
	return &UserHandler{Repo: repo}
}

func (u *UserHandler) CreateUserHandler(c echo.Context) error {
	var userRQ models.UserCreateRequest
	if err := c.Bind(&userRQ); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}
	if err := c.Validate(userRQ); err != nil {
		return c.JSON(http.StatusBadRequest, "validation error: "+err.Error())
	}
	newUser := &models.User{
		Name:  userRQ.Name,
		Email: userRQ.Email,
	}
	if err := u.Repo.CreateUser(c.Request().Context(), newUser); err != nil {
		return c.JSON(http.StatusInternalServerError, "could not create user")
	}
	userResp := models.UserResponse{
		ID:        newUser.ID,
		Name:      newUser.Name,
		Email:     newUser.Email,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
	}
	return c.JSON(http.StatusCreated, userResp)
}

func (u *UserHandler) GetUserByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	currUser, err := u.Repo.GetActiveUserByID(c.Request().Context(), uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "could not retrieve user")
	}
	if currUser == nil {
		return c.JSON(http.StatusNotFound, "user not found")
	}
	userResp := models.UserResponse{
		ID:        currUser.ID,
		Name:      currUser.Name,
		Email:     currUser.Email,
		CreatedAt: currUser.CreatedAt,
		UpdatedAt: currUser.UpdatedAt,
	}
	return c.JSON(http.StatusOK, userResp)
}
func (u *UserHandler) UpdateUserHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid user ID")
	}

	var userRQ models.UserUpdateRequest
	if err := c.Bind(&userRQ); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}
	if err := c.Validate(userRQ); err != nil {
		return c.JSON(http.StatusBadRequest, "validation error: "+err.Error())
	}

	updates := make(map[string]interface{})
	if strings.TrimSpace(userRQ.Name) != "" {
		updates["name"] = userRQ.Name
	}
	if strings.TrimSpace(userRQ.Email) != "" {
		updates["email"] = userRQ.Email
	}

	if err := u.Repo.UpdateUser(c.Request().Context(), uint(id), updates); err != nil {
		return c.JSON(http.StatusInternalServerError, "could not update user")
	}

	return c.JSON(http.StatusOK, "user updated")
}
func (u *UserHandler) DeleteUserHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := u.Repo.DeleteUser(c.Request().Context(), uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, "could not delete user")
	}

	return c.JSON(http.StatusOK, "user deleted")
}
