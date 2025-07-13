package handlers

import (
	"context"
	"myproject/internal/userservice"
	"myproject/internal/web/users"
	"time"
)

type UserHandlers struct {
	service userservice.UserService
}

func NewUserHandlers(s userservice.UserService) *UserHandlers {
	return &UserHandlers{service: s}
}

func (h *UserHandlers) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	usersAll, err := h.service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var responseUsers []users.User
	for _, t := range usersAll {
		tCopy := t

		var deletedAt *time.Time
		if tCopy.DeletedAt.Valid {
			deletedAt = &tCopy.DeletedAt.Time
		}

		responseUsers = append(responseUsers, users.User{
			Id:        &tCopy.ID,
			Email:     tCopy.Email,
			Password:  tCopy.Password,
			CreatedAt: &tCopy.CreatedAt,
			UpdatedAt: &tCopy.UpdatedAt,
			DeletedAt: deletedAt,
		})
	}

	return users.GetUsers200JSONResponse(responseUsers), nil
}

func (h *UserHandlers) PostUser(ctx context.Context, request users.PostUserRequestObject) (users.PostUserResponseObject, error) {
	newUser := userservice.User{
		Email:    request.Body.Email,
		Password: request.Body.Password,
	}

	createdUser, err := h.service.PostUser(newUser)
	if err != nil {
		return nil, err
	}

	response := users.User{
		Id:       &createdUser.ID,
		Email:    createdUser.Email,
		Password: createdUser.Password,
	}

	return users.PostUser201JSONResponse(response), nil
}

func (h *UserHandlers) PatchUserByID(ctx context.Context, request users.PatchUserByIDRequestObject) (users.PatchUserByIDResponseObject, error) {
	updatedUser, err := h.service.PatchUserByID(request.Id, *request.Body.Email, *request.Body.Password)

	if err != nil {
		return nil, err
	}

	response := users.User{
		Id:       &updatedUser.ID,
		Email:    updatedUser.Email,
		Password: updatedUser.Password,
	}

	return users.PatchUserByID200JSONResponse(response), nil
}

func (h *UserHandlers) DeleteUserByID(ctx context.Context, request users.DeleteUserByIDRequestObject) (users.DeleteUserByIDResponseObject, error) {
	err := h.service.DeleteUserByID(request.Id)
	if err != nil {
		return nil, err
	}

	return users.DeleteUserByID204Response{}, nil
}
