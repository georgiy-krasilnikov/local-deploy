package db

import (
	"testing"
	"time"

	"grpc-service-template/models"

	"github.com/stretchr/testify/assert"
)

func Test_buildCreateUserQuery(t *testing.T) {
	t.Run("SuccessWithFullUserData", func(t *testing.T) {
		user := models.User{
			ID:        100,
			Name:      "Ivan",
			LastName:  "Ivanov",
			Email:     "ivan@fletn.com",
			Age:       50,
			CreatedAt: time.Now(),
		}

		want := "INSERT INTO users (name,last_name,email,age) VALUES ($1,$2,$3,$4) RETURNING id"
		got, args, err := buildCreateUserQuery(user)

		assert.NoError(t, err)
		assert.Equal(t, want, got)
		assert.Equal(t, 4, len(args))
		assert.Equal(t, "Ivan", args[0].(string))
	})

	t.Run("SuccessWithNonFullUserData", func(t *testing.T) {
		user := models.User{
			Name:     "Petr",
			LastName: "Petrov",
			Email:    "petr@fletn.com",
			Age:      40,
		}

		want := "INSERT INTO users (name,last_name,email,age) VALUES ($1,$2,$3,$4) RETURNING id"
		got, args, err := buildCreateUserQuery(user)

		assert.NoError(t, err)
		assert.Equal(t, want, got)
		assert.Equal(t, 4, len(args))
		assert.Equal(t, "petr@fletn.com", args[2].(string))
	})
}

func Test_buildGetUserByIDQuery(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		id := int32(123)

		want := "SELECT id, name, last_name, email, age FROM users WHERE id = $1"
		got, args, err := buildGetUserByIDQuery(id)

		assert.NoError(t, err)
		assert.Equal(t, want, got)
		assert.Equal(t, 1, len(args))
		assert.Equal(t, id, args[0])
	})
}

func Test_buildGetListOfUsersByIDsQuery(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ids := []int32{1, 2}

		want := "SELECT id, name, last_name, email, age FROM users WHERE id IN ($1,$2)"
		got, args, err := buildGetListOfUsersByIDsQuery(ids)

		assert.NoError(t, err)
		assert.Equal(t, want, got)
		assert.Equal(t, 2, len(args))
		assert.Equal(t, ids[0], args[0])
		assert.Equal(t, ids[1], args[1])
	})
}

func Test_buildCreatePostQuery(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		post := models.Post{
			ID:     5,
			UserID: 7,
			Title:  "Holidays",
			Text:   "Go to cinema",
		}

		want := "INSERT INTO posts (user_id,title,text) VALUES ($1,$2,$3) RETURNING (SELECT (name, last_name) FROM users JOIN posts ON posts.user_id = users.id)"
		got, args, err := buildCreatePostQuery(post)

		assert.NoError(t, err)
		assert.Equal(t, want, got)
		assert.Equal(t, 5, len(args))
		assert.Equal(t, int32(7), args[0])
	})
}

func Test_buildJoinTablesQuery(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		post := models.Post{
			ID:     5,
			UserID: 7,
			Title:  "Holidays",
			Text:   "Go to cinema",
		}

		want := "SELECT posts.id, user_id, title, text, name, last_name FROM posts JOIN users ON posts.user_id = users.id WHERE posts.id = $1"
		got, args, err := buildJoinTablesQuery(post)

		assert.NoError(t, err)
		assert.Equal(t, want, got)
		assert.Equal(t, 1, len(args))
		assert.Equal(t, int32(5), args[0])
	})
}

func Test_buildGetPostsOfUserQuery(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var posts [2]models.Post
		for i, _ := range posts {
			posts[i] = models.Post{
				ID:     int32(i + 6),
				UserID: 7,
				Title:  "Holidays",
				Text:   "Go to cinema",
			}
		}

		want := "SELECT id, user_id, title, text FROM posts WHERE user_id = $1"
		got, args, err := buildGetPostsOfUserQuery(7)

		assert.NoError(t, err)
		assert.Equal(t, want, got)
		assert.Equal(t, 1, len(args))
	})
}

func Test_buildJoinPostsAndUsersTablesQuery(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		post := models.Post{
			ID:     5,
			UserID: 7,
			Title:  "Holidays",
			Text:   "Go to cinema",
		}

		want := "SELECT posts.id, title, text, name, last_name FROM posts JOIN users ON posts.user_id = users.id WHERE posts.id = $1"
		got, args, err := buildJoinPostsAndUsersTablesQuery(post.ID)

		assert.NoError(t, err)
		assert.Equal(t, want, got)
		assert.Equal(t, 1, len(args))
		assert.Equal(t, int32(5), args[0])
	})
}

func Test_buildCreateTableWithCommentQuery(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		comment := models.Comment{
			ID:     5,
			UserID: 7,
			Text:   "which film?",
		}

		want := "SELECT comments.id, text, name, last_name FROM comments JOIN users ON comments.user_id = users.id WHERE post_id = $1 AND user_id = $2"
		got, args, err := buildCreateTableWithCommentQuery(comment)

		assert.NoError(t, err)
		assert.Equal(t, want, got)
		assert.Equal(t, 2, len(args))
		assert.Equal(t, int32(7), args[1])
	})
}

func Test_buildDeleteCommentFromPost(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		commentID := int32(5)

		want := "DELETE FROM comments WHERE id = $1"
		got, args, err := buildDeleteCommentFromPostQuery(commentID)

		assert.NoError(t, err)
		assert.Equal(t, want, got)
		assert.Equal(t, 1, len(args))
		assert.Equal(t, commentID, args[0])
	})
}

func Test_buildGetPostWithCommentsQuery(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		comment := models.Comment{
			ID:     5,
			UserID: 7,
			Text:   "which film?",
		}

		want := "SELECT * FROM posts JOIN comments ON posts.id = comments.post_id WHERE posts.id = $1"
		got, args, err := buildGetPostWithCommentsQuery(comment.ID)

		assert.NoError(t, err)
		assert.Equal(t, want, got)
		assert.Equal(t, 1, len(args))
		assert.Equal(t, comment.ID, args[0])
	})
}
