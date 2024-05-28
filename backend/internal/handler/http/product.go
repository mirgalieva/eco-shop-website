package handler

import (
	"encoding/json"
	"log"
	"mirgalievaal-project/backend/internal/entity"
	"net/http"
)

func (h *Handler) getAllProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		products, err := h.service.Product.GetAll()
		if err != nil {
			log.Printf("error while getting all products, error:=%v\n", err.Error())
			w.WriteHeader(http.StatusTeapot)
			return
		}
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(products)
		if err != nil {
			log.Printf("Error while encoding products, error:%v", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) addProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		product := &entity.Product{}
		err := json.NewDecoder(r.Body).Decode(product)
		if err != nil {
			log.Printf("error while trying to descode useReg to register a new user, error:=%v\n", err.Error())
			w.WriteHeader(http.StatusTeapot)
			return
		}
		product, err = h.service.Product.Create(product)
		if err != nil {
			log.Printf("error while getting all users, error:=%v\n", err.Error())
			w.WriteHeader(http.StatusTeapot)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}
