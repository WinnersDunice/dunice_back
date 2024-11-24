package handler

import (
	"net/http"

	"github.com/WinnersDunice/dunice_back/proxy/pkg/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Handler struct {
	//server - config the http server
	Server *service.Server
}

const (
	IP        = "195.80.238.9"
	UsersPort = "8003"
	MongoPort = "8004"
	SSOPort   = "8530"
)

func SetCORSOriginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Set("Access-Control-Allow-Origin", "https://dunicewinners.ru, https://r.dunicewinners.ru")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type RegRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Mac_address string `json:"mac_address"`
}

// Constructor of a handler
func NewHandler(server *service.Server) *Handler {
	return &Handler{Server: server}
}

func (h *Handler) InitRoutes() *chi.Mux {
	z := NewUserStatsHandler();
	/////////////////////////////////////////////////////////////////////////////////////////////
	//init new router
	r := chi.NewRouter()
	// redirect /auth/ to /auth
	r.Use(middleware.RedirectSlashes)
	//serve all the r-routes

	r.Use(SetCORSOriginMiddleware)

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/login", (h.Login))

	r.Post("/reg", (h.Register))

	r.Post("/user_info/{mac}", (z.SetUserInfo))
	r.Get("/user_info/{mac}", (z.GetUserInfo))
	r.Post("/tables", h.CreateTable)
	r.Post("/chairs", h.CreateChair)
	r.Post("/equipment", h.CreateEquipment)
	r.Post("/furniture", h.CreateFurniture)
	r.Post("/kitchen", h.CreateKitchen)
	r.Get("/belongsTo/{belongsTo}", h.GetObjectsByBelongsTo)
	r.Get("/officeId/{officeId}", h.GetObjectsByOfficeID)
	r.Delete("/{collection}/{id}", h.DeleteObjectByID)
	r.Delete("/logout", (h.Logout))
	

	r.Post("/users", h.CreateUser)
	r.Get("/users/{id}", h.GetUserByID)
	r.Put("/users/{id}", h.UpdateUser)
	r.Delete("/users/{id}", h.DeleteUser)
	r.Get("/users/login/{login}", h.GetUserByLogin)
	r.Put("/users/login/{id}", h.UpdateUserLogin)
	r.Put("/users/password/{id}", h.UpdateUserPassword)
	r.Put("/users/macaddress/{id}", h.UpdateUserMacAddress)
	r.Get("/users/login2/{id}", h.GetUserLogin)
	r.Get("/users/password/{id}", h.GetUserPassword)
	r.Post("/users/auth", h.AuthUser)
	r.Get("/users/isadmin/{userid}/{officeid}", h.IsAdmin)
	r.Post("/users/makeadmin", h.MakeAdmin)
	r.Get("/get/offices", h.GetAllOffices)
	r.Get("/offices/{officeid}/users", h.GetUsersByOfficeID)
	r.Get("/users/{userid}/office", h.GetOfficeByUserID)
	r.Post("/offices", h.CreateOffice)
	r.Get("/offices/{id}", h.GetOfficeByID)
	r.Put("/offices/{id}", h.UpdateOffice)
	r.Delete("/offices/{id}", h.DeleteOffice)
	return r

}
