package auth

import (
	errortype "github.com/hmdnu/fintr/pkg/errorType"
	"github.com/hmdnu/fintr/pkg/token"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	db *sqlx.DB
}

func NewService(db *sqlx.DB) *Service {
	return &Service{db: db}
}

func (s *Service) Login(authDto AuthDto) (string, error) {
	user := Auth{}
	schema := "SELECT id, username, password FROM accounts WHERE username = $1"
	err := s.db.Get(&user, schema, authDto.Username)
	if err != nil {
		return "", errortype.CredInvalidErr
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authDto.Password))
	if err != nil {
		return "", errortype.CredInvalidErr
	}
	token, err := token.GenerateToken(user.Id)
	if err != nil {
		return "", err
	}
	return token, nil
}
