package main

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// handlers

// CreateUserHandler ...
func CreateUserHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "Create User Success")
}

// LoginHandler ...
func LoginHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	username := p.ByName("username")
	io.WriteString(w, username)
}
