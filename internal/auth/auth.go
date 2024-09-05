package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/golang-jwt/jwt/v4"
)

// Claims — структура утверждений, которая включает стандартные утверждения
// и одно пользовательское — UserID
type Claims struct {
	jwt.RegisteredClaims
	UserID int
}

// BuildJWTString создаёт токен и возвращает его в виде строки.
func BuildJWTString() (string, error) {
	// создаём новый токен с алгоритмом подписи HS256 и утверждениями — Claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			// когда создан токен
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(config.TOKENEXP))),
		},
		// собственное утверждение
		UserID: 0,
	})

	// создаём строку токена
	tokenString, err := token.SignedString([]byte(config.SECRETKEY))
	if err != nil {
		return "", err
	}

	// возвращаем строку токена
	return tokenString, nil
}

func GetUserID(tokenString string) (int, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(config.SECRETKEY), nil
		})
	if err != nil {
		return -1, err
	}

	if !token.Valid {
		fmt.Println("token is not valid")
		str := "token is not valid"
		err = errors.New(str)
		return -1, err
	}

	fmt.Println("Token os valid")
	return claims.UserID, nil
}
