package libs

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var mySecrets = []byte(os.Getenv("JWT_KEYS"))

type Claims struct {
	UserId string
	Role   string
	jwt.StandardClaims
}

func NewToken(userId, role string) *Claims {
	return &Claims{
		UserId: userId,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}
}

func (c *Claims) Create() (string, error) {
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return tokens.SignedString(mySecrets)
}

func CheckToken(token string) (*Claims, error) {
	tokens, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(mySecrets), nil
	})
	if err != nil {
		return nil, err
	}
	Claims := tokens.Claims.(*Claims)
	return Claims, err
}
