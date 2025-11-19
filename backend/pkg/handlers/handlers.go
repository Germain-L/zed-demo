package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"zed-demo/pkg/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type handler struct {
	db *gorm.DB
}

func New(db *gorm.DB) *handler {
	return &handler{db}
}

func (h *handler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.db.Create(&todo).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

func (h *handler) GetTodos(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo
	if err := h.db.Find(&todos).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todos)
}

func (h *handler) GetTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var todo models.Todo
	if err := h.db.First(&todo, id).Error; err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

func (h *handler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var todo models.Todo
	if err := h.db.First(&todo, id).Error; err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.db.Save(&todo).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

func (h *handler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var todo models.Todo
	if err := h.db.First(&todo, id).Error; err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}
	if err := h.db.Delete(&todo).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
