package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/thirteenths/WEB_BMSTU23/backend/internal/model"
	"net/http"
	"strconv"
)

// GetAllPoster godoc
// @Summary Show All info about posters.
// @Description get posters.
// @ID get-all-poster
// @Tags posters
// @Accept */*
// @Produce json
// @Success 200 {object} getAllPostersResponse
// @Failure 500 {object} errorResponse
// @Router /posters [get]
func (h *Handler) GetAllPoster(c *gin.Context) {
	posters, err := h.service.IPoster.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllPostersResponse{
		Data: posters,
	})
}

// GetPoster godoc
// @Summary Show info about poster.
// @Description get poster by ID.
// @ID get-poster
// @Tags posters
// @Accept */*
// @Produce json
// @Param id path string true "poster ID"
// @Success 200 {object} modelUI.Poster
// @Failure 500 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Router /posters/{id} [get]
func (h *Handler) GetPoster(c *gin.Context) {
	posterId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	poster, err := h.service.IPoster.Get(posterId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, poster)
}

// CreatePoster godoc
// @Summary Create new poster.
// @Description post poster
// @ID create-poster
// @Tags posters
// @Accept */*
// @Produce json
// @Param data body modelUI.Poster true "modelUI.Poster"
// @Success 200 {object} int
// @Failure 500 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Failure 403 {object} errorResponse
// @Router /posters [post]
func (h *Handler) CreatePoster(c *gin.Context) {
	var poster model.Poster
	if err := c.BindJSON(&poster); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	roleUser, err := getUserRole(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if roleUser != 0 {
		newErrorResponse(c, http.StatusForbidden, err.Error())
	}

	posterId, err := h.service.IPoster.Add(poster)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id_poster": posterId,
	})
}

// DeletePoster godoc
// @Summary Delete a poster item by ID.
// @Description delete poster by ID
// @ID delete-poster
// @Tags posters
// @Accept */*
// @Produce json
// @Param id path string true "poster ID"
// @Success 200 {object} statusResponse
// @Failure 500 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Failure 403 {object} errorResponse
// @Router /posters/{id} [delete]
func (h *Handler) DeletePoster(c *gin.Context) {
	roleUser, err := getUserRole(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if roleUser != 0 {
		newErrorResponse(c, http.StatusForbidden, err.Error())
	}

	posterId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	err = h.service.IPoster.Delete(posterId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// CheckInPoster godoc
// @Summary CheckIn poster item by ID.
// @Description check in comic by ID
// @ID check-in-poster
// @Tags posters
// @Accept */*
// @Produce json
// @Param id path string true "poster ID"
// @Success 200 {object} statusResponse
// @Failure 500 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Router /posters/{id} [post]
func (h *Handler) CheckInPoster(c *gin.Context) {
	idUser, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	roleUser, err := getUserRole(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if roleUser != 1 {
		newErrorResponse(c, http.StatusForbidden, err.Error())
	}

	posterId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	err = h.service.IPoster.CheckIn(posterId, idUser)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// UpdatePoster godoc
// @Summary Update info about poster item by ID.
// @Description update poster by ID
// @ID update-poster
// @Tags posters
// @Accept */*
// @Produce json
// @Param id path string true "poster ID"
// @Param data body modelUI.Poster true "modelUI.poster"
// @Success 200 {object} statusResponse
// @Failure 500 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Router /posters/{id} [put]
func (h *Handler) UpdatePoster(c *gin.Context) {
	posterId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	roleUser, err := getUserRole(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if roleUser != 0 {
		newErrorResponse(c, http.StatusForbidden, err.Error())
	}

	var poster model.Poster
	if err := c.BindJSON(&poster); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.IPoster.Update(posterId, poster)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
