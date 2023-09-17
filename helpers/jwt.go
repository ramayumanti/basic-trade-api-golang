package helpers

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = "your-256-bit-secret"

func GenerateToken(id string, email string) string {
	// Generate an ECDSA private key
	//privateKey, err := GenerateEcdsaPrivateKey()
	//if err != nil {
	//	return err.Error()
	//}

	claims := jwt.MapClaims{
		"uuid":  id,
		"email": email,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := parseToken.SignedString([]byte(secretKey))
	if err != nil {
		return err.Error()
	}

	return signedToken
}

func VerifyToken(ctx *gin.Context) (interface{}, error) {
	// Generate an ECDSA private key
	//privateKey, err := GenerateEcdsaPrivateKey()
	//if err != nil {
	//	return nil, err
	//}

	errResponse := errors.New("sign in to proceed")
	headerToken := ctx.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errResponse
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(secretKey), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResponse
	}

	return token.Claims.(jwt.MapClaims), nil
}
