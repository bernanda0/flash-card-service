package handlers

import (
	"br/simple-service/db/sqlc"
	"log"
	"net/http"
	"strconv"
)

func NewDeckHandler(l *log.Logger, q *sqlc.Queries) *DeckHandler {
	var c uint = 0
	return &DeckHandler{&Handler{l, q, &c}}
}

func (dh *DeckHandler) CreateDeckH(w http.ResponseWriter, r *http.Request) {
	hp := HandlerParam{w, r, http.MethodPost, dh.createDeck}
	dh.h.handleRequest(hp)
}

func (dh *DeckHandler) DeleteDeckH(w http.ResponseWriter, r *http.Request) {
	hp := HandlerParam{w, r, http.MethodDelete, dh.deleteDeck}
	dh.h.handleRequest(hp)
}

func (dh *DeckHandler) GetDeckH(w http.ResponseWriter, r *http.Request) {
	hp := HandlerParam{w, r, http.MethodGet, dh.getDeck}
	dh.h.handleRequest(hp)
}

func (dh *DeckHandler) ListDecksByAccountH(w http.ResponseWriter, r *http.Request) {
	hp := HandlerParam{w, r, http.MethodGet, dh.listDecksByAccount}
	dh.h.handleRequest(hp)
}

func (dh *DeckHandler) UpdateDeckTitleH(w http.ResponseWriter, r *http.Request) {
	hp := HandlerParam{w, r, http.MethodPut, dh.updateDeckTitle}
	dh.h.handleRequest(hp)
}

// implementation
func (dh *DeckHandler) createDeck(w http.ResponseWriter, r *http.Request) error {
	// Parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return err
	}

	// Retrieve form values
	accountID, err := strconv.Atoi(r.FormValue("account_id"))
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return err
	}
	title := r.FormValue("title")

	deckParam := sqlc.CreateDeckParams{
		AccountID: int32(accountID),
		Title:     title,
	}

	deck, err := dh.h.q.CreateDeck(r.Context(), deckParam)
	if err != nil {
		http.Error(w, "Error creating deck", http.StatusInternalServerError)
		return err
	}

	w.WriteHeader(http.StatusCreated)
	toJSON(w, deck)
	return nil
}

func (dh *DeckHandler) deleteDeck(w http.ResponseWriter, r *http.Request) error {
	deckID, err := strconv.Atoi(r.URL.Query().Get("deck_id"))
	if err != nil {
		http.Error(w, "Invalid deck ID", http.StatusBadRequest)
		return err
	}

	deck, err := dh.h.q.DeleteDeck(r.Context(), int32(deckID))
	if err != nil {
		http.Error(w, "Error deleting deck", http.StatusInternalServerError)
		return err
	}

	w.WriteHeader(http.StatusOK)
	toJSON(w, deck)
	return nil
}

func (dh *DeckHandler) getDeck(w http.ResponseWriter, r *http.Request) error {
	deckID, err := strconv.Atoi(r.URL.Query().Get("deck_id"))
	if err != nil {
		http.Error(w, "Invalid deck ID", http.StatusBadRequest)
		return err
	}

	deck, err := dh.h.q.GetDeck(r.Context(), int32(deckID))
	if err != nil {
		http.Error(w, "Deck not found", http.StatusInternalServerError)
		return err
	}

	w.WriteHeader(http.StatusOK)
	toJSON(w, deck)
	return nil
}

func (dh *DeckHandler) listDecksByAccount(w http.ResponseWriter, r *http.Request) error {
	accountID, err := strconv.Atoi(r.URL.Query().Get("account_id"))
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return err
	}

	decks, err := dh.h.q.ListDecksByAccount(r.Context(), int32(accountID))
	if err != nil {
		http.Error(w, "Cannot retrieve deck", http.StatusInternalServerError)
		return err
	}

	w.WriteHeader(http.StatusOK)
	toJSON(w, decks)
	return nil
}

func (dh *DeckHandler) updateDeckTitle(w http.ResponseWriter, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return err
	}

	deckID, err := strconv.Atoi(r.FormValue("deck_id"))
	if err != nil {
		http.Error(w, "Invalid deck ID", http.StatusBadRequest)
		return err
	}
	title := r.FormValue("title")

	updateDeckParam := sqlc.UpdateDeckTitleParams{
		Title:  title,
		DeckID: int32(deckID),
	}

	deck, err := dh.h.q.UpdateDeckTitle(r.Context(), updateDeckParam)
	if err != nil {
		http.Error(w, "Cannot update deck", http.StatusInternalServerError)
		return err
	}

	w.WriteHeader(http.StatusOK)
	toJSON(w, deck)
	return nil
}
