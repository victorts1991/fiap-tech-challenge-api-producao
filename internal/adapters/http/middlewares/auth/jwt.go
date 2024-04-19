package auth

import (
	serverErr "fiap-tech-challenge-api/internal/adapters/http/error"
	"fiap-tech-challenge-api/internal/core/commons"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"os"
	"strings"
	"time"
)

//go:generate mockgen -source=$GOFILE -package=mock_auth -destination=../../../../test/mock/auth/$GOFILE

type Token interface {
	GenerateToken(email string) (string, error)
	VerifyToken(next echo.HandlerFunc) echo.HandlerFunc
}

type JwtToken struct {
}

func NewJwtToken() Token {
	return &JwtToken{}
}

type jwtCustomClaims struct {
	Cpf     string `json:"cpf"`
	IsAdmin bool
	jwt.RegisteredClaims
}

func (jt *JwtToken) GenerateToken(cpf string) (string, error) {
	secret := []byte(os.Getenv("AUTH_SECRET"))
	claims := &jwtCustomClaims{
		cpf,
		true,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return t, nil
}

func (jt *JwtToken) VerifyToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenStr := jt.getToken(c)
		if tokenStr == "" {
			return serverErr.HandleError(c, commons.Unauthorized.New("authentication key not found"))
		}

		secret := []byte(os.Getenv("AUTH_SECRET"))
		tkn, err := jwt.ParseWithClaims(tokenStr, &jwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})
		if err != nil {
			return commons.Unauthorized.New(err.Error())
		}

		if !tkn.Valid {
			return commons.Unauthorized.New("authentication is not valid")
		}

		return next(c)
	}
}

func (jt *JwtToken) getToken(c echo.Context) string {
	tokenStr := ""
	if bearer := c.Request().Header.Get("Authorization"); bearer != "" {
		if strings.Contains(bearer, "Bearer") {
			tokenStr = strings.Split(bearer, " ")[1]
		}
	}

	if tokenStr == "" {
		tokenStr = c.Request().Header.Get("token")
	}

	return tokenStr
}
