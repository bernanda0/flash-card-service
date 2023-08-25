package handlers

import (
	"br/simple-service/db/sqlc"
	"br/simple-service/token"
	"br/simple-service/utils"
	"errors"
	"log"
	"net/http"
	"time"
)

const (
	DURATION = 15
)

func NewAuthHandler(l *log.Logger, q *sqlc.Queries, t *token.Maker) *AuthHandler {
	var c uint = 0
	return &AuthHandler{&Handler{l, q, &c}, *t}
}

func (auth_h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	hp := HandlerParam{w, r, http.MethodPost, auth_h.login}
	auth_h.h.handleRequest(hp)
}

func (auth_h *AuthHandler) login(w http.ResponseWriter, r *http.Request) error {
	// Parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return err
	}

	// Retrieve form values
	email := r.FormValue("email")
	password := r.FormValue("password")

	ok := utils.EmailIsValid(email)
	if !ok {
		http.Error(w, "Invalid Email", http.StatusInternalServerError)
		return errors.New("invalid email")
	}

	ok = utils.PasswordIsValid(password)
	if !ok {
		http.Error(w, "Invalid Password", http.StatusInternalServerError)
		return errors.New("invalid password")
	}

	// check if user exist
	user, err := auth_h.h.q.GetAccountbyEmail(r.Context(), email)
	if err != nil {
		http.Error(w, "Account not found! Register first", http.StatusInternalServerError)
		return errors.New("account not found, register first")
	}

	duration := time.Minute * DURATION
	token, err := auth_h.t.GenerateToken(uint(user.AccountID), user.Username, duration)
	if err != nil {
		return errors.New("failed generate token for user")
	}

	res := LoginUserResponse{
		AccessToken: token,
		UserID:      uint(user.AccountID),
		Username:    user.Username,
	}

	w.WriteHeader(http.StatusCreated)
	toJSON(w, res)
	return nil
}
