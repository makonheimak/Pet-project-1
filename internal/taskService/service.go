package taskservice

type TaskService interface {
	CreateTask(req Task) (Task, error)
	GetAllTasks() ([]Task, error)
	GetTaskByID(id string) (Task, error)
	UpdateTask(id, task string) (Task, error)
	DeleteTask(id string) error
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

func (s *taskService) GetTaskByID(id string) (Task, error) {
	tasks, err := s.repo.GetTaskByID(id)
	if err != nil {
		return Task{}, err
	}
	return tasks, nil
}

func (s *taskService) UpdateTask(id string, newText string) (Task, error) {
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

func (s *taskService) DeleteTask(id string) error {
	err := s.repo.DeleteTask(id)
	if err != nil {
		return err
	}
	return nil
}
