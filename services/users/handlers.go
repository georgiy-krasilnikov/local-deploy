package users

import (
	"context"

	"grpc-service-template/models"
	"grpc-service-template/pb"

	log "github.com/sirupsen/logrus"
)

// Условимся начинать интерфейсы с приставки I(Interface)
type IDatabase interface {
	CreateUser(ctx context.Context, newUser models.User) (int32, error)
	GetUserByID(ctx context.Context, userID int32) (*models.User, error)
	GetListOfUsersByIDs(ctx context.Context, usersIDs []int32) ([]models.User, error)
	CreatePost(ctx context.Context, newPost models.Post) (*models.PostTable, error)
	GetPostsOfUser(ctx context.Context, userID int32) ([]models.Post, error)
	CreateCommentForPost(ctx context.Context, newComment models.Comment) (*models.PostTableWithComment, error)
	DeleteCommentFromPost(ctx context.Context, commentID int32) (string, error)
	GetPostWithComments(ctx context.Context, postID int32) (*models.Post, error)
}

type Handler struct {
	pb.UsersServer

	db IDatabase
}

func New(db IDatabase) *Handler {
	return &Handler{db: db}
}

func (h *Handler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	log.WithFields(log.Fields{
		"name":      req.Name,
		"last_name": req.LastName,
		"email":     req.Age,
		"age":       req.Age,
	}).Debug("create new user req")

	userID, err := h.db.CreateUser(ctx, models.User{
		Name:     req.Name,
		LastName: req.LastName,
		Email:    req.Email,
		Age:      req.Age,
	})
	if err != nil {
		log.WithError(err).Error("failed to create new user")
		return nil, err
	}

	return &pb.CreateUserResponse{
		Id: userID,
	}, nil
}

func (h *Handler) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {

	log.WithFields(log.Fields{
		"id": req.Id,
	}).Debug("get user by id req")

	user, err := h.db.GetUserByID(ctx, req.Id)
	if err != nil {
		log.WithError(err).Error("failed to get user by id")
		return nil, err
	}

	return &pb.GetUserByIdResponse{
		User: convertUserToPb(user),
	}, nil
}

func (h *Handler) GetListOfUsersByIds(ctx context.Context, req *pb.GetListOfUsersByIdsRequest) (*pb.GetListOfUsersByIdsResponse, error) {
	log.WithFields(log.Fields{
		"id": req.Id,
	}).Debug("get users by ids req")

	users, err := h.db.GetListOfUsersByIDs(ctx, req.Id)
	if err != nil {
		log.WithError(err).Error("failed to get users by ids")
		return nil, err
	}

	return &pb.GetListOfUsersByIdsResponse{
		Users: convertUsersToPb(users),
	}, nil
}

func (h *Handler) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	log.WithFields(log.Fields{
		"user_id": req.UserId,
		"title":   req.Title,
		"text":    req.Text,
	}).Debug("create new post req")

	postTable, err := h.db.CreatePost(ctx, models.Post{
		UserID: req.UserId,
		Title:  req.Title,
		Text:   req.Text,
	})
	if err != nil {
		log.WithError(err).Error("failed to create new post")
		return nil, err
	}

	return &pb.CreatePostResponse{
		PostTable: convertPostTableToPb(postTable),
	}, nil
}

func (h *Handler) GetPostsOfUser(ctx context.Context, req *pb.GetPostsOfUserRequest) (*pb.GetPostsOfUserResponse, error) {
	log.WithFields(log.Fields{
		"id": req.UserId,
	}).Debug("get posts of user req")

	posts, err := h.db.GetPostsOfUser(ctx, req.UserId)
	if err != nil {
		log.WithError(err).Error("failed to get posts of user")
		return nil, err
	}

	return &pb.GetPostsOfUserResponse{
		Posts: convertPostsToPb(posts),
	}, nil
}

func (h *Handler) CreateCommentForPost(ctx context.Context, req *pb.CreateCommentForPostRequest) (*pb.CreateCommentForPostResponse, error) {
	log.WithFields(log.Fields{
		"post_id": req.PostId,
		"user_id": req.UserId,
		"text":    req.Text,
	}).Debug("create comment for post req")

	postTable, err := h.db.CreateCommentForPost(ctx, models.Comment{
		PostID: req.PostId,
		UserID: req.UserId,
		Text:   req.Text,
	})
	if err != nil {
		log.WithError(err).Error("failed create comment for post")
		return nil, err
	}

	return &pb.CreateCommentForPostResponse{
		PostTableWithComment: convertCommentPostTableToPb(postTable),
	}, nil
}

func (h *Handler) DeleteCommentFromPost(ctx context.Context, req *pb.DeleteCommentFromPostRequest) (*pb.DeleteCommentFromPostResponse, error) {
	log.WithFields(log.Fields{
		"comment_id": req.CommentId,
	}).Debug("delete comment from post req")

	message, err := h.db.DeleteCommentFromPost(ctx, req.CommentId)
	if err != nil {
		log.WithError(err).Error("failed to delete comment from post")
		return nil, err
	}

	return &pb.DeleteCommentFromPostResponse{
		Message: message,
	}, nil
}

func (h *Handler) GetPostWithComments(ctx context.Context, req *pb.GetPostWithCommentsRequest) (*pb.GetPostWithCommentsResponse, error) {
	log.WithFields(log.Fields{
		"post_id": req.PostId,
	}).Debug("get post with comments req")

	post, err := h.db.GetPostWithComments(ctx, req.PostId)
	if err != nil {
		log.WithError(err).Error("failed get post with comments")
		return nil, err
	}

	return &pb.GetPostWithCommentsResponse{
		Post: convertPostToPb(post),
	}, nil
}
