package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/thirteenths/WEB_BMSTU23/backend/internal/modelUI"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}

type getAllComicsResponse struct {
	Data []modelUI.Comic `json:"comics"`
}

type getAllEventsResponse struct {
	Data []modelUI.Event `json:"events"`
}

type getAllPostersResponse struct {
	Data []modelUI.Poster `json:"posters"`
}

type getAllPlacesResponse struct {
	Data []modelUI.Place `json:"places"`
}
