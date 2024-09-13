package main

import (
	"net/http"
	"new-go-project/internal/database"
	"new-go-project/internal/handlers"
	"new-go-project/internal/messagesService"

	"github.com/gorilla/mux"
)

func main() {
	// Инициализация базы данных
	database.InitDB()
	database.DB.AutoMigrate(&messagesService.Message{})

	// Создание необходимых объектов репозитория, сервиса и хендлеров
	repo := messagesService.NewMessageRepository(database.DB)
	service := messagesService.NewService(repo)
	handler := handlers.NewHandler(service)

	// Создание роутера и маршрутов
	router := mux.NewRouter()
	router.HandleFunc("/api/post", handler.PostMessageHandler).Methods("POST")
	router.HandleFunc("/api/get", handler.GetMessagesHandler).Methods("GET")
	router.HandleFunc("/api/update/{id:[0-9]+}", handler.PatchMessageHandler).Methods("PATCH")
	router.HandleFunc("/api/delete/{id:[0-9]+}", handler.DeleteMessageHandler).Methods("DELETE")

	// Запуск сервера на порту 8080
	http.ListenAndServe(":8080", router)
}
