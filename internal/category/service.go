package category

import (
	"strings"

	errortype "github.com/hmdnu/fintr/pkg/errorType"
	"github.com/jmoiron/sqlx"
)

type Service struct {
	db *sqlx.DB
}

func NewService(db *sqlx.DB) *Service {
	return &Service{db: db}
}

func (s *Service) Create(categoryDto *CategoryDto) error {
	query := `INSERT INTO categories (name, type) VALUES (:name, :type);`
	tx := s.db.MustBegin()
	_, err := tx.NamedExec(query, &CategoryDto{Name: categoryDto.Name, Type: categoryDto.Type})
	if err != nil {
		if strings.Contains(err.Error(), "2067") {
			return errortype.ConstraintErr(err, "name")
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) List() ([]Category, error) {
	query := `SELECT * FROM categories;`
	categories := []Category{}
	err := s.db.Select(&categories, query)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
