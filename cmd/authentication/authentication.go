package authentication

import (
	"errors"
	"time"

	"github.com/ILLIDOM/gin-webapp/cmd/domain"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("secretkey") // TODO - extract to env var

type JWTClaim struct {
	UserID         int           `json:"userid"`
	Username       string        `json:"username"`
	Roles          []domain.Role `json:"roles"`
	StandardClaims jwt.RegisteredClaims
}

func (claims JWTClaim) Valid() error {
	var now = time.Now()
	if claims.StandardClaims.VerifyExpiresAt(now, true) {
		return nil
	}
	return errors.New("token is invalid")
}

type LoginRequest struct {
	UserName string `json:"username" form:"UserName" binding:"required"`
	Password string `json:"password" form:"Password" binding:"required"`
}

func GenerateJWT(user domain.User) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		UserID:   user.ID,
		Username: user.Fullname,
		Roles:    user.Roles,
		StandardClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test",
			Subject:   "somebody",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}
