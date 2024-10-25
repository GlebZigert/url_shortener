package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/GlebZigert/url_shortener.git/internal/packerr"
	"github.com/golang-jwt/jwt/v4"
)

// Струтура с методами для аутентификации и авторизации
type AuthController struct {
	sekretKey string
	tokenExp  int
	UIDkey    int
}

// ошибка если доступ запрещен
var ErrAccessDenied = errors.New("access denied")

// Claims — структура утверждений, которая включает стандартные утверждения
// и одно пользовательское — UID
type Claims struct {
	jwt.RegisteredClaims
	UID int
}

// Конструктор
func NewAuth(sekretKey string, tokenExp int, key int) *AuthController {
	return &AuthController{sekretKey, tokenExp, key}
}

// ошибка при формировании токена
var ErrBuildJWTString error = errors.New("ошибка формирования JWT")

// BuildJWTString создаёт токен и возвращает его в виде строки.
func (auc *AuthController) BuildJWTString(id int) (string, error) {

	// создаём новый токен с алгоритмом подписи HS256 и утверждениями — Claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			// когда создан токен
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(auc.tokenExp))),
		},
		// собственное утверждение
		UID: id,
	})

	// создаём строку токена
	tokenString, err := token.SignedString([]byte(auc.sekretKey))
	if err != nil {
		return "", packerr.NewTimeError(ErrBuildJWTString)
	}

	// возвращаем строку токена
	return tokenString, nil
}

// Получить uid из токена
func (auc *AuthController) GetUserID(tokenString string) (int, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, packerr.NewTimeError(fmt.Errorf("unexpected signing method: %v", t.Header["alg"]))
			}
			return []byte(auc.sekretKey), nil
		})
	if err != nil {
		return -1, packerr.NewTimeError(err)
	}

	if !token.Valid {

		str := "token is not valid"
		err = errors.New(str)
		return -1, packerr.NewTimeError(err)
	}

	return claims.UID, nil
}

// достать uid из контекста реквеста - применяется в ендпойнтах
func (auc *AuthController) CheckUID(ctx context.Context) (user int, ok bool) {

	user, ok = ctx.Value(auc.UIDkey).(int)
	return
}
