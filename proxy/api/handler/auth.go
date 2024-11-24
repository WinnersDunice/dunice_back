package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"

	"net/http"

	"github.com/WinnersDunice/dunice_back/proxy/entities"
)



func ValidateSessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://%s:%s/api_auth_dunice_server_sso/validate")
		if err != nil {
			http.Error(w, "Error sending validation request", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Check the response status code
		if resp.StatusCode != http.StatusOK {
			http.Error(w, "Validation failed", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

	url := fmt.Sprintf("http://%s:%s/api_auth_dunice_server_sso/login", IP, SSOPort)
	log.Print(url)

	resp, err := http.Post(url, "application/json", r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	for _, cookie := range resp.Cookies() {
		cookie.SameSite = http.SameSiteNoneMode
		cookie.Domain = ".winnersdunice.ru" // Устанавливаем домен для куки
		http.SetCookie(w, cookie)
	}
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var user entities.SmallUser
	err := json.NewDecoder(r.Body).Decode(&user)
	log.Print(user)
	log.Print(err)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusBadRequest)
		return
	}

	url_reg := fmt.Sprintf("http://%s:%s/database/users", IP, UsersPort)
	url_login := fmt.Sprintf("http://%s:%s/api_auth_dunice_server_sso/login", IP, SSOPort)
	log.Print(url_login)
	log.Print(user)
	userJSON, err := json.Marshal(user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusBadRequest)
		return
	}

	resp, err := http.Post(url_reg, "application/json", bytes.NewBuffer(userJSON))
	if err != nil {
		log.Print(1)
		log.Print(err)
		http.Error(w, fmt.Sprintf("Error 1: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close() // Закрываем тело ответа

	log.Print(resp.Body)
	log.Println("Sent to ", url_reg)

	resp2, err := http.Post(url_login, "application/json", bytes.NewBuffer(userJSON))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error 2: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp2.Body.Close() // Закрываем тело ответа

	body2, err := io.ReadAll(resp2.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error 3: %v", err), http.StatusInternalServerError)
		return
	}

	// Устанавливаем куки из ответа на вход
	for _, cookie := range resp2.Cookies() {
		cookie.SameSite = http.SameSiteNoneMode
		cookie.Domain = ".dunicewinners.ru" // Устанавливаем домен для куки
		http.SetCookie(w, cookie)
	}

	// Устанавливаем заголовки ответа
	w.Header().Set("Content-Type", resp2.Header.Get("Content-Type"))
	w.WriteHeader(http.StatusCreated)
	w.Write(body2)
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	client := &http.Client{}
	url := fmt.Sprintf("http://%s:%s/api_auth_dunice_server_sso/logout", IP, SSOPort)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error 1: %v", err), http.StatusInternalServerError)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error 2: %v", err), http.StatusInternalServerError)
		return
	}
	for _, cookie := range resp.Cookies() {
		http.SetCookie(w, cookie)
	}

	resp.Body.Close()
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	w.WriteHeader(http.StatusNoContent)


}
