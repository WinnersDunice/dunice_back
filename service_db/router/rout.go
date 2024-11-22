package router

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	us "github.com/WinnersDunice/dunice_back/service_db/internal/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/crypto/bcrypt"
)

func Rout(db *sql.DB) error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("never should have pwn here, FSB went 4 u <3"))
	})

	api := chi.NewRouter()
	r.Mount("/database_zov_russ_cbo", api)

	v1 := chi.NewRouter()
	api.Mount("/users", v1)

	// Create a new user
	v1.Post("/", func(w http.ResponseWriter, r *http.Request) {
		user := new(us.User)
		if err := json.NewDecoder(r.Body).Decode(user); err != nil {
			http.Error(w, "cannot parse JSON", http.StatusBadRequest)
			return
		}
		if err := us.CreateUser(db, user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(user)
	})

	// Get a user by ID
	v1.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Print(err)
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		user, err := us.GetUserByID(db, id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				http.Error(w, "user not found", http.StatusNotFound)
				return
			}
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(user)
	})

	// Update a user by ID
	v1.Put("/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Print(err)
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		updatedUser := new(us.User)
		if err := json.NewDecoder(r.Body).Decode(updatedUser); err != nil {
			http.Error(w, "cannot parse JSON", http.StatusBadRequest)
			return
		}
		if err := us.UpdateUser(db, id, updatedUser); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(updatedUser)
	})

	// Update user login
	v1.Put("/login/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Print(err)
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		type LoginRequest struct {
			Login string `json:"login"`
		}
		var req LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "cannot parse JSON", http.StatusBadRequest)
			return
		}
		if err := us.UpdateUserLogin(db, id, req.Login); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"message": "login updated successfully"})
	})

	// Update user password with old password check
	v1.Put("/password/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Print(err)
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		type PasswordRequest struct {
			OldPassword string `json:"old_password"`
			NewPassword string `json:"new_password"`
		}
		var req PasswordRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "cannot parse JSON", http.StatusBadRequest)
			return
		}
		// Get the current hashed password
		currentHashedPassword, err := us.GetUserPassword(db, id)
		if err != nil {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}
		// Check if the old password matches the current hashed password
		if err := bcrypt.CompareHashAndPassword([]byte(currentHashedPassword), []byte(req.OldPassword)); err != nil {
			http.Error(w, "old password is incorrect", http.StatusUnauthorized)
			return
		}
		if err := us.UpdateUserPassword(db, id, req.NewPassword); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"message": "password updated successfully"})
	})

	// Update user MAC address
	v1.Put("/macaddress/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Print(err)
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		type MacAddressRequest struct {
			MacAddress string `json:"macaddress"`
		}
		var req MacAddressRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "cannot parse JSON", http.StatusBadRequest)
			return
		}
		if err := us.UpdateUserMacAddress(db, id, req.MacAddress); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"message": "MAC address updated successfully"})
	})

	// Get user login
	v1.Get("/login/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Print(err)
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		login, err := us.GetUserLogin(db, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"login": login})
	})

	// Get user password
	v1.Get("/password/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Print(err)
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		password, err := us.GetUserPassword(db, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"password": password})
	})

	// Delete a user by ID
	v1.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Print(err)
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		if err := us.DeleteUser(db, id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"status": "user deleted"})
	})

	return http.ListenAndServe(":8003", r)
}
