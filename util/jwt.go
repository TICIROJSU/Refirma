package util

import (
	"time"

)

var SECRET_KEY_JWT string
var TIME_EXPIRE_TOKEN int64

// Creamos Token
func GenerarJWT() (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(TIME_EXPIRE_TOKEN) * time.Minute)),
		Issuer:    "mpazp",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(SECRET_KEY_JWT))
}

// Verificamos Token
func VerificarJWT(tokenString string) error {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY_JWT), nil
	})

	if token.Valid {
		return nil
	} else {
		return err
	}

}
