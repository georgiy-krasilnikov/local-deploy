package models

const postsTableName = "posts"

type Post struct {
	ID       int32
	UserID   int32
	Title    string
	Text     string
	Comments []Comment
}

type PostTable struct {
	ID           int32
	UserID       int32
	UserName     string
	UserLastName string
	Title        string
	Text         string
}

func (Post) TableName() string {
	return postsTableName
}
