package deleteuser

type DeleteUser struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Friends []int  `json:"friends"`
}
