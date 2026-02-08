package auth

type AuthDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Auth struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
