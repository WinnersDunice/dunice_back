package handler

import (
	"encoding/json"
	"net/http"
	"sync"
	"log"
	"github.com/go-chi/chi"
)

// UserInfo представляет информацию о пользователе
type UserInfo struct {
	Apps []string `json:"apps"`
}

// UserStatsHandler обрабатывает запросы к информации о пользователях
type UserStatsHandler struct {
	mu    sync.Mutex
	users map[string]UserInfo
}

// NewUserStatsHandler создает новый обработчик для статистики пользователей
func NewUserStatsHandler() *UserStatsHandler {
	return &UserStatsHandler{
		users: make(map[string]UserInfo),
	}
}

// SetUserInfo устанавливает информацию о пользователе
func (h *UserStatsHandler) SetUserInfo(w http.ResponseWriter, r *http.Request) {
	mac := chi.URLParam(r, "mac")
	
	var userInfo UserInfo
	if err := json.NewDecoder(r.Body).Decode(&userInfo); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	log.Print(userInfo)
	h.mu.Lock()
	h.users[mac] = userInfo
	h.mu.Unlock()

	w.WriteHeader(http.StatusNoContent)
}

// GetUserInfo получает информацию о пользователе
func (h *UserStatsHandler) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	mac := chi.URLParam(r, "mac")

	h.mu.Lock()
	userInfo, exists := h.users[mac]
	h.mu.Unlock()

	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(userInfo); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}
