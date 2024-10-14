package router

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/haqisaurus/poskita/util"
)
 

func FirstAuthMiddleware(c *fiber.Ctx) error {
	publicKey, err := util.LoadPublicKey(os.Getenv("PUBLIC_KEY"))
	if err != nil {
		// log.Fatalf("Error loading private key: %v", err)
		return c.Status(fiber.StatusUnauthorized).JSON(util.GenerateError(util.ErrUnauthorized, err.Error()))

	}
	tokenStr := c.Get("Authorization")
	tokenStr = strings.ReplaceAll(tokenStr, "Bearer ", "")

	// Parse the token
	// Parse the token
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Validate that the signing method is RS256
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the public key for validation
		return publicKey, nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(util.GenerateError(util.ErrUnauthorized, err.Error()))
	}
	// Check if the token is valid
	if !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(util.GenerateError(util.ErrUnauthorized, "Token is invalid"))

	}

	// Save the token in the context for use in the handler
	c.Locals("user", token)

	// Call the next handler
	return c.Next()
}
 
func CompanyAuthMiddleware(c *fiber.Ctx) error {
	publicKey, err := util.LoadPublicKey(os.Getenv("PUBLIC_KEY"))
	if err != nil {
		// log.Fatalf("Error loading private key: %v", err)
		return c.Status(fiber.StatusUnauthorized).JSON(util.GenerateError(util.ErrUnauthorized, err.Error()))

	}
	tokenStr := c.Get("Authorization")
	tokenStr = strings.ReplaceAll(tokenStr, "Bearer ", "")

	// Parse the token
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Validate that the signing method is RS256
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the public key for validation
		return publicKey, nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(util.GenerateError(util.ErrUnauthorized, err.Error()))
	}
	// Check if the token is valid
	if !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(util.GenerateError(util.ErrUnauthorized, "Token is invalid"))

	}
	claims := token.Claims.(jwt.MapClaims)
	if claims["companyId"] == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(util.GenerateError(util.ErrUnauthorized, "Token is invalid, not for access"))
	}
	// Save the token in the context for use in the handler
	c.Locals("user", token)

	// Call the next handler
	return c.Next()
}
 