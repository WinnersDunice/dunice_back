package router

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	us "github.com/WinnersDunice/dunice_back/service_db/internal/user"
	of "github.com/WinnersDunice/dunice_back/service_db/internal/offices"
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
	r.Mount("/database", api)

	v1 := chi.NewRouter()
	api.Mount("/users", v1)
	v2 := chi.NewRouter()
	api.Mount("/offices", v2)

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
	v1.Get("/get/{id}", func(w http.ResponseWriter, r *http.Request) {
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

	v1.Get("/get/login/{login}", func(w http.ResponseWriter, r *http.Request) {
		login := chi.URLParam(r, "login")
		log.Print("z")
		user, err := us.GetUserByLogin(db, login)
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

	v1.Post("/auth", func(w http.ResponseWriter, r *http.Request) {
		type AuthRequest struct {
			Login    string `json:"login"`
			Password string `json:"password"`
		}
		var req AuthRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "cannot parse JSON", http.StatusBadRequest)
			return
		}
		authenticated, err := us.AuthUser(db, req.Login, req.Password)
		if err != nil {
			log.Print("1")
			log.Print(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if !authenticated {
			http.Error(w, "invalid login or password", http.StatusUnauthorized)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"message": "authentication successful"})
	})

	v1.Get("/isadmin/{userid}/{officeid}", func(w http.ResponseWriter, r *http.Request) {
		useridStr := chi.URLParam(r, "userid")
		officeidStr := chi.URLParam(r, "officeid")
		userid, err := strconv.Atoi(useridStr)
		if err != nil {
			log.Print(err)
			http.Error(w, "invalid userid", http.StatusBadRequest)
			return
		}
		officeid, err := strconv.Atoi(officeidStr)
		if err != nil {
			log.Print(err)
			http.Error(w, "invalid officeid", http.StatusBadRequest)
			return
		}
		isAdmin, err := us.IsAdmin(db, userid, officeid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]bool{"isadmin": isAdmin})
	})

	// Make user admin
	v1.Post("/makeadmin", func(w http.ResponseWriter, r *http.Request) {
		type MakeAdminRequest struct {
			UserID    int `json:"userid"`
			OfficeID  int `json:"officeid"`
		}
		var req MakeAdminRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "cannot parse JSON", http.StatusBadRequest)
			return
		}
		if err := us.MakeAdmin(db, req.UserID, req.OfficeID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"message": "user made admin successfully"})
	})

	// Create a new office
	v2.Post("/", func(w http.ResponseWriter, r *http.Request) {
		office := new(of.Office)
		if err := json.NewDecoder(r.Body).Decode(office); err != nil {
			http.Error(w, "cannot parse JSON", http.StatusBadRequest)
			return
		}
		if err := of.CreateOffice(db, office); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(office)
	})

	// Get an office by ID
	v2.Get("/getss/{id}", func(w http.ResponseWriter, r *http.Request) {
		log.Print("q")
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Print(err)
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		office, err := of.GetOfficeByID(db, id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				http.Error(w, "office not found", http.StatusNotFound)
				return
			}
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(office)
	})

	// Update an office by ID
	v2.Put("/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Print(err)
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		updatedOffice := new(of.Office)
		if err := json.NewDecoder(r.Body).Decode(updatedOffice); err != nil {
			http.Error(w, "cannot parse JSON", http.StatusBadRequest)
			return
		}
		if err := of.UpdateOffice(db, id, updatedOffice); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(updatedOffice)
	})

	// Delete an office by ID
	v2.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Print(err)
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		if err := of.DeleteOffice(db, id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"status": "office deleted"})
	})

	v2.Get("/gets", func(w http.ResponseWriter, r *http.Request) {
		log.Print("z")
		offices, err := us.GetAllOffices(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(offices)
	})

	// Get all users of an office
	v2.Get("/offices/{officeid}/users", func(w http.ResponseWriter, r *http.Request) {
		log.Print("t")
		officeidStr := chi.URLParam(r, "officeid")
		officeid, err := strconv.Atoi(officeidStr)
		if err != nil {
			log.Print(err)
			http.Error(w, "invalid officeid", http.StatusBadRequest)
			return
		}
		users, err := us.GetUsersByOfficeID(db, officeid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(users)
	})

	// Get the office of a user
	v2.Get("/tmp/{userid}/office", func(w http.ResponseWriter, r *http.Request) {
		log.Print("v")
		useridStr := chi.URLParam(r, "userid")
		userid, err := strconv.Atoi(useridStr)
		if err != nil {
			log.Print(err)
			http.Error(w, "invalid userid", http.StatusBadRequest)
			return
		}
		office, err := us.GetOfficeByUserID(db, userid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(office)
	})

	// Get all users of an office
	v2.Get("/offices/{officeid}/users", func(w http.ResponseWriter, r *http.Request) {
		log.Print("y")
		officeidStr := chi.URLParam(r, "officeid")
		officeid, err := strconv.Atoi(officeidStr)
		if err != nil {
			log.Print(err)
			http.Error(w, "invalid officeid", http.StatusBadRequest)
			return
		}
		users, err := us.GetUsersByOfficeID(db, officeid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(users)
	})

	// Get the office of a user
	v2.Get("/get/{userid}/office", func(w http.ResponseWriter, r *http.Request) {
		log.Print("zq")
		useridStr := chi.URLParam(r, "userid")
		userid, err := strconv.Atoi(useridStr)
		if err != nil {
			log.Print(err)
			http.Error(w, "invalid userid", http.StatusBadRequest)
			return
		}
		office, err := us.GetOfficeByUserID(db, userid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(office)
	})

	return http.ListenAndServe(":8003", r)
}
