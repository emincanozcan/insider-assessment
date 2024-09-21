-- name: CreateMessage :one
INSERT INTO messages (content, recipient, status)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetPendingMessagesAndMarkAsSending :many
WITH updated AS (
  SELECT id
  FROM messages
  WHERE status = 0
  ORDER BY id ASC
  LIMIT $1
)
UPDATE messages
SET status = 1
WHERE id IN (SELECT id FROM updated)
RETURNING *;

-- name: GetMessageByID :one
SELECT * FROM messages
WHERE id = $1 LIMIT 1;

-- name: ListPendingMessages :many
SELECT * FROM messages
WHERE status = 0
ORDER BY id
LIMIT $1;

-- name: ListProcessingMessages :many
SELECT * FROM messages
WHERE status = 1
ORDER BY id DESC
LIMIT $1;

-- name: ListSentMessages :many
SELECT * FROM messages
WHERE status = 2
ORDER BY id DESC
LIMIT $1;

-- name: UpdateMessageStatus :exec
UPDATE messages
SET status = $2
WHERE id = $1;

