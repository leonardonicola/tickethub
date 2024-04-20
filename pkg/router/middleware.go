package router

import (
	"errors"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/leonardonicola/tickethub/config"
	user "github.com/leonardonicola/tickethub/internal/user/domain"
	userDTO "github.com/leonardonicola/tickethub/internal/user/dto"
	"golang.org/x/crypto/bcrypt"
)

func InitAuthMiddleware() (*jwt.GinJWTMiddleware, error) {
	// Lookup for variables and return error if coulnd't find

	JWT_SECRET, ok := os.LookupEnv("JWT_SECRET")
	if !ok {
		return nil, errors.New("failed getting jwt secret from env variables")
	}

	db := config.GetDB()

	// Set identity key for JWT
	identityKey := "id"

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "tickethub-authentication",
		Key:         []byte(JWT_SECRET),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(user.User); ok {
				return jwt.MapClaims{
					identityKey: v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return user.User{
				ID: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals userDTO.Login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			email := loginVals.Email
			password := loginVals.Password

			var user user.User

			err := db.QueryRow("SELECT email, password FROM users WHERE email = $1", email).Scan(&user.Email, &user.Password)
			if err != nil {
				logger.Errorf("AUTH - not found: %v", err)
				return nil, jwt.ErrFailedAuthentication
			}

			if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
				logger.Errorf("AUTH - hashing error: %v", err)
				return nil, jwt.ErrFailedAuthentication
			}

			return user, nil
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup: "header: Authorization, query: token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
	if err != nil {
		return nil, err
	}
	return authMiddleware, nil
}
