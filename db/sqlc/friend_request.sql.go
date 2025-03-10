// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: friend_request.sql

package db

import (
	"context"
)

const createFriendRequest = `-- name: CreateFriendRequest :exec
INSERT INTO friend_requests (
    user_id, friend_id, request_desc
) VALUES (
    $1, $2, $3
)
`

type CreateFriendRequestParams struct {
	UserID      int32  `json:"user_id"`
	FriendID    int32  `json:"friend_id"`
	RequestDesc string `json:"request_desc"`
}

func (q *Queries) CreateFriendRequest(ctx context.Context, arg *CreateFriendRequestParams) error {
	_, err := q.db.Exec(ctx, createFriendRequest, arg.UserID, arg.FriendID, arg.RequestDesc)
	return err
}

const existsFriendRequest = `-- name: ExistsFriendRequest :one
SELECT COUNT(*) FROM friend_requests 
WHERE 
	((user_id = $1 AND friend_id = $2) OR 
	(user_id = $2 AND friend_id = $1)) AND status = 1
`

type ExistsFriendRequestParams struct {
	UserID   int32 `json:"user_id"`
	FriendID int32 `json:"friend_id"`
}

func (q *Queries) ExistsFriendRequest(ctx context.Context, arg *ExistsFriendRequestParams) (int64, error) {
	row := q.db.QueryRow(ctx, existsFriendRequest, arg.UserID, arg.FriendID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const updateFriendRequest = `-- name: UpdateFriendRequest :exec
UPDATE friend_requests
SET
	status  = $3,
	updated_at = now()
WHERE
user_id = $1 AND friend_id = $2 AND status = 1
`

type UpdateFriendRequestParams struct {
	UserID   int32 `json:"user_id"`
	FriendID int32 `json:"friend_id"`
	Status   int8  `json:"status"`
}

func (q *Queries) UpdateFriendRequest(ctx context.Context, arg *UpdateFriendRequestParams) error {
	_, err := q.db.Exec(ctx, updateFriendRequest, arg.UserID, arg.FriendID, arg.Status)
	return err
}
