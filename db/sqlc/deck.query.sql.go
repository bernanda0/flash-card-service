// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: deck.query.sql

package sqlc

import (
	"context"
)

const createDeck = `-- name: CreateDeck :one
INSERT INTO deck (
  account_id, title
) VALUES (
  $1, $2
)
RETURNING deck_id, account_id, title, created_at
`

type CreateDeckParams struct {
	AccountID int32  `json:"account_id"`
	Title     string `json:"title"`
}

func (q *Queries) CreateDeck(ctx context.Context, arg CreateDeckParams) (Deck, error) {
	row := q.db.QueryRowContext(ctx, createDeck, arg.AccountID, arg.Title)
	var i Deck
	err := row.Scan(
		&i.DeckID,
		&i.AccountID,
		&i.Title,
		&i.CreatedAt,
	)
	return i, err
}

const deleteDeck = `-- name: DeleteDeck :one
DELETE FROM deck
WHERE deck_id = $1
RETURNING deck_id, account_id, title, created_at
`

func (q *Queries) DeleteDeck(ctx context.Context, deckID int32) (Deck, error) {
	row := q.db.QueryRowContext(ctx, deleteDeck, deckID)
	var i Deck
	err := row.Scan(
		&i.DeckID,
		&i.AccountID,
		&i.Title,
		&i.CreatedAt,
	)
	return i, err
}

const getDeck = `-- name: GetDeck :one
SELECT deck_id, account_id, title, created_at FROM deck
WHERE deck_id = $1 LIMIT 1
`

func (q *Queries) GetDeck(ctx context.Context, deckID int32) (Deck, error) {
	row := q.db.QueryRowContext(ctx, getDeck, deckID)
	var i Deck
	err := row.Scan(
		&i.DeckID,
		&i.AccountID,
		&i.Title,
		&i.CreatedAt,
	)
	return i, err
}

const getOwner = `-- name: GetOwner :one
SELECT account_id FROM deck
WHERE deck_id = $1
`

func (q *Queries) GetOwner(ctx context.Context, deckID int32) (int32, error) {
	row := q.db.QueryRowContext(ctx, getOwner, deckID)
	var account_id int32
	err := row.Scan(&account_id)
	return account_id, err
}

const listDecksByAccount = `-- name: ListDecksByAccount :many
SELECT deck_id, account_id, title, created_at FROM deck
WHERE account_id = $1
ORDER BY title
`

func (q *Queries) ListDecksByAccount(ctx context.Context, accountID int32) ([]Deck, error) {
	rows, err := q.db.QueryContext(ctx, listDecksByAccount, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Deck
	for rows.Next() {
		var i Deck
		if err := rows.Scan(
			&i.DeckID,
			&i.AccountID,
			&i.Title,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateDeckTitle = `-- name: UpdateDeckTitle :one
UPDATE deck SET title = $1
WHERE deck_id = $2
RETURNING deck_id, account_id, title, created_at
`

type UpdateDeckTitleParams struct {
	Title  string `json:"title"`
	DeckID int32  `json:"deck_id"`
}

func (q *Queries) UpdateDeckTitle(ctx context.Context, arg UpdateDeckTitleParams) (Deck, error) {
	row := q.db.QueryRowContext(ctx, updateDeckTitle, arg.Title, arg.DeckID)
	var i Deck
	err := row.Scan(
		&i.DeckID,
		&i.AccountID,
		&i.Title,
		&i.CreatedAt,
	)
	return i, err
}
