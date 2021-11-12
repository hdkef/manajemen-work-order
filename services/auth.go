package services

import (
	"errors"
	"fmt"
	"manajemen-work-order/models"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var SECRET string

func init() {
	godotenv.Load()
	SECRET = os.Getenv("SECRET")
}

func GenerateTokenFromEntity(entity *models.Entity) (string, error) {
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":        entity.ID,
		"Fullname":  entity.Fullname,
		"Username":  entity.Username,
		"Email":     entity.Email,
		"Role":      entity.Role,
		"Signature": entity.Signature,
	})
	signedToken, err := newToken.SignedString([]byte(SECRET))
	if err != nil {
		return "", err
	} else {
		return signedToken, nil
	}
}

func CompareTwoPassword(password *string, hashedPassword *string) error {
	return bcrypt.CompareHashAndPassword([]byte(*hashedPassword), []byte(*password))
}

func HashPassword(pass *string) (string, error) {
	passByte, err := bcrypt.GenerateFromPassword([]byte(*pass), 5)
	if err != nil {
		return "", err
	}
	return string(passByte), nil
}

func validateTokenString(token *string) (models.Entity, error) {
	//parse token string to jwt.Token
	parsedToken, err := jwt.Parse(*token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error")
		}
		return []byte(SECRET), nil
	})
	if err != nil {
		return models.Entity{}, err
	}

	if !parsedToken.Valid {
		return models.Entity{}, errors.New("token invalid")
	}
	mapclaims := parsedToken.Claims.(jwt.MapClaims)

	return models.Entity{
		ID:        int64(mapclaims["ID"].(float64)),
		Fullname:  mapclaims["Fullname"].(string),
		Username:  mapclaims["Username"].(string),
		Email:     mapclaims["Email"].(string),
		Role:      mapclaims["Role"].(string),
		Signature: mapclaims["Signature"].(string),
	}, nil
}

func ValidateTokenFromHeader(c *gin.Context) (models.Entity, error) {
	token := c.GetHeader("Authorization")

	return validateTokenString(&token)
}
