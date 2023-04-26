package models

import "errors"

var (
	ErrNoDocuments = errors.New("no documents")
)

type User struct {
	ID        int    `bson:"id"`
	Email     string `bson:"email"`
	Name      string `bson:"name"`
	Age       int    `bson:"age"`
	Friends   []int  `bson:"friends"`
	CreatedAt int64  `bson:"created_at"`
}

type Users []*User
