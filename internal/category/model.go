package category

type CategoryDto struct {
	Name string `json:"name" validate:"required"`
	Type string `json:"type" validate:"required"`
}

type Category struct {
	Id   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Type string `json:"type" db:"type"`
}
