package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetMessages(w http.ResponseWriter, r *http.Request) {
	var messages []Message
	DB.Find(&messages)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var message Message
	json.NewDecoder(r.Body).Decode(&message)
	DB.Create(&message)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Message created successfully!"}
	json.NewEncoder(w).Encode(response)
}

func UpdateMessage(w http.ResponseWriter, r *http.Request) {
	var message Message
	vars := mux.Vars(r)
	id := vars["id"]

	if err := DB.First(&message, id).Error; err != nil {
		http.Error(w, "Message not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := DB.Save(&message).Error; err != nil {
		http.Error(w, "Failed to update message", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Message updated successfully!"}
	json.NewEncoder(w).Encode(response)
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	var message Message
	vars := mux.Vars(r)
	id := vars["id"]

	if err := DB.First(&message, id).Error; err != nil {
		http.Error(w, "Message not found", http.StatusNotFound)
		return
	}

	if err := DB.Delete(&message).Error; err != nil {
		http.Error(w, "Failed to delete message", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Message deleted successfully!"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	InitDB()
	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/messages", CreateMessage).Methods("POST")
	router.HandleFunc("/api/messages", GetMessages).Methods("GET")
	router.HandleFunc("/api/messages/{id:[0-9]+}", UpdateMessage).Methods("PATCH")
	router.HandleFunc("/api/messages/{id:[0-9]+}", DeleteMessage).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
