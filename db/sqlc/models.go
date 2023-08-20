// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package sqlc

import (
	"database/sql"
)

type Account struct {
	AccountID    int32        `json:"account_id"`
	Username     string       `json:"username"`
	Email        string       `json:"email"`
	PasswordHash string       `json:"password_hash"`
	CreatedAt    sql.NullTime `json:"created_at"`
}

type Deck struct {
	DeckID    int32         `json:"deck_id"`
	AccountID sql.NullInt32 `json:"account_id"`
	Title     string        `json:"title"`
	CreatedAt sql.NullTime  `json:"created_at"`
}

type Flashcard struct {
	FlashcardID    int32           `json:"flashcard_id"`
	DeckID         sql.NullInt32   `json:"deck_id"`
	Question       string          `json:"question"`
	Answer         string          `json:"answer"`
	NextReviewDate sql.NullTime    `json:"next_review_date"`
	Interval       sql.NullInt32   `json:"interval"`
	Repetitions    sql.NullInt32   `json:"repetitions"`
	EasinessFactor sql.NullFloat64 `json:"easiness_factor"`
	CreatedAt      sql.NullTime    `json:"created_at"`
	UpdatedAt      sql.NullTime    `json:"updated_at"`
	IsArchived     sql.NullBool    `json:"is_archived"`
}
