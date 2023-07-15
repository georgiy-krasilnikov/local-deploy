package db

import (
	"grpc-service-template/models"

	sq "github.com/Masterminds/squirrel"
)

var (
	idCol       = "id"
	nameCol     = "name"
	lastNameCol = "last_name"
	emailCol    = "email"
	ageCol      = "age"
	userIdCol   = "user_id"
	titleCol    = "title"
	textCol     = "text"
	postIdCol   = "post_id"
)

func buildCreateUserQuery(u models.User) (string, []interface{}, error) {
	return sq.Insert(models.User{}.TableName()).
		Columns(nameCol, lastNameCol, emailCol, ageCol).
		Values(u.Name, u.LastName, u.Email, u.Age).
		Suffix("RETURNING id").PlaceholderFormat(sq.Dollar).ToSql()
}

func buildGetUserByIDQuery(userID int32) (string, []interface{}, error) {
	return sq.Select(idCol, nameCol, lastNameCol, emailCol, ageCol).
		From(models.User{}.TableName()).Where(sq.Eq{idCol: userID}).
		PlaceholderFormat(sq.Dollar).ToSql()
}

func buildGetListOfUsersByIDsQuery(usersIDs []int32) (string, []interface{}, error) {
	return sq.Select(idCol, nameCol, lastNameCol, emailCol, ageCol).
		From(models.User{}.TableName()).Where(sq.Eq{idCol: usersIDs}).
		PlaceholderFormat(sq.Dollar).ToSql()
}

func buildCreatePostQuery(p models.Post) (string, []interface{}, error) {
	return sq.Insert(models.Post{}.TableName()).
		Columns(userIdCol, titleCol, textCol).
		Values(p.UserID, p.Title, p.Text).
		Suffix("RETURNING id").PlaceholderFormat(sq.Dollar).ToSql()
}

func buildJoinTablesQuery(p models.Post) (string, []interface{}, error) {
	return sq.Select("posts.id", userIdCol, titleCol, textCol, nameCol, lastNameCol).
		From(models.Post{}.TableName()).Join("users ON posts.user_id = users.id").
		Where("posts.id = ?", p.ID).PlaceholderFormat(sq.Dollar).ToSql()
}

func buildGetPostsOfUserQuery(userID int32) (string, []interface{}, error) {
	return sq.Select(idCol, userIdCol, titleCol, textCol).
		From(models.Post{}.TableName()).Where(sq.Eq{userIdCol: userID}).
		PlaceholderFormat(sq.Dollar).ToSql()
}

func buildCreateCommentForPostQuery(c models.Comment) (string, []interface{}, error) {
	return sq.Insert(models.Comment{}.TableName()).
		Columns(postIdCol, userIdCol, textCol).
		Values(c.PostID, c.UserID, c.Text).
		Suffix("RETURNING id, post_id").PlaceholderFormat(sq.Dollar).ToSql()
}

func buildJoinPostsAndUsersTablesQuery(postID int32) (string, []interface{}, error) {
	return sq.Select("posts.id", titleCol, textCol, nameCol, lastNameCol).
		From(models.Post{}.TableName()).Join("users ON posts.user_id = users.id").
		Where("posts.id = ?", postID).PlaceholderFormat(sq.Dollar).ToSql()
}

func buildCreateTableWithCommentQuery(c models.Comment) (string, []interface{}, error) {
	return sq.Select("comments.id", textCol, nameCol, lastNameCol).From(models.Comment{}.TableName()).
		Join("users ON comments.user_id = users.id").
		Where(sq.Eq{userIdCol: c.UserID, postIdCol: c.PostID}).PlaceholderFormat(sq.Dollar).ToSql()
}

func buildDeleteCommentFromPostQuery(commentID int32) (string, []interface{}, error) {
	return sq.Delete(models.Comment{}.TableName()).Where(sq.Eq{idCol: commentID}).
		PlaceholderFormat(sq.Dollar).ToSql()
}

func buildGetPostWithCommentsQuery(postID int32) (string, []interface{}, error) {
	return sq.Select("*").From(models.Post{}.TableName()).
		Join("comments ON posts.id = comments.post_id").
		Where(sq.Eq{"posts.id": postID}).PlaceholderFormat(sq.Dollar).ToSql()
}
