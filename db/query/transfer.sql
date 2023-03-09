-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: GetTransferForUpdate :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListTransfer :many
SELECT * FROM transfers
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateTransfer :one
INSERT INTO transfers (
  from_account_id, to_account_id, amount, currency
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateTransfer :one
UPDATE transfers
SET amount = $2
WHERE id = $1
RETURNING *;

-- name: DeleteTransfer :exec
DELETE FROM transfers
WHERE id = $1;
