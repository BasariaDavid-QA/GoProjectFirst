package userService

type UserService struct {
	Repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(user *User) error {
	return s.Repo.Create(user)
}

func (s *UserService) GetUsers() ([]User, error) {
	return s.Repo.GetAll()
}

func (s *UserService) GetUserByID(id uint) (*User, error) {
	return s.Repo.GetByID(id)
}

func (s *UserService) UpdateUserFields(id uint, updates map[string]interface{}) error {
	return s.Repo.UpdateUserFields(id, updates)
}

func (s *UserService) UpdateUser(user *User) error {
	return s.Repo.Update(user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.Repo.Delete(id)
}
