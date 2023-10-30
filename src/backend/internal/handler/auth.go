package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/thirteenths/WEB_BMSTU23/backend/internal/modelUI"
	"net/http"
)

// SingUp godoc
// @Summary Create new person.
// @Description post person
// @ID create-person
// @Tags auth
// @Accept */*
// @Produce json
// @Param data body modelUI.Person true "modelUI.Person"
// @Success 200 {object} int
// @Failure 500 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Router /auth/singUp [post]
func (h *Handler) singUp(c *gin.Context) {
	var input modelUI.Person

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.IAuthorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

type singInInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// SingIn godoc
// @Summary Sing in person.
// @Description post person
// @ID sing-in-person
// @Tags auth
// @Accept */*
// @Produce json
// @Param data body singInInput true "singInInput"
// @Success 200 {object} string
// @Failure 500 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Router /auth/singIn [post]
func (h *Handler) singIn(c *gin.Context) {
	var input singInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.IAuthorization.GenerationToken(input.Login, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
