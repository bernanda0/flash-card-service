package handlers

import (
	"br/simple-service/db/sqlc"
	"br/simple-service/token"
	"log"
	"net/http"
)

type Handler struct {
	l *log.Logger
	q *sqlc.Queries
	c *uint
}

type AuthHandler struct {
	h *Handler
	t token.Maker
}

type AccountHandler struct {
	h *Handler
}

type DeckHandler struct {
	h *Handler
}

type CardHandler struct {
	h *Handler
}

type HandlerParam struct {
	w           http.ResponseWriter
	r           *http.Request
	method      string
	handlerFunc func(http.ResponseWriter, *http.Request) error
}

type LoginUserResponse struct {
	AccessToken string `json:"access_token"`
	UserID      uint   `json:"user_id"`
	Username    string `json:"username"`
}
