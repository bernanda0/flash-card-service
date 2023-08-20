package handlers

import (
	"br/simple-service/db/sqlc"
	"br/simple-service/utils"
	"log"
	"net/http"
	"strconv"
)

type AccountHandler struct {
	l *log.Logger
	q *sqlc.Queries
	c *uint
}

func NewAccountHandler(l *log.Logger, q *sqlc.Queries) *AccountHandler {
	var c uint = 0
	return &AccountHandler{l, q, &c}
}

func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	*h.c++
	var err error = nil
	defer func() {
		apiLog(h.l, h.c, &r.RequestURI, err)
	}()

	err = checkHTTPMethod(w, r.Method, http.MethodPost)
	if err != nil {
		return
	}

	// Parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	// Retrieve form values
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	hashedPassword, _ := utils.HashPassword(password)

	// Create accountParams using retrieved form values
	accountParams := sqlc.CreateAccountParams{
		Username:     username,
		Email:        email,
		PasswordHash: hashedPassword, // Don't forget to hash the password
	}

	account, err := h.q.CreateAccount(r.Context(), accountParams)
	if err != nil {
		http.Error(w, "Error creating account", http.StatusInternalServerError)
		return
	}

	defer apiLog(h.l, h.c, &r.RequestURI, err)

	w.WriteHeader(http.StatusCreated)
	toJSON(w, account)
}

func (h *AccountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	*h.c++
	var err error = nil
	defer func() {
		apiLog(h.l, h.c, &r.RequestURI, err)
	}()

	err = checkHTTPMethod(w, r.Method, http.MethodGet)
	if err != nil {
		return
	}

	accountID, err := strconv.Atoi(r.URL.Query().Get("account_id"))
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	account, err := h.q.GetAccount(r.Context(), int32(accountID))
	if err != nil {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	toJSON(w, account)
}

func (h *AccountHandler) ListAccounts(w http.ResponseWriter, r *http.Request) {
	*h.c++
	var err error = nil
	defer func() {
		apiLog(h.l, h.c, &r.RequestURI, err)
	}()

	err = checkHTTPMethod(w, r.Method, http.MethodGet)
	if err != nil {
		return
	}

	accounts, err := h.q.ListAccounts(r.Context())
	if err != nil {
		h.l.Println(err)
		http.Error(w, "Error listing accounts", http.StatusInternalServerError)
		return
	}

	toJSON(w, accounts)
}

func (h *AccountHandler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	var err error = nil
	defer func() {
		apiLog(h.l, h.c, &r.RequestURI, err)
	}()

	err = checkHTTPMethod(w, r.Method, http.MethodDelete)
	if err != nil {
		return
	}

	accountID, err := strconv.Atoi(r.URL.Query().Get("account_id"))
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	account, err := h.q.DeleteAccount(r.Context(), int32(accountID))
	if err != nil {
		http.Error(w, "Error deleting account", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	toJSON(w, account)
}
