package util

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/haqisaurus/poskita/dto"
)
const (
    ErrInvalidInput = 1001
    ErrNotFound     = 1002
    ErrUnauthorized = 1003
    ErrInternal     = 2001
)

// Create an error message map to map error codes to error messages
var errorMessages = map[int]string{
    ErrInvalidInput: "Invalid input provided",
    ErrNotFound:     "Resource not found",
    ErrUnauthorized: "Unauthorized access",
    ErrInternal:     "Internal server error",
}

func GenerateError(code int, errorMessage string) dto.ResponseError {
	newUUID := uuid.New()
	message := "Unknown error"
	if msg, exists := errorMessages[code]; exists {
        message = msg
    }

	return dto.ResponseError{
		Message:         message,
		Code:            code,
		MessageError:    errorMessage,
		ActivityRefCode: newUUID.String(),
	}
}
 
func GenerateResponse[T any](data T) dto.ResponseSuccess[T] {
	newUUID := uuid.New()
	 
	return dto.ResponseSuccess[T]{
		Message:         "SUCCESS",
		Code:            "0",
		ActivityRefCode: newUUID.String(),
		Data : data,
	}
}

func LoadPrivateKey(filePath string) (*rsa.PrivateKey, error) {
	keyData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyData)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing the private key")
	}

	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func LoadPublicKey(filePath string) (*rsa.PublicKey, error) {
	keyData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyData)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing the public key")
	}

	// return x509.ParsePKIXPublicKey(block.Bytes).(*rsa.PublicKey), nil
	pubKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// Type assert to *rsa.PublicKey
	publicKey, ok := pubKeyInterface.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("not an RSA public key")
	}

	return publicKey, nil
}

func GetUser(c *fiber.Ctx) jwt.MapClaims {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims
}
