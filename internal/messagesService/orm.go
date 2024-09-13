package messagesService

import "gorm.io/gorm"

// Определение структуры Message
type Message struct {
	gorm.Model
	Text string `json:"text"`
}
