package util

import "time"
import "github.com/golang-jwt/jwt/v4"

type MyCustomClaims struct {
	jwt.RegisteredClaims
	ID int64 `json:"id"`
}

func GenJWT(id int64) (string, error) {
	mySigningKey := []byte("~a1a2a1a3a8&")
	// Create the Claims
	myClaims := &MyCustomClaims{
		jwt.RegisteredClaims{
			Issuer:    "test",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
		},
		id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)
	return token.SignedString(mySigningKey)
}

func ParseJWT(tokenString string) (*MyCustomClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (any, error) {
		return []byte("~a1a2a1a3a8&"), nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
