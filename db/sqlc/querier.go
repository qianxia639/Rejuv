// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"context"
)

type Querier interface {
	CreateFriendRequest(ctx context.Context, arg *CreateFriendRequestParams) error
	CreateFriendship(ctx context.Context, arg *CreateFriendshipParams) (Friendship, error)
	CreateGroup(ctx context.Context, arg *CreateGroupParams) (Group, error)
	CreateGroupMember(ctx context.Context, arg *CreateGroupMemberParams) (GroupMember, error)
	CreateUser(ctx context.Context, arg *CreateUserParams) (User, error)
	DeleteFriend(ctx context.Context, arg *DeleteFriendParams) error
	ExistsEmail(ctx context.Context, email string) (int64, error)
	ExistsFriendship(ctx context.Context, arg *ExistsFriendshipParams) (bool, error)
	ExistsGroupMember(ctx context.Context, arg *ExistsGroupMemberParams) (bool, error)
	ExistsNickname(ctx context.Context, nickname string) (int64, error)
	ExistsUsername(ctx context.Context, username string) (int64, error)
	GetFriendList(ctx context.Context, userID int32) ([]Friendship, error)
	GetFriendRequest(ctx context.Context, arg *GetFriendRequestParams) (FriendRequest, error)
	GetGroup(ctx context.Context, groupName string) (Group, error)
	GetGroupMemberList(ctx context.Context, groupID int32) ([]GroupMember, error)
	GetUser(ctx context.Context, username string) (User, error)
	GetUserById(ctx context.Context, id int32) (User, error)
	// 用于更新已过期的申请记录
	UpdateExpiredFriendRequest(ctx context.Context, arg *UpdateExpiredFriendRequestParams) ([]int32, error)
	UpdateFriendRequest(ctx context.Context, arg *UpdateFriendRequestParams) error
	UpdateUser(ctx context.Context, arg *UpdateUserParams) error
}

var _ Querier = (*Queries)(nil)
