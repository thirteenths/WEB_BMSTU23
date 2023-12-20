package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/thirteenths/WEB_BMSTU23/backend/internal/model"
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
	Data []model.Comic `json:"comics"`
}

type getAllEventsResponse struct {
	Data []model.Event `json:"events"`
}

type getAllPostersResponse struct {
	Data []model.Poster `json:"posters"`
}

type getAllPlacesResponse struct {
	Data []model.Place `json:"places"`
}
