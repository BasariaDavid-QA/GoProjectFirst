package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetMessages(w http.ResponseWriter, r *http.Request) {
	var messages []Message
	result := DB.Find(&messages)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var message Message

	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := DB.Create(&message)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Message created successfully!"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	InitDB() // Инициализация базы данных

	DB.AutoMigrate(&Message{}) // Автоматическая миграция модели Message

	router := mux.NewRouter()
	router.HandleFunc("/api/messages", CreateMessage).Methods("POST")
	router.HandleFunc("/api/messages", GetMessages).Methods("GET")
	http.ListenAndServe(":8080", router)
}
