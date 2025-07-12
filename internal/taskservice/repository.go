package taskservice

import "gorm.io/gorm"

//CRUD
type TaskRepository interface {
	CreateTask(req *Task) error
	GetAllTasks() ([]Task, error)
	GetTaskByID(id int64) (Task, error)
	UpdateTask(task Task) error
	DeleteTask(id int64) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(req *Task) error {
	return r.db.Create(&req).Error
}

func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks = []Task{}
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) GetTaskByID(id int64) (Task, error) {
	var task Task
	err := r.db.First(&task, "id = ?", id).Error
	return task, err
}

func (r *taskRepository) UpdateTask(task Task) error {
	return r.db.Save(&task).Error
}

func (r *taskRepository) DeleteTask(id int64) error {
	return r.db.Delete(&Task{}, "id = ?", id).Error
}
