-- name: CreateTransfer :one
INSERT INTO transfers (
  from_account_id, to_account_id, amount
) VALUES (
  $1, $2, $3
) 
RETURNING * ;

-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: ListTransfersFromAccount :many
SELECT * FROM transfers
WHERE from_account_id = $1
LIMIT $1
OFFSET $2;

-- name: ListTransfersToAccount :many
SELECT * FROM transfers
WHERE to_account_id = $1
LIMIT $1
OFFSET $2;

-- name: ListTransfersInvolvingAccount :many
SELECT * FROM transfers
WHERE from_account_id = $1
OR to_account_id = $1
LIMIT $1
OFFSET $2;
