package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"time"
)

// EncodeToken encodes the desired information into a
// signed JWT.
func EncodeToken(id uuid.UUID, secret string) (string, error){

	claims := jwt.MapClaims{}
	claims["user_id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24)

	// encode token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign encoded token
	signedToken, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ValidateToken validates incoming JWT token
func ValidateToken(encodedToken, secret string) (jwt.MapClaims , error) {

	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, valid := token.Method.(*jwt.SigningMethodHMAC); !valid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil || !token.Valid{
		return nil, err
	}

	return token.Claims.(jwt.MapClaims), nil

}