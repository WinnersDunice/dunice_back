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
		w.Header().Set("Access-Control-Allow-Origin", "https://winnersdunice.ru")
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
}

// Constructor of a handler
func NewHandler(server *service.Server) *Handler {
	return &Handler{Server: server}
}

func (h *Handler) InitRoutes() *chi.Mux {
	/////////////////////////////////////////////////////////////////////////////////////////////
	//init new router
	r := chi.NewRouter()
	// redirect /auth/ to /auth
	r.Use(middleware.RedirectSlashes)
	//serve all the api-routes

	r.Use(SetCORSOriginMiddleware)

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/login", (h.Login))

	r.Post("/reg", (h.Register))

	r.Post("/tables", h.CreateTable)
	r.Post("/chairs", h.CreateChair)
	r.Post("/equipment", h.CreateEquipment)
	r.Post("/furniture", h.CreateFurniture)
	r.Post("/kitchen", h.CreateKitchen)
	r.Get("/belongsTo/{belongsTo}", h.GetObjectsByBelongsTo)
	r.Get("/officeId/{officeId}", h.GetObjectsByOfficeID)
	r.Delete("/{collection}/{id}", h.DeleteObjectByID)
	r.Delete("/logout", (h.Logout))

	return r

}
