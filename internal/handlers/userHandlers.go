package handlers

import (
	"context"
	"errors"
	"new-go-project/internal/userService" // Импортируем наш сервис
	"new-go-project/internal/web2/users"  // Импортируем пакет users
)

type UserHandler struct {
	Service *userService.UserService
}

// Нужна для создания структуры UserHandler на этапе инициализации приложения
func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

func (h *UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	// Получение всех пользователей из сервиса
	allUsers, err := h.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	// Создаем переменную респон типа 200джейсонРеспонс
	// Которую мы потом передадим в качестве ответа
	response := users.GetUsers200JSONResponse{}

	// Заполняем слайс response всеми пользователями из БД
	for _, usr := range allUsers {
		user := users.User{
			Id:       &usr.ID,
			Email:    &usr.Email,
			Password: &usr.Password,
		}
		response = append(response, user)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

func (h *UserHandler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	userRequest := request.Body
	// Обращаемся к сервису и создаем пользователя
	userToCreate := userService.User{Email: *userRequest.Email, Password: *userRequest.Password}
	createdUser, err := h.Service.CreateUser(userToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := users.PostUsers201JSONResponse{
		Id:       &createdUser.ID,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}
	// Просто возвращаем респонс!
	return response, nil
}

func (h *UserHandler) PatchUsers(_ context.Context, request users.PatchUsersRequestObject) (users.PatchUsersResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	userRequest := request.Body
	// Обращаемся к сервису и обновляем пользователя
	userToUpdate := userService.User{Email: *userRequest.Email, Password: *userRequest.Password}
	updatedUser, err := h.Service.UpdateUserByID(*userRequest.Id, userToUpdate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := users.PatchUsers200JSONResponse{
		Id:       &updatedUser.ID,
		Email:    &updatedUser.Email,
		Password: &updatedUser.Password,
	}
	// Просто возвращаем респонс!
	return response, nil
}

func (h *UserHandler) DeleteUsers(_ context.Context, request users.DeleteUsersRequestObject) (users.DeleteUsersResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	userRequest := request.Body
	if userRequest == nil {
		return nil, errors.New("userRequest is nil")
	}

	// Обращаемся к сервису и удаляем пользователя
	deletedUser, err := h.Service.DeleteUserByID(*userRequest.Id)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := users.DeleteUsers200JSONResponse{
		Id:       &deletedUser.ID,
		Email:    &deletedUser.Email,
		Password: &deletedUser.Password,
	}
	// Просто возвращаем респонс!
	return response, nil
}
