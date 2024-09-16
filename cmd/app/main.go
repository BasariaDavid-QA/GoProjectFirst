package main

import (
	"log"
	"new-go-project/internal/database"
	"new-go-project/internal/handlers"
	"new-go-project/internal/messagesService"
	"new-go-project/userService"

	"github.com/labstack/echo/v4"
)

func main() {
	database.InitDB()
	if err := database.DB.AutoMigrate(&messagesService.Message{}, &userService.User{}); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

	messageRepo := messagesService.NewMessageRepository(database.DB)
	messageService := messagesService.NewMessageService(messageRepo)

	messageHandler := handlers.MessageHandler{Service: messageService}

	userRepo := userService.NewUserRepository(database.DB)
	userService := userService.NewUserService(userRepo)
	userHandler := handlers.UserHandler{Service: userService}

	e := echo.New()

	e.GET("/messages", messageHandler.GetMessages)
	e.POST("/messages", messageHandler.PostMessages)
	e.PATCH("/messages/:id", messageHandler.PatchMessages)
	e.DELETE("/messages/:id", messageHandler.DeleteMessages)

	e.GET("/users", userHandler.GetUsers)
	e.POST("/users", userHandler.PostUsers)
	e.PATCH("/users/:id", userHandler.PatchUserByID)
	e.DELETE("/users/:id", userHandler.DeleteUserByID)

	e.Logger.Fatal(e.Start(":8080"))
}
