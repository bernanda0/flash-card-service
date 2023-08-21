package handlers

import (
	"br/simple-service/db/sqlc"
	"log"
	"net/http"
)

type Handler struct {
	l *log.Logger
	q *sqlc.Queries
	c *uint
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
