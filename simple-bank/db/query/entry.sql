-- name: CreateEntry :one
INSERT INTO entries (
  account_id, amount
) VALUES (
  $1, $2
) 
RETURNING * ;

-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM entries
WHERE account_id = $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;

-- name: UpdateEntryAccount :exec
UPDATE entries
SET account_id = $2
WHERE id = $1;

-- name: UpdateEntryAmount :exec
UPDATE entries
SET amount = $2
WHERE id = $1;

-- name: DeleteEntry :exec
DELETE FROM entries 
WHERE id = $1;