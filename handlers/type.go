package handlers

import (
	"br/simple-service/db/sqlc"
	"br/simple-service/token"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Handler struct {
	l *log.Logger
	q *sqlc.Queries
	c *uint
	u *AuthedUser
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

type AuthedUser struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
}

type LoginUserResponse struct {
	SessionID      uuid.UUID `json:"session_id"`
	AccessToken    string    `json:"access_token"`
	AccessTokenEx  time.Time `json:"access_token_expire"`
	RefreshToken   string    `json:"refresh_token"`
	RefreshTokenEx time.Time `json:"refresh_token_expire"`
	UserID         uint      `json:"user_id"`
	Username       string    `json:"username"`
}

type renewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type renewAccessTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expire"`
}
