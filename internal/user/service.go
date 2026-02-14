package user

import (
	"strings"

	errortype "github.com/hmdnu/fintr/pkg/errorType"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	db *sqlx.DB
}

func NewService(db *sqlx.DB) *Service {
	return &Service{db: db}
}

func (s *Service) Create(newUser CreateUserDto) error {
	query := "INSERT INTO accounts (name, username, password, balance) VALUES (:name, :username, :password, :balance)"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	tx := s.db.MustBegin()
	_, err = tx.NamedExec(query, &CreateUserDto{Name: newUser.Name, Username: newUser.Username, Password: string(hashedPassword), Balance: newUser.Balance})
	if err != nil {
		if strings.Contains(err.Error(), "2067") {
			return errortype.ConstraintErr(err, "username")
		}
		return err
	}
	tx.Commit()
	return nil
}

func (s *Service) List() ([]User, error) {
	query := "SELECT id, name, username, balance, is_active FROM accounts WHERE is_active = true;"
	users := []User{}
	err := s.db.Select(&users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *Service) Get(id int) (*User, error) {
	query := "SELECT id, name, username, balance FROM accoutns WHERE id = :id AND is_active = true"
	user := User{}
	err := s.db.Get(&user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
