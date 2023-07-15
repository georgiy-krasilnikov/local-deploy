package db

import (
	"context"

	"grpc-service-template/models"
)

func (db *DB) CreateUser(ctx context.Context, newUser models.User) (int32, error) {
	query, args, err := buildCreateUserQuery(newUser)
	if err != nil {
		return 0, err
	}

	var newUserID int32
	err = db.QueryRow(ctx, query, args...).Scan(&newUserID)
	if err != nil {
		return 0, err
	}

	return newUserID, nil
}

func (db *DB) GetUserByID(ctx context.Context, userID int32) (*models.User, error) {
	query, args, err := buildGetUserByIDQuery(userID)
	if err != nil {
		return nil, err
	}

	var user models.User
	if err = db.QueryRow(ctx, query, args...).Scan(
		&user.ID,
		&user.Name,
		&user.LastName,
		&user.Email,
		&user.Age,
	); err != nil {
		return nil, err
	}

	return &user, nil
}

func (db *DB) GetListOfUsersByIDs(ctx context.Context, usersIDs []int32) ([]models.User, error) {
	query, args, err := buildGetListOfUsersByIDsQuery(usersIDs)

	if err != nil {
		return nil, err
	}
	var users []models.User
	rows, err := db.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	var u models.User
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&u.ID, &u.Name, &u.LastName, &u.Email, &u.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}