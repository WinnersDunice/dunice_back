package validate

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/WinnersDunice/dunice_back/sso/internal/structs"
	"github.com/gorilla/sessions"
)

func ValidateHandler(store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Получаем сессию
		log.Print("ValidateHandler called")
		cookies := r.Cookies()
		for _, cookie := range cookies {
			log.Printf("Cookie: Name=%s, Value=%s", cookie.Name, cookie.Value)
		}
		session, err := store.Get(r, "auth")

		if err != nil {
			log.Print(err)
			http.Error(w, "Error getting session: "+err.Error(), http.StatusInternalServerError)
			return
		}
		log.Print("Session retrieved successfully")

		// Проверяем, аутентифицирован ли пользователь
		authenticated, ok := session.Values["authenticated"].(bool)
		if !ok || !authenticated {
			log.Print("User is not authenticated")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Логируем userID, officeid и isadmin из сессии
		login, ok := session.Values["login"].(string)
		if !ok {
			log.Print("login not found in session")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		officeid, ok := session.Values["officeid"].(int)
		if !ok {
			log.Print("OfficeID not found in session")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		isadmin, ok := session.Values["isadmin"].(bool)
		if !ok {
			log.Print("IsAdmin not found in session")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		log.Printf("UserID from session: %s, OfficeID: %d, IsAdmin: %t", login, officeid, isadmin)

		// Формируем ответ
		user := structs.User{
			Login:   login,
			OfficeID: officeid,
			IsAdmin: isadmin,
		}

		// Сохраняем данные в контексте
		ctx := context.WithValue(r.Context(), "userID", login)
		ctx = context.WithValue(ctx, "officeid", officeid)
		ctx = context.WithValue(ctx, "isadmin", isadmin)

		// Передаем новый контекст дальше
		r = r.WithContext(ctx)

		// Устанавливаем заголовок Content-Type
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Отправляем JSON-ответ
		if err := json.NewEncoder(w).Encode(user); err != nil {
			log.Print("Error encoding JSON response: ", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}
