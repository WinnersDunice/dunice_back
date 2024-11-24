package handler

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

const ssoEndpoint = "http://195.80.238.9:8530/api_auth_dunice_server_sso/validate"

type UserPrivileges struct {
	Email      string `json:"email"`
	UserID     int    `json:"id"`
	Password   string `json:"password"`
	Privileges int    `json:"privilege_level"`
}

func NotStrictFetchUserMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		//log.Print("Unstrict")

		cookies := r.Cookies()
		var userID int
		var watched int
		for _, cookie := range cookies {
			if cookie.Name == "userID" {
				if id, err := strconv.Atoi(cookie.Value); err == nil {
					userID = id
				}

			}
			if cookie.Name == "watched" {
				if id, err := strconv.Atoi(cookie.Value); err == nil {
					log.Print("Watched: ", id)
					watched = id

				}

			}
			//log.Printf("Cookie: Name=%s, Value=%s", cookie.Name, cookie.Value)

		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "userID", userID)
		ctx = context.WithValue(ctx, "watched", watched)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func FetchUserMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		log.Print("ZZZZ")
		// Логируем куки
		var watched int
		cookies := r.Cookies()
		for _, cookie := range cookies {
			if cookie.Name == "watched" {
				if id, err := strconv.Atoi(cookie.Value); err == nil {
					watched = id
				}
				break
			}
			//log.Printf("Cookie: Name=%s, Value=%s", cookie.Name, cookie.Value)
		}

		// Создаем новый запрос к ssoEndpoint с текущими куками
		req, err := http.NewRequest("GET", ssoEndpoint, nil)
		if err != nil {
			http.Error(w, "Error creating request", http.StatusInternalServerError)
			return
		}

		// Копируем куки из текущего запроса
		for _, cookie := range cookies {
			req.AddCookie(cookie)
		}

		// Отправляем запрос
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, "Error sending request", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Error reading response body", http.StatusInternalServerError)
			return
		}

		var userPrivileges UserPrivileges
		if err := json.Unmarshal(body, &userPrivileges); err != nil {
			http.Error(w, "Error unmarshalling response", http.StatusInternalServerError)
			return
		}

		// Добавляем данные в контекст
		ctx := r.Context()
		ctx = context.WithValue(ctx, "watched", watched)
		ctx = context.WithValue(ctx, "userID", userPrivileges.UserID)
		ctx = context.WithValue(ctx, "privileges", userPrivileges.Privileges)
		log.Print("userID: ", userPrivileges.UserID)
		log.Print("privileges: ", userPrivileges.Privileges)

		// Передаем управление следующему обработчику
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
