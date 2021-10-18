package controllers

import (
	"net/http"
)

func RegisterControllers() {
	nc := newNumbersController()
	cc := newCharacterController()
	tc := newTextController()

	http.HandleFunc("/api/text", tc.TextController)
	http.HandleFunc("/api/text/", tc.TextController)
	http.HandleFunc("/api/character", cc.CharacterController)
	http.HandleFunc("/api/character/", cc.CharacterController)
	http.HandleFunc("/api/number", nc.NumbersController)
	http.HandleFunc("/api/number/", nc.NumbersController)

	http.HandleFunc("/api/input", HandleInput)
	http.HandleFunc("/api/input/", HandleInput)

}
