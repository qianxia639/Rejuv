// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: friendship.sql

package db

import (
	"context"
)

const createFriendship = `-- name: CreateFriendship :one
INSERT INTO friendships (
    user_id, friend_id, comment
) VALUES (
    $1, $2, $3
)
RETURNING user_id, friend_id, comment, created_at
`

type CreateFriendshipParams struct {
	UserID   int32  `json:"user_id"`
	FriendID int32  `json:"friend_id"`
	Comment  string `json:"comment"`
}

func (q *Queries) CreateFriendship(ctx context.Context, arg *CreateFriendshipParams) (Friendship, error) {
	row := q.db.QueryRow(ctx, createFriendship, arg.UserID, arg.FriendID, arg.Comment)
	var i Friendship
	err := row.Scan(
		&i.UserID,
		&i.FriendID,
		&i.Comment,
		&i.CreatedAt,
	)
	return i, err
}

const deleteFriend = `-- name: DeleteFriend :exec
DELETE FROM friendships 
WHERE (user_id = $1 AND friend_id = $2) 
    OR (user_id = $2 AND friend_id = $1)
`

type DeleteFriendParams struct {
	UserID   int32 `json:"user_id"`
	FriendID int32 `json:"friend_id"`
}

func (q *Queries) DeleteFriend(ctx context.Context, arg *DeleteFriendParams) error {
	_, err := q.db.Exec(ctx, deleteFriend, arg.UserID, arg.FriendID)
	return err
}

const existsFriendship = `-- name: ExistsFriendship :one
SELECT COUNT(*) FROM friend_requests 
WHERE user_id = $1 AND friend_id = $2 AND status = 2
`

type ExistsFriendshipParams struct {
	UserID   int32 `json:"user_id"`
	FriendID int32 `json:"friend_id"`
}

func (q *Queries) ExistsFriendship(ctx context.Context, arg *ExistsFriendshipParams) (int64, error) {
	row := q.db.QueryRow(ctx, existsFriendship, arg.UserID, arg.FriendID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getFriendList = `-- name: GetFriendList :many
SELECT user_id, friend_id, comment, created_at FROM friendships WHERE user_id = $1
`

func (q *Queries) GetFriendList(ctx context.Context, userID int32) ([]Friendship, error) {
	rows, err := q.db.Query(ctx, getFriendList, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Friendship{}
	for rows.Next() {
		var i Friendship
		if err := rows.Scan(
			&i.UserID,
			&i.FriendID,
			&i.Comment,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
