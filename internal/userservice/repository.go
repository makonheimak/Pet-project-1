package userservice

import "gorm.io/gorm"

//CRUD
type UserRepository interface {
	PostUser(req *User) error
	GetAllUsers() ([]User, error)
	GetUserByID(id int64) (User, error)
	PatchUserByID(user User) error
	DeleteUserByID(id int64) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) PostUser(req *User) error {
	return r.db.Create(req).Error
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var users = []User{}
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) GetUserByID(id int64) (User, error) {
	var user User
	err := r.db.First(&user, "id = ?", id).Error
	return user, err
}

func (r *userRepository) PatchUserByID(user User) error {
	return r.db.Model(&user).Updates(map[string]interface{}{
		"email":    user.Email,
		"password": user.Password,
	}).Error
}

func (r *userRepository) DeleteUserByID(id int64) error {
	return r.db.Delete(&User{}, "id = ?", id).Error
}
