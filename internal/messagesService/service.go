package messagesService

// Структура для сервиса сообщений
type MessageService struct {
	repo MessageRepository
}

// Создание нового сервиса сообщений
func NewMessageService(repo MessageRepository) *MessageService {
	return &MessageService{repo: repo}
}

// Создание нового сообщения
func (s *MessageService) CreateMessage(message Message) (Message, error) {
	return s.repo.CreateMessage(message)
}

// Получение всех сообщений
func (s *MessageService) GetAllMessages() ([]Message, error) {
	return s.repo.GetAllMessages()
}

// Обновление сообщения по ID
func (s *MessageService) UpdateMessageByID(id int, message Message) (Message, error) {
	return s.repo.UpdateMessageByID(id, message)
}

// Удаление сообщения по ID
func (s *MessageService) DeleteMessageByID(id int) error {
	return s.repo.DeleteMessageByID(id)
}
