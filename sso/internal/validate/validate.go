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
		a, ok := session.Values["authenticated"].(bool)
		log.Print("Session authenticated:", a, ok)

		// Логируем значения сессии
		log.Printf("Session values: %+v", session.Values)

		// Проверяем, аутентифицирован ли пользователь
		authenticated, ok := session.Values["authenticated"].(bool)
		if !ok || !authenticated {
			log.Print("User is not authenticated")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Логируем userID и уровень привилегий из сессии
		userID, ok := session.Values["userID"].(int) // Предполагаем, что userID хранится как int
		if !ok {
			log.Print("UserID not found in session")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		role, ok := session.Values["role"].(int) // Предполагаем, что роль хранится как int
		if !ok {
			log.Print("Role not found in session")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		log.Printf("UserID from session: %d, Role: %d", userID, role)

		// Формируем ответ
		user := structs.User{
			UserID: userID,
			Role:   role,
		}

		// Сохраняем данные в контексте
		ctx := context.WithValue(r.Context(), "userID", userID)
		ctx = context.WithValue(ctx, "role", role)

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
