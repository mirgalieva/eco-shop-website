package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"mirgalievaal-project/backend/internal/service"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", index).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/users", h.getAllUsers).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/users/{user_id}", h.getUserByID).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/users/register", h.registerUser).Methods(http.MethodPost, http.MethodOptions)

	r.HandleFunc("/products", h.getAllProducts).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/products/add", h.addProduct).Methods(http.MethodPost, http.MethodOptions)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(CustomCORSMiddleware(r))

	return r
}
func CustomCORSMiddleware(r *mux.Router) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Length, Content-Type, Authorization, Host, Origin, X-CSRF-Token")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization")

			next.ServeHTTP(w, req)
		})
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world!\n")
}
