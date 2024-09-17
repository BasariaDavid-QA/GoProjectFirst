package handlers

import (
	"context"
	"errors"
	"new-go-project/internal/messagesService" // Импортируем наш сервис
	"new-go-project/internal/web1/messages"   // Импортируем пакет messages
)

type MessageHandler struct {
	Service *messagesService.MessageService
}

// Нужна для создания структуры Handler на этапе инициализации приложения
func NewMessageHandler(service *messagesService.MessageService) *MessageHandler {
	return &MessageHandler{
		Service: service,
	}
}

func (h *MessageHandler) GetMessages(_ context.Context, _ messages.GetMessagesRequestObject) (messages.GetMessagesResponseObject, error) {
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
			Message: &msg.Message,
		}
		response = append(response, message)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

func (h *MessageHandler) PostMessages(_ context.Context, request messages.PostMessagesRequestObject) (messages.PostMessagesResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	messageRequest := request.Body
	// Обращаемся к сервису и создаем сообщение
	messageToCreate := messagesService.Message{Message: *messageRequest.Message}
	createdMessage, err := h.Service.CreateMessage(messageToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := messages.PostMessages201JSONResponse{
		Id:      &createdMessage.ID,
		Message: &createdMessage.Message,
	}
	// Просто возвращаем респонс!
	return response, nil
}

func (h *MessageHandler) PatchMessages(_ context.Context, request messages.PatchMessagesRequestObject) (messages.PatchMessagesResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	messageRequest := request.Body
	// Обращаемся к сервису и обновляем сообщение
	messageToUpdate := messagesService.Message{Message: *messageRequest.Message}
	updatedMessage, err := h.Service.UpdateMessageByID(*messageRequest.Id, messageToUpdate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := messages.PatchMessages200JSONResponse{
		Id:      &updatedMessage.ID,
		Message: &updatedMessage.Message,
	}
	// Просто возвращаем респонс!
	return response, nil
}

func (h *MessageHandler) DeleteMessages(_ context.Context, request messages.DeleteMessagesRequestObject) (messages.DeleteMessagesResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	messageRequest := request.Body
	if messageRequest == nil {
		return nil, errors.New("messageRequest is nil")
	}

	// Обращаемся к сервису и удаляем сообщение
	deletedMessage, err := h.Service.DeleteMessageByID(*messageRequest.Id)

	if err != nil {
		return nil, err
	}

	// создаем структуру респонс
	response := messages.DeleteMessages200JSONResponse{
		Id:      &deletedMessage.ID,
		Message: &deletedMessage.Message,
	}
	// Просто возвращаем респонс!
	return response, nil
}
