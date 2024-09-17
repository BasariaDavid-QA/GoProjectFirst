package userService

import (
	"gorm.io/gorm"
)

type UserRepository interface {
	// CreateUser - возвращаем созданно пользователя и ошибку
	CreateUser(user User) (User, error)

	// GetAllUsers - Возвращаем массив из всех пользователей в БД и ошибку
	GetAllUsers() ([]User, error)

	// UpdateUserByID - Передаем id и User, возвращаем обновленный User и ошибку
	UpdateUserByID(id uint, user User) (User, error)

	// DeleteUserByID - Передаем id для удаления, возвращаем удаленного пользователя и ошибку
	DeleteUserByID(id uint) (User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

// (r *messageRepository) привязывает данную функцию к нашему репозиторию

func (r *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) CreateUser(user User) (User, error) {
	createUser := r.db.Create(&user)
	if createUser.Error != nil {
		return User{}, createUser.Error
	}
	return user, nil
}

func (r *userRepository) UpdateUserByID(id uint, user User) (User, error) {
	var existingUser User // В этой переменной храним существующее сообщение из БД

	// Ищем существующее сообщение в БД по заданному id
	result := r.db.First(&existingUser, id)
	if result.Error != nil {
		return User{}, result.Error
	}
	// Обозначаем поля, которые будем обновлять
	updates := map[string]interface{}{
		"Email":    user.Email,
		"Password": user.Password,
	}
	// Обновляем запись в БД, соответствующую заданному id
	result = r.db.Model(&User{}).Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		return User{}, result.Error
	}

	// Обновляем existingMessage, чтобы вернуть актуальные данные
	result = r.db.First(&existingUser, id)
	if result.Error != nil {
		return User{}, result.Error
	}

	return existingUser, nil
}

func (r *userRepository) DeleteUserByID(id uint) (User, error) {
	var user User
	deletedUser := r.db.First(&user, id)
	if deletedUser.Error != nil {
		return User{}, deletedUser.Error
	}

	deletedUser = r.db.Delete(&user, id)
	if deletedUser.Error != nil {
		return User{}, deletedUser.Error
	}
	return user, nil
}
