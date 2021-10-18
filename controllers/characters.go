package controllers

import (
	"apalabrados/models"
	"encoding/json"
	"net/http"
	"strings"
)

type characterController struct{}

func newCharacterController() *characterController {
	return &characterController{}
}

func (cc characterController) CharacterController(w http.ResponseWriter, r *http.Request) {
	url_parts := strings.Split(r.URL.Path, "/")
	if r.URL.Path == "/api/character" || r.URL.Path == "/api/character/" {
		if r.Method == http.MethodGet {
			cc.getAll(w)
		} else {
			w.WriteHeader(http.StatusNotImplemented)
		}
		return
	}
	id := url_parts[3]
	if id != "" {
		if r.Method == http.MethodDelete {
			cc.delete(id, w)
		} else {
			w.WriteHeader(http.StatusNotImplemented)
		}
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Seems unhandled path error"))
}

func (cc *characterController) getAll(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	numbers, err := models.GetCharacters()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(numbers)
}

func (cc *characterController) delete(id string, w http.ResponseWriter) {

	err := models.DeleteCharacter(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusNoContent)
}
