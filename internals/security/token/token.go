package token

import (
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type TokenMetadata struct {
	UserId float64
	Exp    int64
}

func ExtractTokenMetada(c *fiber.Ctx) (*TokenMetadata, error) {
	token, err := verifyToken(c)
	if err != nil {
		return nil, err
	}

	// Setting and checking token and credentials.
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		// Expires time.
		id := claims["user_id"].(float64)
		// fullname := claims["full_name"].(string)
		exp := int64(claims["exp"].(float64))

		return &TokenMetadata{
			UserId: id,
			Exp:    exp,
			// FullName:  fullname,
		}, nil
	}

	return nil, err

}

func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractToken(c)

	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func extractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")

	// Normally Authorization HTTP header.
	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	// env := config.New()
	return []byte(os.Getenv("JWT_SECRET_KEY")), nil
}
