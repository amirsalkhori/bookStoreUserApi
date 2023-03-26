package users

type User struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Family    string `json:"family"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
}
