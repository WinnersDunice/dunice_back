package login

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	s "github.com/WinnersDunice/dunice_back/sso/internal/structs"
	ts "github.com/WinnersDunice/dunice_back/sso/internal/tokenset"
	"github.com/gorilla/sessions"
)

func LoginHandler(store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user s.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error decoding JSON: %s", err), http.StatusBadRequest)
			return
		}

		z := s.HashPassword(user.Password)
		user.Password = z

		log.Print(user)

		jsonData, err := json.Marshal(user)
		if err != nil {
			http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
			return
		}

		resp, err := http.Post(s.IPDB+"/users/auth", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			http.Error(w, fmt.Sprintf("Error sending request: %s", err), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			http.Error(w, "Invalid credentials "+fmt.Sprintf("%d", resp.StatusCode), http.StatusUnauthorized)
			return
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Error reading response body", http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(body, &user)
		if err != nil {
			http.Error(w, "Error unmarshalling JSON", http.StatusInternalServerError)
			return
		}

		session, err := store.Get(r, "auth")
		if err != nil {
			http.Error(w, "Error getting session: "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "userID",
			Value:    fmt.Sprintf("%d", user.UserID),
			Secure:   true,                  // Убедитесь, что это true, если используете HTTPS
			HttpOnly: false,                 // Защита от XSS
			SameSite: http.SameSiteNoneMode, // Для работы в сторонних контекстах
			MaxAge:   86400,                 // Установите MaxAge, если хотите, чтобы куки сохранялись
		})
		http.SetCookie(w, &http.Cookie{
			Name:     "role",
			HttpOnly: false,
			Value:    fmt.Sprintf("%d", user.Role),
			SameSite: http.SameSiteNoneMode,
			Secure:   true,
			MaxAge:   0,
		})
		// Сохраняем значения в сессии
		session.Values["authenticated"] = true
		session.Values["userID"] = user.UserID
		session.Values["role"] = user.Role

		// Логируем значения перед сохранением
		log.Printf("Saving session values: authenticated=%v, userID=%d, role=%d", session.Values["authenticated"], session.Values["userID"], session.Values["role"])

		session.Options.MaxAge = 86400 // 24 hours
		session.Options.SameSite = http.SameSiteNoneMode
		session.Options.Secure = true

		err = session.Save(r, w)
		if err != nil {
			http.Error(w, "Error saving session: "+err.Error(), http.StatusInternalServerError)
			return
		}

		ts.Add(session)
	}
}
