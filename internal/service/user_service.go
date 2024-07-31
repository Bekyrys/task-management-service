package service

import (
    "github.com/Bekyrys/task-manager/internal/models"
    "github.com/Bekyrys/task-manager/internal/repository"
)

type UserService struct {
    Repo *repository.UserRepository
}

func (s *UserService) CreateUser(user *models.User) error {
    return s.Repo.Create(user)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
    return s.Repo.GetAll()
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
    return s.Repo.GetByID(id)
}

func (s *UserService) UpdateUser(user *models.User) error {
    return s.Repo.Update(user)
}

func (s *UserService) DeleteUser(id int) error {
    return s.Repo.Delete(id)
}

func (s *UserService) FindUsersByName(name string) ([]models.User, error) {
    return s.Repo.FindByName(name)
}

func (s *UserService) FindUsersByEmail(email string) ([]models.User, error) {
    return s.Repo.FindByEmail(email)
}
