// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: message.sql

package sqlc

import (
	"context"
)

const createMessage = `-- name: CreateMessage :one
INSERT INTO messages (content, recipient, status)
VALUES ($1, $2, $3)
RETURNING id, content, recipient, status
`

type CreateMessageParams struct {
	Content   string
	Recipient string
	Status    int32
}

func (q *Queries) CreateMessage(ctx context.Context, arg CreateMessageParams) (Message, error) {
	row := q.db.QueryRowContext(ctx, createMessage, arg.Content, arg.Recipient, arg.Status)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.Content,
		&i.Recipient,
		&i.Status,
	)
	return i, err
}

const getMessageByID = `-- name: GetMessageByID :one
SELECT id, content, recipient, status FROM messages
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetMessageByID(ctx context.Context, id int32) (Message, error) {
	row := q.db.QueryRowContext(ctx, getMessageByID, id)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.Content,
		&i.Recipient,
		&i.Status,
	)
	return i, err
}

const listPendingMessages = `-- name: ListPendingMessages :many
SELECT id, content, recipient, status FROM messages
WHERE status = 0
ORDER BY id
LIMIT $1
`

func (q *Queries) ListPendingMessages(ctx context.Context, limit int32) ([]Message, error) {
	rows, err := q.db.QueryContext(ctx, listPendingMessages, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Message
	for rows.Next() {
		var i Message
		if err := rows.Scan(
			&i.ID,
			&i.Content,
			&i.Recipient,
			&i.Status,
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

const listProcessingMessages = `-- name: ListProcessingMessages :many
SELECT id, content, recipient, status FROM messages
WHERE status = 1
ORDER BY id DESC
LIMIT $1
`

func (q *Queries) ListProcessingMessages(ctx context.Context, limit int32) ([]Message, error) {
	rows, err := q.db.QueryContext(ctx, listProcessingMessages, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Message
	for rows.Next() {
		var i Message
		if err := rows.Scan(
			&i.ID,
			&i.Content,
			&i.Recipient,
			&i.Status,
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

const listSentMessages = `-- name: ListSentMessages :many
SELECT id, content, recipient, status FROM messages
WHERE status = 2
ORDER BY id DESC
LIMIT $1
`

func (q *Queries) ListSentMessages(ctx context.Context, limit int32) ([]Message, error) {
	rows, err := q.db.QueryContext(ctx, listSentMessages, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Message
	for rows.Next() {
		var i Message
		if err := rows.Scan(
			&i.ID,
			&i.Content,
			&i.Recipient,
			&i.Status,
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

const updateMessageStatus = `-- name: UpdateMessageStatus :exec
UPDATE messages
SET status = $2
WHERE id = $1
`

type UpdateMessageStatusParams struct {
	ID     int32
	Status int32
}

func (q *Queries) UpdateMessageStatus(ctx context.Context, arg UpdateMessageStatusParams) error {
	_, err := q.db.ExecContext(ctx, updateMessageStatus, arg.ID, arg.Status)
	return err
}