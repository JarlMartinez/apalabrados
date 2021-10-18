package controllers

import (
	"apalabrados/models"
	"encoding/json"
	"net/http"
	"strings"
)

type numberController struct{}

func newNumbersController() *numberController {
	return &numberController{}
}

func (nc numberController) NumbersController(w http.ResponseWriter, r *http.Request) {
	url_parts := strings.Split(r.URL.Path, "/")
	if r.URL.Path == "/api/number" || r.URL.Path == "/api/number/" {
		if r.Method == http.MethodGet {
			nc.getAll(w)
		} else {
			w.WriteHeader(http.StatusNotImplemented)
		}
		return
	}
	id := url_parts[3]
	if id != "" {
		if r.Method == http.MethodDelete {
			nc.delete(id, w)
		} else {
			w.WriteHeader(http.StatusNotImplemented)
		}
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Seems unhandled path error"))
}

func (nc *numberController) getAll(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	numbers, err := models.GetNumbers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(numbers)
}

func (nc *numberController) delete(id string, w http.ResponseWriter) {

	err := models.DeleteNumber(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusNoContent)
}
