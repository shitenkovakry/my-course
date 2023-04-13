package models

type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Friends   []int  `json:"friends"`
	CreatedAt int64
}

type Users []*User
