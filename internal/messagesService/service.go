package messagesService

type MessageService struct {
	repo MessageRepository
}

func NewService(repo MessageRepository) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) GetAllMessages() ([]Message, error) {
	return s.repo.GetAllMessages()
}

func (s *MessageService) CreateMessage(message Message) (Message, error) {
	return s.repo.CreateMessage(message)
}

func (s *MessageService) UpdateMessageByID(ID uint, message Message) (Message, error) {
	return s.repo.UpdateMessageByID(ID, message)
}

func (s *MessageService) DeleteMessageByID(ID uint) (Message, error) {
	return s.repo.DeleteMessageByID(ID)
}
