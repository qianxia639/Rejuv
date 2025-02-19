// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: group.sql

package db

import (
	"context"
)

const createGroup = `-- name: CreateGroup :one
INSERT INTO groups (
	group_name, creator_id, description
) VALUES (
	$1, $2, $3
) RETURNING id, group_name, creator_id, group_avatar_url, description, max_member, created_at, updated_at
`

type CreateGroupParams struct {
	GroupName   string `json:"group_name"`
	CreatorID   int32  `json:"creator_id"`
	Description string `json:"description"`
}

func (q *Queries) CreateGroup(ctx context.Context, arg *CreateGroupParams) (Group, error) {
	row := q.db.QueryRow(ctx, createGroup, arg.GroupName, arg.CreatorID, arg.Description)
	var i Group
	err := row.Scan(
		&i.ID,
		&i.GroupName,
		&i.CreatorID,
		&i.GroupAvatarUrl,
		&i.Description,
		&i.MaxMember,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getGroup = `-- name: GetGroup :one
SELECT id, group_name, creator_id, group_avatar_url, description, max_member, created_at, updated_at FROM groups WHERE group_name = $1 LIMIT 1
`

func (q *Queries) GetGroup(ctx context.Context, groupName string) (Group, error) {
	row := q.db.QueryRow(ctx, getGroup, groupName)
	var i Group
	err := row.Scan(
		&i.ID,
		&i.GroupName,
		&i.CreatorID,
		&i.GroupAvatarUrl,
		&i.Description,
		&i.MaxMember,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
