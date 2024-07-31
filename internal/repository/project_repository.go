package repository

import (
    "database/sql"
    "github.com/Bekyrys/task-manager/internal/models"
)

type ProjectRepository struct {
    DB *sql.DB
}

func (r *ProjectRepository) Create(project *models.Project) error {
    query := "INSERT INTO projects (title, description, start_date, end_date, manager_id) VALUES ($1, $2, $3, $4, $5) RETURNING id"
    return r.DB.QueryRow(query, project.Title, project.Description, project.StartDate, project.EndDate, project.ManagerID).Scan(&project.ID)
}

func (r *ProjectRepository) GetAll() ([]models.Project, error) {
    query := "SELECT id, title, description, start_date, end_date, manager_id FROM projects"
    rows, err := r.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var projects []models.Project
    for rows.Next() {
        var project models.Project
        if err := rows.Scan(&project.ID, &project.Title, &project.Description, &project.StartDate, &project.EndDate, &project.ManagerID); err != nil {
            return nil, err
        }
        projects = append(projects, project)
    }
    return projects, nil
}

func (r *ProjectRepository) GetByID(id int) (*models.Project, error) {
    query := "SELECT id, title, description, start_date, end_date, manager_id FROM projects WHERE id=$1"
    var project models.Project
    err := r.DB.QueryRow(query, id).Scan(&project.ID, &project.Title, &project.Description, &project.StartDate, &project.EndDate, &project.ManagerID)
    if err != nil {
        return nil, err
    }
    return &project, nil
}

func (r *ProjectRepository) Update(project *models.Project) error {
    query := "UPDATE projects SET title=$1, description=$2, start_date=$3, end_date=$4, manager_id=$5 WHERE id=$6"
    _, err := r.DB.Exec(query, project.Title, project.Description, project.StartDate, project.EndDate, project.ManagerID, project.ID)
    return err
}

func (r *ProjectRepository) Delete(id int) error {
    query := "DELETE FROM projects WHERE id=$1"
    _, err := r.DB.Exec(query, id)
    return err
}

func (r *ProjectRepository) FindByTitle(title string) ([]models.Project, error) {
    query := "SELECT id, title, description, start_date, end_date, manager_id FROM projects WHERE title ILIKE '%' || $1 || '%'"
    rows, err := r.DB.Query(query, title)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var projects []models.Project
    for rows.Next() {
        var project models.Project
        if err := rows.Scan(&project.ID, &project.Title, &project.Description, &project.StartDate, &project.EndDate, &project.ManagerID); err != nil {
            return nil, err
        }
        projects = append(projects, project)
    }
    return projects, nil
}

func (r *ProjectRepository) FindByManagerID(managerID int) ([]models.Project, error) {
    query := "SELECT id, title, description, start_date, end_date, manager_id FROM projects WHERE manager_id=$1"
    rows, err := r.DB.Query(query, managerID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var projects []models.Project
    for rows.Next() {
        var project models.Project
        if err := rows.Scan(&project.ID, &project.Title, &project.Description, &project.StartDate, &project.EndDate, &project.ManagerID); err != nil {
            return nil, err
        }
        projects = append(projects, project)
    }
    return projects, nil
}