package users

import (
	"testing"
	"time"

	"grpc-service-template/models"
	"grpc-service-template/pb"

	"github.com/stretchr/testify/assert"
)

func Test_convertUserToPb(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		user := models.User{
			ID:        100,
			Name:      "Ivan",
			LastName:  "Ivanov",
			Email:     "ivan@fletn.com",
			Age:       100,
			CreatedAt: time.Now(),
		}

		want := &pb.User{
			Id:       user.ID,
			Name:     user.Name,
			LastName: user.LastName,
			Email:    user.Email,
			Age:      user.Age,
		}

		got := convertUserToPb(&user)

		assert.NotNil(t, got)
		assert.Equal(t, want, got)
	})
}

func Test_convertUsersToPb(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		users := make([]models.User, 2)
		for i, _ := range users {
			users[i] = models.User{
				ID:        int32(i + 1),
				Name:      "Ivan",
				LastName:  "Ivanov",
				Email:     "ivan@fletn.com",
				Age:       int32((i + 1)) * 10,
				CreatedAt: time.Now(),
			}
		}

		pbUsers := make([]*pb.User, 2)
		for i, u := range users {
			pbUsers[i] = &pb.User{
				Id:       u.ID,
				Name:     u.Name,
				LastName: u.LastName,
				Email:    u.Email,
				Age:      u.Age,
			}
		}

		want := pbUsers
		got := convertUsersToPb(users)

		assert.NotNil(t, got)
		assert.Equal(t, want, got)
	})
}

func Test_convertPostsToPb(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		posts := make([]models.Post, 2)
		for i, _ := range posts {
			posts[i] = models.Post{
				ID:     int32(i + 1),
				UserID: 2,
				Title:  "Holidays",
				Text:   "Meet friends",
			}
		}

		pbPosts := make([]*pb.Post, 2)
		for i, p := range posts {
			pbPosts[i] = &pb.Post{
				Id:     p.ID,
				UserId: p.UserID,
				Title:  p.Title,
				Text:   p.Text,
			}
		}

		want := pbPosts
		got := convertPostsToPb(posts)

		assert.NotNil(t, got)
		assert.Equal(t, want, got)
	})
}

func Test_convertPostTableToPb(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		p := &models.PostTableWithComment{
			ID:                  int32(1),
			Title:               "Weekends",
			Text:                "watch film",
			UserName:            "Gosha",
			UserLastName:        "Krasilnikov",
			CommentID:           int32(2),
			CommentText:         "how are you",
			CommentUserName:     "Ilya",
			CommentUserLastName: "Krasnov",
		}

		pbPostTable := &pb.PostTableWithComment{
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

		want := pbPostTable
		got := convertCommentPostTableToPb(p)

		assert.NotNil(t, got)
		assert.Equal(t, want, got)
	})
}

func Test_convertPostTable(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		p := &models.PostTable{
			ID:           int32(1),
			UserID:       int32(2),
			Title:        "Weekends",
			Text:         "watch film",
			UserName:     "Gosha",
			UserLastName: "Krasilnikov",
		}

		pbPostTable := &pb.PostTable{
			Id:           p.ID,
			UserId:       p.UserID,
			UserName:     p.UserName,
			UserLastName: p.UserLastName,
			Title:        p.Title,
			Text:         p.Text,
		}

		want := pbPostTable
		got := convertPostTableToPb(p)

		assert.NotNil(t, got)
		assert.Equal(t, want, got)
	})
}

func Test_convertCommentsToPb(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		comments := make([]models.Comment, 2)
		for i, _ := range comments {
			comments[i] = models.Comment{
				ID:     int32(i + 1),
				PostID: int32(i * 2),
				UserID: 2,
				Text:   "hello",
			}
		}

		pbComments := make([]*pb.Comment, 2)
		for i, c := range comments {
			pbComments[i] = &pb.Comment{
				Id:     c.ID,
				PostId: c.PostID,
				UserId: c.UserID,
				Text:   c.Text,
			}
		}

		want := pbComments
		got := convertCommentsToPb(comments)

		assert.NotNil(t, got)
		assert.Equal(t, want, got)
	})
}

func Test_convertPostToPb(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		post := &models.Post{
			ID:     int32(1),
			UserID: int32(2),
			Title:  "holidays",
			Text:   "meet with freinds",
		}

		var comments [2]models.Comment
		for i, _ := range comments {
			comments[i] = models.Comment{
				ID:     int32(i + 1),
				PostID: int32(1),
				UserID: 2,
				Text:   "hello",
			}
			post.Comments = append(post.Comments, comments[i])
		}

		pbPost := &pb.Post{
			Id:       post.ID,
			UserId:   post.UserID,
			Title:    post.Title,
			Text:     post.Text,
			Comments: convertCommentsToPb(post.Comments),
		}

		want := pbPost
		got := convertPostToPb(post)

		assert.NotNil(t, got)
		assert.Equal(t, want, got)
	})
}
