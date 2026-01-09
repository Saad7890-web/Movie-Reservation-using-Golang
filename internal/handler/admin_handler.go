package handler

import "net/http"

type AdminHandler struct{}

func NewAdminHandler()*AdminHandler{
	return &AdminHandler{}
}

func(h *AdminHandler)Dashboard(w http.ResponseWriter, r*http.Request){
	w.Write([]byte("Welcome Admin"))
}