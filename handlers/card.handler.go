package handlers

import (
	"br/simple-service/db/sqlc"
	"log"
	"net/http"
	"strconv"
)

func NewCardHandler(l *log.Logger, q *sqlc.Queries) *CardHandler {
	var c uint = 0
	return &CardHandler{&Handler{l, q, &c}}
}

func (ch *CardHandler) CreateCardH(w http.ResponseWriter, r *http.Request) {
	hp := HandlerParam{w, r, http.MethodPost, ch.createCard}
	ch.h.handleRequest(hp)
}

func (ch *CardHandler) GetCardH(w http.ResponseWriter, r *http.Request) {
	hp := HandlerParam{w, r, http.MethodGet, ch.getCard}
	ch.h.handleRequest(hp)
}

func (ch *CardHandler) ListCardsH(w http.ResponseWriter, r *http.Request) {
	hp := HandlerParam{w, r, http.MethodGet, ch.listCards}
	ch.h.handleRequest(hp)
}

func (ch *CardHandler) UpdateCardH(w http.ResponseWriter, r *http.Request) {
	hp := HandlerParam{w, r, http.MethodPut, ch.updateCard}
	ch.h.handleRequest(hp)
}

func (ch *CardHandler) DeleteCardH(w http.ResponseWriter, r *http.Request) {
	hp := HandlerParam{w, r, http.MethodDelete, ch.deleteCard}
	ch.h.handleRequest(hp)
}

// the implementation

func (ch *CardHandler) createCard(w http.ResponseWriter, r *http.Request) error {
	// Parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return err
	}

	// Retrieve form values
	deckID, err := strconv.Atoi(r.FormValue("deck_id"))
	if err != nil {
		http.Error(w, "Invalid deck ID", http.StatusBadRequest)
		return err
	}
	question := r.FormValue("question")
	answer := r.FormValue("answer")

	// Create cardParams using retrieved form values
	cardParams := sqlc.CreateFlashcardParams{
		DeckID:   int32(deckID),
		Question: question,
		Answer:   answer,
	}

	card, err := ch.h.q.CreateFlashcard(r.Context(), cardParams)
	if err != nil {
		http.Error(w, "Error creating card", http.StatusInternalServerError)
		return err
	}

	w.WriteHeader(http.StatusCreated)
	toJSON(w, card)
	return nil
}

func (ch *CardHandler) getCard(w http.ResponseWriter, r *http.Request) error {
	cardID, err := strconv.Atoi(r.URL.Query().Get("card_id"))
	if err != nil {
		http.Error(w, "Invalid card ID", http.StatusBadRequest)
		return err
	}

	card, err := ch.h.q.GetFlashcard(r.Context(), int32(cardID))
	if err != nil {
		http.Error(w, "Card not found", http.StatusNotFound)
		return err
	}

	toJSON(w, card)
	return nil
}

func (ch *CardHandler) listCards(w http.ResponseWriter, r *http.Request) error {
	deckID, err := strconv.Atoi(r.URL.Query().Get("deck_id"))
	if err != nil {
		http.Error(w, "Invalid deck ID", http.StatusBadRequest)
		return err
	}

	cards, err := ch.h.q.ListFlashcardsByDeck(r.Context(), int32(deckID))
	if err != nil {
		http.Error(w, "Error listing cards", http.StatusInternalServerError)
		return err
	}

	toJSON(w, cards)
	return nil
}

func (ch *CardHandler) updateCard(w http.ResponseWriter, r *http.Request) error {
	// Parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return err
	}

	// Retrieve form values
	cardID, err := strconv.Atoi(r.FormValue("card_id"))
	if err != nil {
		http.Error(w, "Invalid card ID", http.StatusBadRequest)
		return err
	}
	question := r.FormValue("question")
	answer := r.FormValue("answer")

	// Create cardParams using retrieved form values
	cardParams := sqlc.UpdateFlashcardParams{
		Question:    question,
		Answer:      answer,
		FlashcardID: int32(cardID),
	}

	card, err := ch.h.q.UpdateFlashcard(r.Context(), cardParams)
	if err != nil {
		http.Error(w, "Error updating card", http.StatusInternalServerError)
		return err
	}

	toJSON(w, card)
	return nil
}

func (ch *CardHandler) deleteCard(w http.ResponseWriter, r *http.Request) error {
	cardID, err := strconv.Atoi(r.URL.Query().Get("card_id"))
	if err != nil {
		http.Error(w, "Invalid card ID", http.StatusBadRequest)
		return err
	}

	card, err := ch.h.q.DeleteFlashcard(r.Context(), int32(cardID))
	if err != nil {
		http.Error(w, "Error deleting card", http.StatusInternalServerError)
		return err
	}

	w.WriteHeader(http.StatusOK)
	toJSON(w, card)
	return nil
}
