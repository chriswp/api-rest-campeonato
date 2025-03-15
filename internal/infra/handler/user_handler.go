package handler

import (
	"github.com/chriswp/api-rest-campeonato/internal/constants"
	"github.com/chriswp/api-rest-campeonato/internal/domain/repository"
	"github.com/chriswp/api-rest-campeonato/internal/dto"
	"github.com/gin-gonic/gin"
	"github.com/go-chi/jwtauth"
	"net/http"
	"time"
)

type UserHandler struct {
	UserRepository repository.UserRepository
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/login", h.GetJWT)
}

func (h *UserHandler) GetJWT(c *gin.Context) {
	jwt, ok := c.Value("jwt").(*jwtauth.JWTAuth)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": constants.ErrJWTNotFound})
		return
	}

	jwtExpiresIn, ok := c.Value("JwtExpiresIn").(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": constants.ErrInvalidJwtExp})
		return
	}

	var user dto.GetJWTInput
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrInvalidData})
		return
	}

	u, err := h.UserRepository.FindUserByEmail(c, user.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": constants.ErrUserNotFound})
		return
	}

	if !u.CheckPassword(user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": constants.ErrInvalidPassword})
		return
	}

	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})

	c.JSON(http.StatusOK, dto.GetJWTOutput{Token: tokenString})
}
