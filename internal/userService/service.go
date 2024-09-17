package userService

type UserService struct {
	repo UserRepository
}

func NewService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) CreateUser(user User) (User, error) {
	return s.repo.CreateUser(user)
}

func (s *UserService) UpdateUserByID(ID uint, user User) (User, error) {
	return s.repo.UpdateUserByID(ID, user)
}

func (s *UserService) DeleteUserByID(ID uint) (User, error) {
	return s.repo.DeleteUserByID(ID)
}
