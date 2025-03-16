package config

import (
	"fmt"
	"github.com/chriswp/api-rest-campeonato/internal/constants"
	"github.com/chriswp/api-rest-campeonato/internal/domain/usecase"
	"time"

	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func NewAuthMiddleware(authUseCase usecase.AuthUseCase) (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "api",
		Key:         []byte(Envs.JWTSecret),
		Timeout:     time.Duration(Envs.JWTExpiresIn) * time.Second,
		MaxRefresh:  time.Hour,
		IdentityKey: "user_id",
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var login struct {
				Email    string `json:"email"`
				Password string `json:"password"`
			}
			if err := c.ShouldBindJSON(&login); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			if login.Email == "" {
				return nil, fmt.Errorf(constants.RequiredField("email"))
			}
			if login.Password == "" {
				return nil, fmt.Errorf(constants.RequiredField("password"))
			}

			userData, err := authUseCase.Authenticate(c.Request.Context(), login.Email, login.Password)
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}
			return userData, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(map[string]interface{}); ok && v["role"] == "admin" {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{"error": message})
		},
		TokenLookup:   "header: Authorization, query: token",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		return nil, err
	}

	return authMiddleware, nil
}
