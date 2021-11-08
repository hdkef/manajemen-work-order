package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

type User struct {
	ID       int64  `json:"user_id"`
	FullName string `json:"user_full_name"`
	Username string `json:"user_username"`
	Password string `json:"user_password"`
	Role     string `json:"user_role"`
}

const AUTH_COOKIE_NAME = "Authorization"
const ERR_NEED_REFRESH_TOKEN = "ERR_NEED_REFRESH_TOKEN"
const REFRESH_TOKEN_DIFF = 10000
const TOKEN_EXPIRES_DUR = 100000

var SECRET string

func init() {
	_ = godotenv.Load()
	SECRET = os.Getenv("SECRET")
}

func (u *User) Authenticate(c *gin.Context) error {
	err := json.NewDecoder(c.Request.Body).Decode(u)
	if err != nil {
		return err
	}
	////////TOBEIMPLEMENTED
	//get user detail from db
	//compare hashed password
	//if valid, set role and generate token and save to cookies
	u.ID = 1
	u.FullName = "full name"
	u.Role = "PUM"
	err = u.generateTokenAndSave(c)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) ValidateTokenStringGetUser(token *string) error {
	//parse token string to jwt.Token
	parsedToken, err := jwt.Parse(*token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error")
		}
		return []byte(SECRET), nil
	})
	if err != nil {
		return err
	}

	if !parsedToken.Valid {
		return errors.New("token invalid")
	}

	renew := checkTokenRenew(parsedToken)
	mapclaims := parsedToken.Claims.(jwt.MapClaims)
	if renew {
		//send refreshToken
		//cast jwt.MapClaims from parsedToken.Claims
		return errors.New(ERR_NEED_REFRESH_TOKEN)
	}

	id := int64(mapclaims["ID"].(float64))
	u.ID = id
	u.Username = mapclaims["Username"].(string)
	u.FullName = mapclaims["FullName"].(string)
	u.Role = mapclaims["Role"].(string)

	return nil
}

//generateTokenAndSave will create a new token and save to cookies
func (u *User) generateTokenAndSave(c *gin.Context) error {
	claims := newClaimsMap(u)

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := newToken.SignedString([]byte(SECRET))
	if err != nil {
		return err
	}

	SaveTokenCookie(c, &signedToken)
	return nil
}

//create jwt map claims from user struct
func newClaimsMap(user *User) jwt.MapClaims {
	var claims jwt.MapClaims = make(jwt.MapClaims)

	claims["Username"] = user.Username
	claims["FullName"] = user.FullName
	claims["Role"] = user.Role
	claims["ID"] = user.ID

	claims["exp"] = time.Now().Unix() + int64(TOKEN_EXPIRES_DUR)
	//this code is intended to be place after for loop so that new exp override old exp for refresh token

	return claims
}

//validate token from cookies, if it needs refreshment then send new token if invalid then return error if valid cast jwt maps into user struct
func (u *User) ValidateTokenFromCookies(c *gin.Context) error {
	//get token string from cookies
	tokenString, err := c.Cookie(AUTH_COOKIE_NAME)
	if err != nil {
		removeTokenCookie(c)
		return err
	}
	mapclaims, err := validateTokenString(&tokenString)
	if err != nil {
		if err.Error() == ERR_NEED_REFRESH_TOKEN {
			tokenString, err := createRefreshToken(mapclaims)
			if err != nil {
				return err
			}
			SaveTokenCookie(c, &tokenString)
			//cast jwt to Auth struct
			//jwt maps claims pasti anggap id itu float64, makanya harus di convert dulu ke int64
			id := int64((*mapclaims)["ID"].(float64))
			u.ID = id
			u.Username = (*mapclaims)["Username"].(string)
			u.FullName = (*mapclaims)["FullName"].(string)
			u.Role = (*mapclaims)["Role"].(string)
			return nil
		}
		removeTokenCookie(c)
		return err
	}
	//cast jwt map to Auth struct
	//jwt maps claims pasti anggap id itu float64, makanya harus di convert dulu ke int64
	id := int64((*mapclaims)["ID"].(float64))
	u.ID = id
	u.Username = (*mapclaims)["Username"].(string)
	u.FullName = (*mapclaims)["FullName"].(string)
	u.Role = (*mapclaims)["Role"].(string)
	return nil
}

func removeTokenCookie(c *gin.Context) {
	c.SetCookie(AUTH_COOKIE_NAME, "", -1, "/", "", false, false)
}

func SaveTokenCookie(c *gin.Context, tokenString *string) {
	c.SetCookie(AUTH_COOKIE_NAME, *tokenString, TOKEN_EXPIRES_DUR, "/", "", false, false)
}

func validateTokenString(token *string) (*jwt.MapClaims, error) {
	//parse token string to jwt.Token
	parsedToken, err := jwt.Parse(*token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error")
		}
		return []byte(SECRET), nil
	})
	if err != nil {
		return nil, err
	}

	if !parsedToken.Valid {
		return nil, errors.New("token invalid")
	}

	renew := checkTokenRenew(parsedToken)
	mapclaims := parsedToken.Claims.(jwt.MapClaims)
	if renew {
		//send refreshToken
		//cast jwt.MapClaims from parsedToken.Claims
		return &mapclaims, errors.New(ERR_NEED_REFRESH_TOKEN)
	}

	return &mapclaims, nil
}

//checkTokenRenew will return true if token expiration time between range that need to be renewed
func checkTokenRenew(token *jwt.Token) bool {

	now := time.Now().Unix()
	timeSubNow := (*token).Claims.(jwt.MapClaims)["exp"].(float64) - float64(now)

	return timeSubNow <= float64(REFRESH_TOKEN_DIFF)
}

//create refresh token
func createRefreshToken(mapclaims *jwt.MapClaims) (string, error) {
	//From jwt.MapClaims
	tokenString, err := generateTokenStringFromMapClaims(mapclaims)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//generate new token from map claims
func generateTokenStringFromMapClaims(mapclaims *jwt.MapClaims) (string, error) {
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, mapclaims)

	signedToken, err := newToken.SignedString([]byte(SECRET))
	if err != nil {
		return "", err
	} else {
		return signedToken, nil
	}
}
