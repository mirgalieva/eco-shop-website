package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"mirgalievaal-project/backend/internal/entity"
	"mirgalievaal-project/backend/pkg/helpers"
	"net/http"
	"strconv"
)

func (h *Handler) getAllUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		users, err := h.service.User.GetAll()
		if err != nil {
			log.Printf("error while getting all users, error:=%v\n", err.Error())
			w.WriteHeader(http.StatusTeapot)
			return
		}
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(users)
		if err != nil {
			log.Printf("Error while encoding users, error:%v", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) getUserByID(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		vars := mux.Vars(r)
		userId, err := strconv.ParseUint(vars["user_id"], 10, 32)
		if err != nil {
			log.Printf("Error while pasing user_id to uint. Eroror:%v\n", err.Error())
			return
		}
		user, err := h.service.User.Get(uint(userId))
		if err != nil {
			log.Printf("Error when getting user with id=%v, error:%v\n", userId, err.Error())
			w.WriteHeader(http.StatusForbidden)
		}
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			log.Printf("Error while encoding users, error:%v", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) registerUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		defer helpers.Untrace(helpers.Trace("handler.registerUser "))
		userReg := entity.UserRegister{}
		err := json.NewDecoder(r.Body).Decode(&userReg)
		if err != nil {
			log.Printf("error while trying to descode useReg to register a new user, error:=%v\n", err.Error())
			w.WriteHeader(http.StatusTeapot)
			return
		}
		err = h.service.User.Register(&userReg)
		if err != nil {
			log.Printf("error while getting all users, error:=%v\n", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			jsonResp, _ := json.Marshal(NewResponse(responseStatusError, "", err.Error()))
			w.Write(jsonResp)
			return
		}
		w.WriteHeader(http.StatusCreated)

	}
}
