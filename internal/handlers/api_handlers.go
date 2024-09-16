package handlers

import (
	"net/http"
	"new-go-project/internal/messagesService"
	"new-go-project/internal/web/messages"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MessageHandler struct {
	Service *messagesService.MessageService
}

func NewHandler(service *messagesService.MessageService) *MessageHandler {
	return &MessageHandler{
		Service: service,
	}
}

func (h *MessageHandler) GetMessages(c echo.Context) error {
	allMessages, err := h.Service.GetAllMessages()
	if err != nil {
		return err
	}

	var response []messages.Message
	for _, msg := range allMessages {
		message := messages.Message{
			Id:      &msg.ID,
			Message: &msg.Text,
		}
		response = append(response, message)
	}

	return c.JSON(http.StatusOK, response)
}

func (h *MessageHandler) PostMessages(c echo.Context) error {
	var request messages.PostMessagesRequestObject
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	if request.Body.Message == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Message field is required")
	}

	messageToCreate := messagesService.Message{Text: *request.Body.Message}
	createdMessage, err := h.Service.CreateMessage(messageToCreate)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create message")
	}

	response := messages.PostMessages201JSONResponse{
		Id:      &createdMessage.ID,
		Message: &createdMessage.Text,
	}
	return c.JSON(http.StatusCreated, response)
}

func (h *MessageHandler) PatchMessages(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	var request messages.PatchMessagesRequestObject
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	if request.Body.Message == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Message field is required")
	}

	messageUpdate := messagesService.Message{Text: *request.Body.Message}
	updatedMessage, err := h.Service.UpdateMessageByID(id, messageUpdate)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update message")
	}

	response := messages.PatchMessages200JSONResponse{
		Id:      &updatedMessage.ID,
		Message: &updatedMessage.Text,
	}
	return c.JSON(http.StatusOK, response)
}

func (h *MessageHandler) DeleteMessages(c echo.Context) error {
	idStr := c.Param("id")         // Get ID as string
	id, err := strconv.Atoi(idStr) // Convert string to int
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	err = h.Service.DeleteMessageByID(id)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
