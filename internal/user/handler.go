package user

import (
	"github.com/arxon31/user-service/internal/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

const (
	usersURL = "/users"
	userURL  = "/users/{uuid}"
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}
func (h *handler) Register(router *chi.Mux) {
	router.Get(usersURL, h.GetListOfUsersHandler)
	router.Get(userURL, h.GetUserByUUIDHandler)
	router.Post(userURL, h.CreateNewUserUserHandler)
	router.Put(userURL, h.UpdateUserHandler)
	router.Patch(userURL, h.PartiallyUpdateUserHandler)
	router.Delete(userURL, h.DeleteUserHandler)
}

func (h *handler) GetListOfUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is list of users"))
}

func (h *handler) GetUserByUUIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is user by uuid"))
}

func (h *handler) CreateNewUserUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is create new user"))
}

func (h *handler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is update user"))
}

func (h *handler) PartiallyUpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is partially update user"))
}

func (h *handler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is delete user"))
}
