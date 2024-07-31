package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "github.com/Bekyrys/task-manager/internal/models"
    "github.com/Bekyrys/task-manager/internal/service"
)

type ProjectHandler struct {
    Service *service.ProjectService
}

func (h *ProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
    var project models.Project
    if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if err := h.Service.CreateProject(&project); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(project)
}

func (h *ProjectHandler) GetAllProjects(w http.ResponseWriter, r *http.Request) {
    projects, err := h.Service.GetAllProjects()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(projects)
}

func (h *ProjectHandler) GetProjectByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid project ID", http.StatusBadRequest)
        return
    }
    project, err := h.Service.GetProjectByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(project)
}

func (h *ProjectHandler) UpdateProject(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid project ID", http.StatusBadRequest)
        return
    }
    var project models.Project
    if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    project.ID = id
    if err := h.Service.UpdateProject(&project); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(project)
}

func (h *ProjectHandler) DeleteProject(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid project ID", http.StatusBadRequest)
        return
    }
    if err := h.Service.DeleteProject(id); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}
