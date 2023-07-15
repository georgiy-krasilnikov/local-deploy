package models

import "time"

const usersTableName = "users"

type User struct {
	ID        int32
	Name      string
	LastName  string
	Email     string
	Age       int32
	CreatedAt time.Time
}

func (User) TableName() string {
	return usersTableName
}
