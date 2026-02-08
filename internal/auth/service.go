package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
	token, err := generateToken(user, "secret")
	if err != nil {
		return "", err
	}
	return token, nil
}

func generateToken(user Auth, secret string) (string, error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   strconv.Itoa(user.Id),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
