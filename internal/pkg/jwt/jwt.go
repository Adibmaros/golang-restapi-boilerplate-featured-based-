package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type Claims struct {
	UserID uint
	Email  string
	Role   Role
	jwt.RegisteredClaims
}

type Service interface {
	GenerateToken(userID uint, email string, role Role) (string, error)
	ValidateToken(tokenString string) (*Claims, error)
}

type service struct {
	jwtToken []byte
}

func NewJWTService(jwtToken []byte) *service {
	return &service{
		jwtToken: jwtToken,
	}
}

func (s *service) GenerateToken(userID uint, email string, role Role) (string, error) {
	claims := Claims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(s.jwtToken)
}

func (s *service) ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("ada yang salah")
		}
		return s.jwtToken, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("ada yang salah")
	}

	return claims, nil
}
