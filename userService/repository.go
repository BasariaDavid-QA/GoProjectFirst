package userService

import (
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user *User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) GetAll() ([]User, error) {
	var users []User
	err := r.DB.Find(&users).Error
	return users, err
}
func (r *UserRepository) GetByID(id uint) (*User, error) {
	var user User
	err := r.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *UserRepository) UpdateUserFields(id uint, updates map[string]interface{}) error {
	return r.DB.Model(&User{}).Where("id = ?", id).Updates(updates).Error
}

func (r *UserRepository) Update(user *User) error {
	return r.DB.Save(user).Error
}
func (r *UserRepository) Delete(id uint) error {
	return r.DB.Delete(&User{}, id).Error
}
