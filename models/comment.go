package models

const commentsTableName = "comments"

type Comment struct {
	ID     int32
	PostID int32
	UserID int32
	Text   string
}

type PostTableWithComment struct {
	ID                  int32
	UserName            string
	UserLastName        string
	Title               string
	Text                string
	CommentID           int32
	CommentUserName     string
	CommentUserLastName string
	CommentText         string
}

func (Comment) TableName() string {
	return commentsTableName
}
