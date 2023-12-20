package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	userR               = "userRole"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, userRole, err := h.service.IAuthorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	logrus.Printf("Parse token id: %d, role: %d", userId, userRole)
	c.Set(userCtx, userId)
	c.Set(userR, userRole)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}

func getUserRole(c *gin.Context) (int, error) {
	role, ok := c.Get(userR)
	if !ok {
		return 0, errors.New("user role not found")
	}

	roleInt, ok := role.(int)
	if !ok {
		return 0, errors.New("user role is of invalid type")
	}

	return roleInt, nil
}
