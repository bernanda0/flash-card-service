-- name: GetDeck :one
SELECT * FROM deck
WHERE deck_id = $1 LIMIT 1;

-- name: ListDecksByAccount :many
SELECT * FROM deck
WHERE account_id = $1
ORDER BY title;

-- name: CreateDeck :one
INSERT INTO deck (
  account_id, title
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateDeckTitle :exec
UPDATE deck SET title = $1
WHERE deck_id = $2;

-- name: DeleteDeck :exec
DELETE FROM deck
WHERE deck_id = $1;
