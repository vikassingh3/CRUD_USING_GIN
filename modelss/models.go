package modelss

type User struct {
	ID       int    `json:"id"`
	Names    string `json:"names"`
	Username string `json:"username"`
	Email    string ` json:"email"`
	Password string ` json:"password"`
}
