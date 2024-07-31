package service

import (
    "github.com/Bekyrys/task-manager/internal/models"
    "github.com/Bekyrys/task-manager/internal/repository"
)

type TaskService struct {
    Repo *repository.TaskRepository
}

func (s *TaskService) CreateTask(task *models.Task) error {
    return s.Repo.Create(task)
}

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
    return s.Repo.GetAll()
}

func (s *TaskService) GetTaskByID(id int) (*models.Task, error) {
    return s.Repo.GetByID(id)
}

func (s *TaskService) UpdateTask(task *models.Task) error {
    return s.Repo.Update(task)
}

func (s *TaskService) DeleteTask(id int) error {
    return s.Repo.Delete(id)
}

func (s *TaskService) FindTasksByTitle(title string) ([]models.Task, error) {
    return s.Repo.FindByTitle(title)
}

func (s *TaskService) FindTasksByStatus(status string) ([]models.Task, error) {
    return s.Repo.FindByStatus(status)
}

func (s *TaskService) FindTasksByPriority(priority string) ([]models.Task, error) {
    return s.Repo.FindByPriority(priority)
}

func (s *TaskService) FindTasksByAssigneeID(assigneeID int) ([]models.Task, error) {
    return s.Repo.FindByAssigneeID(assigneeID)
}

func (s *TaskService) FindTasksByProjectID(projectID int) ([]models.Task, error) {
    return s.Repo.FindByProjectID(projectID)
}
