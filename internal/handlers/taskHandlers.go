package handlers

import (
	"context"
	"myproject/internal/taskservice"
	"myproject/internal/web/tasks"
)

type TaskHandlers struct {
	service taskservice.TaskService
}

func NewTaskHandlers(s taskservice.TaskService) *TaskHandlers {
	return &TaskHandlers{service: s}
}

func (h *TaskHandlers) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	tasksAll, err := h.service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	var responseTasks []tasks.Task
	for _, t := range tasksAll {
		tCopy := t
		responseTasks = append(responseTasks, tasks.Task{
			Id:   tCopy.ID,
			Task: tCopy.Task,
		})
	}

	return tasks.GetTasks200JSONResponse(responseTasks), nil
}

func (h *TaskHandlers) GetTasksByUserID(ctx context.Context, request tasks.GetTasksByUserIDRequestObject) (tasks.GetTasksByUserIDResponseObject, error) {
	tasksByUser, err := h.service.GetTasksByUserID(uint(request.UserId))
	if err != nil {
		return nil, err
	}

	var responseTasks []tasks.Task
	for _, t := range tasksByUser {
		tCopy := t
		responseTasks = append(responseTasks, tasks.Task{
			Id:     tCopy.ID,
			Task:   tCopy.Task,
			UserId: int64(tCopy.UserID),
		})
	}

	return tasks.GetTasksByUserID200JSONResponse(responseTasks), nil
}

func (h *TaskHandlers) PostTask(ctx context.Context, request tasks.PostTaskRequestObject) (tasks.PostTaskResponseObject, error) {
	newTask := taskservice.Task{
		Task:   request.Body.Task,
		UserID: uint(request.Body.UserId),
	}

	createdTask, err := h.service.CreateTask(newTask)
	if err != nil {
		return nil, err
	}

	response := tasks.Task{
		Id:     createdTask.ID,
		Task:   createdTask.Task,
		UserId: int64(createdTask.UserID),
	}

	return tasks.PostTask201JSONResponse(response), nil
}

func (h *TaskHandlers) UpdateTask(ctx context.Context, request tasks.UpdateTaskRequestObject) (tasks.UpdateTaskResponseObject, error) {
	updatedTask, err := h.service.UpdateTask(request.Id, request.Body.Task)
	if err != nil {
		return nil, err
	}

	response := tasks.Task{
		Id:     updatedTask.ID,
		Task:   updatedTask.Task,
		UserId: int64(updatedTask.UserID),
	}

	return tasks.UpdateTask200JSONResponse(response), nil
}

func (h *TaskHandlers) DeleteTask(ctx context.Context, request tasks.DeleteTaskRequestObject) (tasks.DeleteTaskResponseObject, error) {
	err := h.service.DeleteTask(request.Id)
	if err != nil {
		return nil, err
	}

	return tasks.DeleteTask204Response{}, nil
}
