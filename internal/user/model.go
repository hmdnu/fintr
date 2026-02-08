package user

type User struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Username string `json:"username" db:"username"`
	Balance  int    `json:"balance" db:"balance"`
	IsActive bool   `json:"is_active" db:"is_active"`
}

type CreateUserDto struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Balance  int    `json:"balance" validate:"required"`
}
