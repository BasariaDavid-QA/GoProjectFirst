package main

import "gorm.io/gorm"

// Определяем структуру Message
type Message struct {
	gorm.Model
	Text string `json:"text"`
}
