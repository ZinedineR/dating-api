package jwthelper

import (
	"dating-api/internal/profile/domain"
	"dating-api/pkg/errs"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Age         int       `json:"age"`
	Sex         string    `json:"sex"`
	Orientation string    `json:"orientation"`
	Status      string    `json:"status"`
	Account     string    `json:"account"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	LastLogin   time.Time `json:"last_login"`
	jwt.RegisteredClaims
}

func GenerateJWT(data domain.Profile) (tokenString string, err errs.Error) {
	// expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Id:          data.Id,
		Name:        data.Name,
		Age:         data.Age,
		Sex:         data.Sex,
		Orientation: data.Orientation,
		Status:      data.Status,
		Account:     data.Email,
		Email:       data.Email,
		Password:    data.Password,
		LastLogin:   time.Now(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	completedString, error := token.SignedString(jwtKey)
	if error != nil {
		return "", errs.Wrap(err)
	}
	// tokenString, err = token.SignedString(jwtKey)
	return completedString, nil
}
func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt.Before(time.Now()) {
		err = errors.New("token expired")
		return
	}
	return
}
