package handlers

import (
	"net/http"
	"new-go-project/userService"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	Service *userService.UserService
}

func (h *UserHandler) GetUsers(c echo.Context) error {
	users, err := h.Service.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) PostUsers(c echo.Context) error {
	var user userService.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := h.Service.CreateUser(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) PatchUserByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	existingUser, err := h.Service.GetUserByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, "User not found")
	}

	var updateData userService.User
	if err := c.Bind(&updateData); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	updates := make(map[string]interface{})
	if updateData.Email != "" {
		updates["email"] = updateData.Email
	}
	if updateData.Password != "" {
		updates["password"] = updateData.Password
	}

	if err := h.Service.UpdateUserFields(existingUser.ID, updates); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update user")
	}

	// Получаем обновленного пользователя для подтверждения изменений
	updatedUser, err := h.Service.GetUserByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to retrieve updated user")
	}

	return c.JSON(http.StatusOK, updatedUser)
}

func (h *UserHandler) DeleteUserByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}
	if err := h.Service.DeleteUser(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete user")
	}
	return c.NoContent(http.StatusNoContent)
}
