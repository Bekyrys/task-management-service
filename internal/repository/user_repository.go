package repository

import (
    "database/sql"
    "github.com/Bekyrys/task-manager/internal/models"
)

type UserRepository struct {
    DB *sql.DB
}

func (r *UserRepository) Create(user *models.User) error {
    query := "INSERT INTO users (name, email, role, created_at) VALUES ($1, $2, $3, $4) RETURNING id"
    return r.DB.QueryRow(query, user.Name, user.Email, user.Role, user.CreatedAt).Scan(&user.ID)
}

func (r *UserRepository) GetAll() ([]models.User, error) {
    query := "SELECT id, name, email, role, created_at FROM users"
    rows, err := r.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []models.User
    for rows.Next() {
        var user models.User
        if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt); err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    return users, nil
}

func (r *UserRepository) GetByID(id int) (*models.User, error) {
    query := "SELECT id, name, email, role, created_at FROM users WHERE id=$1"
    var user models.User
    err := r.DB.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserRepository) Update(user *models.User) error {
    query := "UPDATE users SET name=$1, email=$2, role=$3 WHERE id=$4"
    _, err := r.DB.Exec(query, user.Name, user.Email, user.Role, user.ID)
    return err
}

func (r *UserRepository) Delete(id int) error {
    query := "DELETE FROM users WHERE id=$1"
    _, err := r.DB.Exec(query, id)
    return err
}

func (r *UserRepository) FindByName(name string) ([]models.User, error) {
    query := "SELECT id, name, email, role, created_at FROM users WHERE name ILIKE '%' || $1 || '%'"
    rows, err := r.DB.Query(query, name)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []models.User
    for rows.Next() {
        var user models.User
        if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt); err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    return users, nil
}

func (r *UserRepository) FindByEmail(email string) ([]models.User, error) {
    query := "SELECT id, name, email, role, created_at FROM users WHERE email ILIKE '%' || $1 || '%'"
    rows, err := r.DB.Query(query, email)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []models.User
    for rows.Next() {
        var user models.User
        if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt); err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    return users, nil
}