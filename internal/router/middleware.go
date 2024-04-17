package router

import (
	"errors"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/leonardonicola/tickethub/internal/domain"
	"github.com/leonardonicola/tickethub/internal/dto"
)

func InitAuthMiddleware() (*jwt.GinJWTMiddleware, error) {
	// Load godot to retrieve variables from .env
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	// Lookup for variables and return error if coulnd't find
	JWT_SECRET, ok := os.LookupEnv("JWT_SECRET")
	if !ok {
		return nil, errors.New("failed getting jwt secret from env variables")
	}

	// Set identity key for JWT
	identityKey := "id"

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "tickethub-authentication",
		Key:         []byte(JWT_SECRET),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*domain.User); ok {
				return jwt.MapClaims{
					identityKey: v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &domain.User{
				ID: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals dto.Login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			email := loginVals.Email
			password := loginVals.Password

			if email == "admin@admin.com" && password == "admin" {
				return &domain.User{
					ID:    "718d3a7e-cb2b-43b5-a04f-bf312a60d853",
					Email: email,
					Name:  "Admin",
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
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
