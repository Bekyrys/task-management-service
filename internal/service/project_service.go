package service

import (
    "github.com/Bekyrys/task-manager/internal/models"
    "github.com/Bekyrys/task-manager/internal/repository"
)

type ProjectService struct {
    Repo *repository.ProjectRepository
}

func (s *ProjectService) CreateProject(project *models.Project) error {
    return s.Repo.Create(project)
}

func (s *ProjectService) GetAllProjects() ([]models.Project, error) {
    return s.Repo.GetAll()
}

func (s *ProjectService) GetProjectByID(id int) (*models.Project, error) {
    return s.Repo.GetByID(id)
}

func (s *ProjectService) UpdateProject(project *models.Project) error {
    return s.Repo.Update(project)
}

func (s *ProjectService) DeleteProject(id int) error {
    return s.Repo.Delete(id)
}

func (s *ProjectService) FindProjectsByTitle(title string) ([]models.Project, error) {
    return s.Repo.FindByTitle(title)
}

func (s *ProjectService) FindProjectsByManagerID(managerID int) ([]models.Project, error) {
    return s.Repo.FindByManagerID(managerID)
}
