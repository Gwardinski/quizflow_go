package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/pascaldekloe/jwt"
)

// Creates JWT, used in login() views.auth.go
func createJWT(UID int, secret string) (token []byte, err error) {
	var claims jwt.Claims
	claims.Subject = fmt.Sprint(UID)
	claims.Issued = jwt.NewNumericTime(time.Now())
	claims.NotBefore = jwt.NewNumericTime(time.Now())
	claims.Expires = jwt.NewNumericTime(time.Now().Add(72 * time.Hour))
	claims.Issuer = "domain.com"
	claims.Audiences = []string{"domain.com"}
	return claims.HMACSign(jwt.HS256, []byte(secret))
}

// Checks JWT and returns User.ID, used in views.*.go
func (app *application) getUserIdFromJWT(r *http.Request) (int, error) {
	ah := "Authorization"

	authHeader := r.Header.Get(ah)

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		return 0, errors.New("invalid auth header")
	}
	if headerParts[0] != "Bearer" {
		return 0, errors.New("no bearer")
	}

	token := headerParts[1]

	claims, err := jwt.HMACCheck([]byte(token), []byte(app.config.jwt.secret))

	if err != nil {
		return 0, errors.New(err.Error())
	}

	if !claims.Valid(time.Now()) {
		return 0, errors.New("token expired")
	}

	if !claims.AcceptAudience("domain.com") {
		return 0, errors.New("invalid audience")
	}

	if claims.Issuer != "domain.com" {
		return 0, errors.New("invalid issuer")
	}

	userId, err := strconv.ParseInt(claims.Subject, 10, 64)
	if err != nil {
		return 0, errors.New("token err 7")
	}

	return int(userId), nil
}
