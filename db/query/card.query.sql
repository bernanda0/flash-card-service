-- name: GetFlashcard :one
SELECT * FROM flashcard
WHERE flashcard_id = $1 LIMIT 1;

-- name: ListFlashcardsByDeck :many
SELECT * FROM flashcard
WHERE deck_id = $1
ORDER BY created_at;

-- name: CreateFlashcard :one
INSERT INTO flashcard (
  deck_id, question, answer, next_review_date,
  interval, repetitions, easiness_factor
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: UpdateFlashcard :exec
UPDATE flashcard
SET question = $1, answer = $2, next_review_date = $3,
    interval = $4, repetitions = $5, easiness_factor = $6,
    updated_at = NOW()
WHERE flashcard_id = $7;

-- name: DeleteFlashcard :exec
DELETE FROM flashcard
WHERE flashcard_id = $1;
