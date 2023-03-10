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
ORDER BY created_at desc
LIMIT $2
OFFSET $3;

-- name: ListTransfersToAccount :many
SELECT * FROM transfers
WHERE to_account_id = $1
ORDER BY created_at desc
LIMIT $2
OFFSET $3;

-- name: ListTransfersInvolvingAccount :many
SELECT * FROM transfers
WHERE from_account_id = $1
OR to_account_id = $1
ORDER BY created_at desc
LIMIT $2
OFFSET $3;

-- name: UpdateTransferToAccount :one
UPDATE transfers 
SET to_account_id = $2
WHERE id = $1
RETURNING *;

-- name: UpdateTransferFromAccount :one
UPDATE transfers 
SET from_account_id = $2
WHERE id = $1
RETURNING *;

-- name: UpdateTransferAmount :one
UPDATE transfers 
SET amount = $2
WHERE id = $1
RETURNING *;

-- name: DeleteTransfer :exec
DELETE FROM transfers 
WHERE id = $1;