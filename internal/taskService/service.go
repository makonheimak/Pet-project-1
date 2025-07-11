package taskservice

type TaskService interface {
	CreateTask(req Task) (Task, error)
	GetAllTasks() ([]Task, error)
	GetTaskByID(id int64) (Task, error)
	UpdateTask(id int64, task string) (Task, error)
	DeleteTask(id int64) error
}

type taskService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) TaskService {
	return &taskService{repo: r}
}

func (s *taskService) CreateTask(req Task) (Task, error) {
	err := s.repo.CreateTask(&req)
	if err != nil {
		return Task{}, err
	}
	return req, nil
}

func (s *taskService) GetAllTasks() ([]Task, error) {
	tasks, err := s.repo.GetAllTasks()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *taskService) GetTaskByID(id int64) (Task, error) {
	task, err := s.repo.GetTaskByID(id)
	if err != nil {
		return Task{}, err
	}
	return task, nil
}

func (s *taskService) UpdateTask(id int64, newText string) (Task, error) {
	task, err := s.repo.GetTaskByID(id)
	if err != nil {
		return Task{}, err
	}

	task.Task = newText

	if err := s.repo.UpdateTask(task); err != nil {
		return Task{}, err
	}

	return task, nil
}

func (s *taskService) DeleteTask(id int64) error {
	err := s.repo.DeleteTask(id)
	if err != nil {
		return err
	}
	return nil
}
