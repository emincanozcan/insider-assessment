-- name: CreateMessage :one
INSERT INTO messages (content, recipient)
VALUES ($1, $2)
RETURNING *;

-- name: GetPendingMessagesAndMarkAsSending :many
WITH updated AS (
  SELECT id
  FROM messages
  WHERE ((sending_at IS NULL OR sending_at < NOW() - INTERVAL '1 hour') AND tries < 3)
  ORDER BY created_at ASC 
  LIMIT $1
) UPDATE messages
SET sending_at = NOW(), tries = tries + 1
WHERE id IN (SELECT id FROM updated)
RETURNING *;

-- name: ListSentMessages :many
SELECT * FROM messages
WHERE sent_at IS NOT NULL
ORDER BY sent_at DESC
LIMIT $1;

-- name: MarkMessageAsSent :exec
UPDATE messages
SET sent_at = NOW()
WHERE id = $1;

-- name: MarkMessageAsNotSent :exec
UPDATE messages
SET sending_at = null, sent_at = NULL
WHERE id = $1;

