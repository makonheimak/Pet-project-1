package userservice

import "myproject/internal/taskservice"

type UserService interface {
	PostUser(req User) (User, error)
	GetAllUsers() ([]User, error)
	GetUserByID(id int64) (User, error)
	GetTasksForUser(userID uint) ([]taskservice.Task, error)
	PatchUserByID(id int64, email, password string) (User, error)
	DeleteUserByID(id int64) error
}

type userService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) PostUser(req User) (User, error) {
	err := s.repo.PostUser(&req)
	if err != nil {
		return User{}, err
	}
	return req, nil
}

func (s *userService) GetAllUsers() ([]User, error) {
	users, err := s.repo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userService) GetUserByID(id int64) (User, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *userService) GetTasksForUser(userID uint) ([]taskservice.Task, error) {
	return s.repo.GetTasksForUser(userID)
}

func (s *userService) PatchUserByID(id int64, email, password string) (User, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return User{}, err
	}

	user.Email = email
	user.Password = password

	if err := s.repo.PatchUserByID(user); err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *userService) DeleteUserByID(id int64) error {
	err := s.repo.DeleteUserByID(id)
	if err != nil {
		return err
	}
	return nil
}
