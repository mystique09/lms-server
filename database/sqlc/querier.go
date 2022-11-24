// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package database

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	AddNewClassroomMember(ctx context.Context, arg AddNewClassroomMemberParams) (ClassroomMember, error)
	CreateClass(ctx context.Context, arg CreateClassParams) (Classroom, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (string, error)
	DeleteClass(ctx context.Context, id uuid.UUID) (Classroom, error)
	DeleteClassworkFromClass(ctx context.Context, arg DeleteClassworkFromClassParams) (ClassWork, error)
	DeleteCommentFromPost(ctx context.Context, arg DeleteCommentFromPostParams) (interface{}, error)
	DeletePostFromClass(ctx context.Context, arg DeletePostFromClassParams) (interface{}, error)
	DeleteUser(ctx context.Context, id uuid.UUID) (uuid.UUID, error)
	FollowUser(ctx context.Context, arg FollowUserParams) (UserFollow, error)
	GetAllClassFromUser(ctx context.Context, arg GetAllClassFromUserParams) ([]Classroom, error)
	GetAllClassroomMembers(ctx context.Context, arg GetAllClassroomMembersParams) ([]GetAllClassroomMembersRow, error)
	GetAllCommentLikes(ctx context.Context, commentID uuid.UUID) ([]CommentLike, error)
	GetAllCommentsFromPost(ctx context.Context, arg GetAllCommentsFromPostParams) ([]Comment, error)
	GetAllFollowers(ctx context.Context, arg GetAllFollowersParams) ([]GetAllFollowersRow, error)
	GetAllFollowing(ctx context.Context, arg GetAllFollowingParams) ([]GetAllFollowingRow, error)
	GetAllJoinedClassrooms(ctx context.Context, arg GetAllJoinedClassroomsParams) ([]Classroom, error)
	GetAllPostLikes(ctx context.Context, postID uuid.UUID) ([]PostLike, error)
	GetClass(ctx context.Context, id uuid.UUID) (Classroom, error)
	GetClassWork(ctx context.Context, arg GetClassWorkParams) (ClassWork, error)
	GetClassroomMemberById(ctx context.Context, arg GetClassroomMemberByIdParams) (ClassroomMember, error)
	GetClassroomWithInviteCode(ctx context.Context, inviteCode uuid.UUID) (uuid.UUID, error)
	GetFollowerById(ctx context.Context, id uuid.UUID) (GetFollowerByIdRow, error)
	GetOneFollower(ctx context.Context, arg GetOneFollowerParams) (UserFollow, error)
	GetOnePost(ctx context.Context, id uuid.UUID) (Post, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetUser(ctx context.Context, id uuid.UUID) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	GetUsers(ctx context.Context, offset int32) ([]User, error)
	InsertNewClasswork(ctx context.Context, arg InsertNewClassworkParams) (ClassWork, error)
	InsertNewCommentInPost(ctx context.Context, arg InsertNewCommentInPostParams) (Comment, error)
	InsertNewPost(ctx context.Context, arg InsertNewPostParams) (Post, error)
	LeaveClassroom(ctx context.Context, arg LeaveClassroomParams) (ClassroomMember, error)
	LikeComment(ctx context.Context, arg LikeCommentParams) (CommentLike, error)
	LikePost(ctx context.Context, arg LikePostParams) (PostLike, error)
	ListAllPostsByUser(ctx context.Context, arg ListAllPostsByUserParams) ([]Post, error)
	ListAllPostsFromClass(ctx context.Context, arg ListAllPostsFromClassParams) ([]Post, error)
	ListAllPublicClass(ctx context.Context, offset int32) ([]Classroom, error)
	ListClassworkAdmin(ctx context.Context, arg ListClassworkAdminParams) ([]ClassWork, error)
	ListSubmittedClassworks(ctx context.Context, arg ListSubmittedClassworksParams) ([]ClassWork, error)
	UnfollowUser(ctx context.Context, id uuid.UUID) (UserFollow, error)
	UnlikeComment(ctx context.Context, arg UnlikeCommentParams) (CommentLike, error)
	UnlikePost(ctx context.Context, arg UnlikePostParams) (PostLike, error)
	UpdateAClassworkMark(ctx context.Context, arg UpdateAClassworkMarkParams) error
	UpdateClassroomDescription(ctx context.Context, arg UpdateClassroomDescriptionParams) (Classroom, error)
	UpdateClassroomInviteCode(ctx context.Context, arg UpdateClassroomInviteCodeParams) (Classroom, error)
	UpdateClassroomName(ctx context.Context, arg UpdateClassroomNameParams) (Classroom, error)
	UpdateClassroomRoom(ctx context.Context, arg UpdateClassroomRoomParams) (Classroom, error)
	UpdateClassroomSection(ctx context.Context, arg UpdateClassroomSectionParams) (Classroom, error)
	UpdateClassroomSubject(ctx context.Context, arg UpdateClassroomSubjectParams) (Classroom, error)
	UpdateCommentContentInPost(ctx context.Context, arg UpdateCommentContentInPostParams) (interface{}, error)
	UpdatePostContent(ctx context.Context, arg UpdatePostContentParams) (interface{}, error)
	UpdateUserEmail(ctx context.Context, arg UpdateUserEmailParams) (uuid.UUID, error)
	UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) (uuid.UUID, error)
	UpdateUsername(ctx context.Context, arg UpdateUsernameParams) (uuid.UUID, error)
}

var _ Querier = (*Queries)(nil)
