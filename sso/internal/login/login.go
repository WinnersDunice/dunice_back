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
		log.Print(resp)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error sending request: %s", err), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Print(resp.StatusCode)
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

		// Set cookies for userID, officeid, and isadmin
		http.SetCookie(w, &http.Cookie{
			Name:     "userID",
			Value:    user.Login,
			Secure:   true,                  // Убедитесь, что это true, если используете HTTPS
			HttpOnly: false,                 // Защита от XSS
			SameSite: http.SameSiteNoneMode, // Для работы в сторонних контекстах
			MaxAge:   86400,                 // Установите MaxAge, если хотите, чтобы куки сохранялись
		})
		http.SetCookie(w, &http.Cookie{
			Name:     "officeid",
			Value:    fmt.Sprintf("%d", user.OfficeID),
			HttpOnly: false,
			SameSite: http.SameSiteNoneMode,
			Secure:   true,
			MaxAge:   86400,
		})
		http.SetCookie(w, &http.Cookie{
			Name:     "isadmin",
			Value:    fmt.Sprintf("%t", user.IsAdmin),
			HttpOnly: false,
			SameSite: http.SameSiteNoneMode,
			Secure:   true,
			MaxAge:   86400,
		})

		// Сохраняем значения в сессии
		session.Values["authenticated"] = true
		session.Values["userID"] = user.Login
		session.Values["officeid"] = user.OfficeID
		session.Values["isadmin"] = user.IsAdmin

		// Логируем значения перед сохранением
		log.Printf("Saving session values: authenticated=%v, userID=%s, officeid=%d, isadmin=%t", session.Values["authenticated"], session.Values["userID"], session.Values["officeid"], session.Values["isadmin"])

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
