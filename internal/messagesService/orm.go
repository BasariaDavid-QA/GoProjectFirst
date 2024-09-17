package messagesService

import "gorm.io/gorm"

// Определение структуры Message
type Message struct {
	gorm.Model
	Message string `json:"Message"`
}
