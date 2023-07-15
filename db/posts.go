package db

import (
	"context"

	"grpc-service-template/models"
)

func (db *DB) CreatePost(ctx context.Context, newPost models.Post) (*models.PostTable, error) {
	query, args, err := buildCreatePostQuery(newPost)
	if err != nil {
		return nil, err
	}

	var postTable models.PostTable
	if err = db.QueryRow(ctx, query, args...).Scan(&newPost.ID); err != nil {
		return nil, err
	}

	query, args, err = buildJoinTablesQuery(newPost)
	if err != nil {
		return nil, err
	}

	if err = db.QueryRow(ctx, query, args...).Scan(
		&postTable.ID,
		&postTable.UserID,
		&postTable.Title,
		&postTable.Text,
		&postTable.UserName,
		&postTable.UserLastName,
	); err != nil {
		return nil, err
	}

	return &postTable, nil
}

func (db *DB) GetPostsOfUser(ctx context.Context, userID int32) ([]models.Post, error) {
	query, args, err := buildGetPostsOfUserQuery(userID)
	if err != nil {
		return nil, err
	}

	var posts []models.Post
	rows, err := db.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	var p models.Post
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&p.ID, &p.UserID, &p.Title, &p.Text)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}

	return posts, nil
}
