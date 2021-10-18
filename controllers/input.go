package controllers

import (
	"apalabrados/models"
	"encoding/json"
	"net/http"
	"strconv"
)

type Input struct {
	Value string
}

func HandleInput(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	dec := json.NewDecoder(r.Body)
	var input Input
	err := dec.Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed decoding request body"))
		return
	}

	// If its number, save as number
	value, err := strconv.Atoi(input.Value)
	if err == nil {
		newNumber := models.Number{
			Value: value,
		}
		err = models.AddNumber(newNumber)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error saving new number"))
			return
		}
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// If includes some non-letter character, save those.
	hasSomeChar := false
	for _, c := range input.Value {
		if !isAlphaNumeric(c) {
			hasSomeChar = true
			char := models.Chartacter{
				Value: string(c),
			}
			err = models.AddCharacter(char)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Error saving new character"))
				return
			}
		}
	}
	if hasSomeChar {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Save as text
	if input.Value != "" && !hasSomeChar {
		newText := models.Text{
			Value: input.Value,
		}
		err = models.AddText(newText)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error saving new text"))
			return
		}
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Unhandled
	w.WriteHeader(http.StatusNotAcceptable)
}

func isAlphaNumeric(asciiVal rune) bool {
	if (asciiVal >= 65 && asciiVal <= 90) || // capital letter
		(asciiVal >= 97 && asciiVal <= 122) || // small letter
		(asciiVal >= 48 && asciiVal <= 57) { // number
		return true
	}
	return false
}
