package db

import (
	"context"

	"grpc-service-template/models"
)

func (db *DB) CreateCommentForPost(ctx context.Context, newComment models.Comment) (*models.PostTableWithComment, error) {
	query, args, err := buildCreateCommentForPostQuery(newComment)
	if err != nil {
		return nil, err
	}

	if err = db.QueryRow(ctx, query, args...).Scan(&newComment.ID, &newComment.PostID); err != nil {
		return nil, err
	}

	var postTable models.PostTableWithComment

	query, args, err = buildJoinPostsAndUsersTablesQuery(newComment.PostID)
	if err != nil {
		return nil, err
	}
	if err = db.QueryRow(ctx, query, args...).Scan(
		&postTable.ID,
		&postTable.Title,
		&postTable.Text,
		&postTable.UserName,
		&postTable.UserLastName,
	); err != nil {
		return nil, err
	}

	query, args, err = buildCreateTableWithCommentQuery(newComment)
	if err != nil {
		return nil, err
	}
	if err = db.QueryRow(ctx, query, args...).Scan(
		&postTable.CommentID,
		&postTable.CommentText,
		&postTable.CommentUserName,
		&postTable.CommentUserLastName,
	); err != nil {
		return nil, err
	}

	return &postTable, nil
}

func (db *DB) DeleteCommentFromPost(ctx context.Context, commentID int32) (string, error) {
	query, args, err := buildDeleteCommentFromPostQuery(commentID)
	if err != nil {
		return "", err
	}

	var message string
	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	return message, nil
}

func (db *DB) GetPostWithComments(ctx context.Context, postID int32) (*models.Post, error) {
	query, args, err := buildGetPostWithCommentsQuery(postID)
	if err != nil {
		return nil, err
	}

	var (
		post models.Post
		c    models.Comment
	)

	rows, err := db.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(
			&post.ID,
			&post.UserID,
			&post.Title,
			&post.Text,
			&c.ID,
			&c.PostID,
			&c.UserID,
			&c.Text,
		); err != nil {
			return nil, err
		}
		post.Comments = append(post.Comments, c)
	}

	return &post, nil
}
