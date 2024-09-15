package handlers

import (
	"context"
	"new-go-project/internal/messagesService"
	"new-go-project/internal/web/messages"
)

type Handler struct {
	Service *messagesService.MessageService
}

func NewHandler(service *messagesService.MessageService) *Handler {
	return &Handler{
		Service: service,
	}
}

// Helper function to convert uint to int
func uintToInt(u uint) int {
	return int(u)
}

func (h *Handler) GetMessages(_ context.Context, _ messages.GetMessagesRequestObject) (messages.GetMessagesResponseObject, error) {
	// Получение всех сообщений из сервиса
	allMessages, err := h.Service.GetAllMessages()
	if err != nil {
		return nil, err
	}

	// Создаем переменную респон типа 200джейсонРеспонс
	// Которую мы потом передадим в качестве ответа
	response := messages.GetMessages200JSONResponse{}

	// Заполняем слайс response всеми сообщениями из БД
	for _, msg := range allMessages {
		message := messages.Message{
			Id:      &msg.ID,
			Message: &msg.Text,
		}
		response = append(response, message)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

func (h *Handler) PostMessages(_ context.Context, request messages.PostMessagesRequestObject) (messages.PostMessagesResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	messageRequest := request.Body
	// Обращаемся к сервису и создаем сообщение
	messageToCreate := messagesService.Message{Text: *messageRequest.Message}
	createdMessage, err := h.Service.CreateMessage(messageToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := messages.PostMessages201JSONResponse{
		Id:      &createdMessage.ID,
		Message: &createdMessage.Text,
	}
	// Просто возвращаем респонс!
	return response, nil
}

func (h *Handler) PatchMessages(_ context.Context, request messages.PatchMessagesRequestObject) (messages.PatchMessagesResponseObject, error) {
	id := uintToInt(*request.Body.Id) // Convert *int to uint
	messageUpdate := messagesService.Message{Text: *request.Body.Message}
	updatedMessage, err := h.Service.UpdateMessageByID(id, messageUpdate)
	if err != nil {
		return nil, err
	}
	response := messages.PatchMessages200JSONResponse{
		Id:      &updatedMessage.ID, // Convert *uint to *int
		Message: &updatedMessage.Text,
	}
	return response, nil
}

func (h *Handler) DeleteMessages(_ context.Context, request messages.DeleteMessagesRequestObject) (messages.DeleteMessagesResponseObject, error) {
	id := uint(request.Params.Id) // Convert int to uint
	err := h.Service.DeleteMessageByID(uintToInt(id))
	if err != nil {
		return nil, err
	}

	return &messages.DeleteMessages204Response{}, nil
}
