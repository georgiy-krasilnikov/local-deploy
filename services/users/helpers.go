package users

import (
	"grpc-service-template/models"
	"grpc-service-template/pb"
)

func convertUserToPb(u *models.User) *pb.User {
	if u == nil {
		return nil
	}

	return &pb.User{
		Id:       u.ID,
		Name:     u.Name,
		LastName: u.LastName,
		Email:    u.Email,
		Age:      u.Age,
	}
}

func convertUsersToPb(u []models.User) []*pb.User {
	users := make([]*pb.User, len(u))
	for i, user := range u {
		users[i] = &pb.User{
			Id:       user.ID,
			Name:     user.Name,
			LastName: user.LastName,
			Email:    user.Email,
			Age:      user.Age,
		}
	}

	return users
}

func convertPostsToPb(p []models.Post) []*pb.Post {
	posts := make([]*pb.Post, len(p))
	for i, post := range p {
		posts[i] = &pb.Post{
			Id:     post.ID,
			UserId: post.UserID,
			Title:  post.Title,
			Text:   post.Text,
		}
	}

	return posts
}

func convertPostTableToPb(p *models.PostTable) *pb.PostTable {
	if p == nil {
		return nil
	}

	return &pb.PostTable{
		Id:           p.ID,
		UserId:       p.UserID,
		UserName:     p.UserName,
		UserLastName: p.UserLastName,
		Title:        p.Title,
		Text:         p.Text,
	}
}

func convertCommentPostTableToPb(p *models.PostTableWithComment) *pb.PostTableWithComment {
	if p == nil {
		return nil
	}

	return &pb.PostTableWithComment{
		Id:                  p.ID,
		UserName:            p.UserName,
		UserLastName:        p.UserLastName,
		Title:               p.Title,
		Text:                p.Text,
		CommentId:           p.CommentID,
		Comment:             p.CommentText,
		CommentUserName:     p.CommentUserName,
		CommentUserLastName: p.CommentUserLastName,
	}
}

func convertCommentsToPb(c []models.Comment) []*pb.Comment {
	comments := make([]*pb.Comment, len(c))
	for i, comment := range c {
		comments[i] = &pb.Comment{
			Id:     comment.ID,
			PostId: comment.PostID,
			UserId: comment.UserID,
			Text:   comment.Text,
		}
	}

	return comments
}

func convertPostToPb(p *models.Post) *pb.Post {
	if p == nil {
		return nil
	}

	return &pb.Post{
		Id:       p.ID,
		UserId:   p.UserID,
		Title:    p.Title,
		Text:     p.Text,
		Comments: convertCommentsToPb(p.Comments),
	}
}
