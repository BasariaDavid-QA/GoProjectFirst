package messagesService

import "gorm.io/gorm"

// Интерфейс для работы с сообщениями
type MessageRepository interface {
	CreateMessage(message Message) (Message, error)
	GetAllMessages() ([]Message, error)
	UpdateMessageByID(id int, message Message) (Message, error)
	DeleteMessageByID(id int) error
}

type messageRepository struct {
	db *gorm.DB
}

// Создание нового репозитория сообщений
func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &messageRepository{db: db}
}

// Создание нового сообщения
func (r *messageRepository) CreateMessage(message Message) (Message, error) {
	result := r.db.Create(&message)
	if result.Error != nil {
		return Message{}, result.Error
	}
	return message, nil
}

// Получение всех сообщений
func (r *messageRepository) GetAllMessages() ([]Message, error) {
	var messages []Message
	err := r.db.Find(&messages).Error
	return messages, err
}

// Обновление сообщения по ID
func (r *messageRepository) UpdateMessageByID(id int, message Message) (Message, error) {
	var existingMessage Message
	if err := r.db.First(&existingMessage, id).Error; err != nil {
		return Message{}, err
	}
	existingMessage.Text = message.Text
	if err := r.db.Save(&existingMessage).Error; err != nil {
		return Message{}, err
	}
	return existingMessage, nil
}

// Удаление сообщения по ID
func (r *messageRepository) DeleteMessageByID(id int) error {
	if err := r.db.Delete(&Message{}, id).Error; err != nil {
		return err
	}
	return nil
}
