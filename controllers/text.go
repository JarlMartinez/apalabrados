package controllers

import (
	"apalabrados/models"
	"encoding/json"
	"net/http"
	"strings"
)

type textController struct{}

func newTextController() *textController {
	return &textController{}
}

func (tc textController) TextController(w http.ResponseWriter, r *http.Request) {
	url_parts := strings.Split(r.URL.Path, "/")
	if r.URL.Path == "/api/text" || r.URL.Path == "/api/text/" {
		if r.Method == http.MethodGet {
			tc.getAll(w)
		} else {
			w.WriteHeader(http.StatusNotImplemented)
		}
		return
	}
	id := url_parts[3]
	if id != "" {
		if r.Method == http.MethodDelete {
			tc.delete(id, w)
		} else {
			w.WriteHeader(http.StatusNotImplemented)
		}
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Seems unhandled path error"))
}

func (tc *textController) getAll(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	texts, err := models.GetTexts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(texts)
}

func (tc *textController) delete(id string, w http.ResponseWriter) {

	err := models.DeleteText(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusNoContent)
}
