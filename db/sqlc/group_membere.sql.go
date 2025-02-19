// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: group_membere.sql

package db

import (
	"context"
)

const createGroupMember = `-- name: CreateGroupMember :one
INSERT INTO group_members (
	group_id, user_id, role, agreed
) VALUES (
	$1, $2, $3, $4
)
RETURNING group_id, user_id, role, agreed, joined_at
`

type CreateGroupMemberParams struct {
	GroupID int32 `json:"group_id"`
	UserID  int32 `json:"user_id"`
	Role    int16 `json:"role"`
	Agreed  bool  `json:"agreed"`
}

func (q *Queries) CreateGroupMember(ctx context.Context, arg *CreateGroupMemberParams) (GroupMember, error) {
	row := q.db.QueryRow(ctx, createGroupMember,
		arg.GroupID,
		arg.UserID,
		arg.Role,
		arg.Agreed,
	)
	var i GroupMember
	err := row.Scan(
		&i.GroupID,
		&i.UserID,
		&i.Role,
		&i.Agreed,
		&i.JoinedAt,
	)
	return i, err
}
