package repository

import (
    "database/sql"
    "github.com/Bekyrys/task-manager/internal/models"
)

type TaskRepository struct {
    DB *sql.DB
}

func (r *TaskRepository) Create(task *models.Task) error {
    query := "INSERT INTO tasks (title, description, priority, status, assignee_id, project_id, created_at, completed_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
    return r.DB.QueryRow(query, task.Title, task.Description, task.Priority, task.Status, task.AssigneeID, task.ProjectID, task.CreatedAt, task.CompletedAt).Scan(&task.ID)
}

func (r *TaskRepository) GetAll() ([]models.Task, error) {
    query := "SELECT id, title, description, priority, status, assignee_id, project_id, created_at, completed_at FROM tasks"
    rows, err := r.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var tasks []models.Task
    for rows.Next() {
        var task models.Task
        if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.Status, &task.AssigneeID, &task.ProjectID, &task.CreatedAt, &task.CompletedAt); err != nil {
            return nil, err
        }
        tasks = append(tasks, task)
    }
    return tasks, nil
}

func (r *TaskRepository) GetByID(id int) (*models.Task, error) {
    query := "SELECT id, title, description, priority, status, assignee_id, project_id, created_at, completed_at FROM tasks WHERE id=$1"
    var task models.Task
    err := r.DB.QueryRow(query, id).Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.Status, &task.AssigneeID, &task.ProjectID, &task.CreatedAt, &task.CompletedAt)
    if err != nil {
        return nil, err
    }
    return &task, nil
}

func (r *TaskRepository) Update(task *models.Task) error {
    query := "UPDATE tasks SET title=$1, description=$2, priority=$3, status=$4, assignee_id=$5, project_id=$6, completed_at=$7 WHERE id=$8"
    _, err := r.DB.Exec(query, task.Title, task.Description, task.Priority, task.Status, task.AssigneeID, task.ProjectID, task.CompletedAt, task.ID)
    return err
}

func (r *TaskRepository) Delete(id int) error {
    query := "DELETE FROM tasks WHERE id=$1"
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