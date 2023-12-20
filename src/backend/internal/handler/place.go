package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/thirteenths/WEB_BMSTU23/backend/internal/model"
	"net/http"
	"strconv"
)

// GetAllPlace godoc
// @Summary Show All info about places.
// @Description get places.
// @ID get-all-place
// @Tags places
// @Accept */*
// @Produce json
// @Success 200 {object} getAllEventsResponse
// @Failure 500 {object} errorResponse
// @Router /places [get]
func (h *Handler) GetAllPlace(c *gin.Context) {
	places, err := h.service.IPlace.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllPlacesResponse{
		Data: places,
	})
}

// GetPlace godoc
// @Summary Show info about place.
// @Description get place by ID.
// @ID get-place
// @Tags places
// @Accept */*
// @Produce json
// @Param id path string true "place ID"
// @Success 200 {object} modelUI.Place
// @Failure 500 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Router /places/{id} [get]
func (h *Handler) GetPlace(c *gin.Context) {
	placeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	place, err := h.service.IPlace.Get(placeId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, place)
}

// CreatePlace godoc
// @Summary Create new place.
// @Description post place
// @ID create-place
// @Tags places
// @Accept */*
// @Produce json
// @Param data body modelUI.Place true "modelUI.place"
// @Success 201 {object} statusResponse
// @Failure 500 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Failure 403 {object} errorResponse
// @Router /places [post]
func (h *Handler) CreatePlace(c *gin.Context) {
	var place model.Place

	if err := c.BindJSON(&place); err != nil {
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

	placeId, err := h.service.IPlace.Add(place)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id_place": placeId,
	})
}

// DeletePlace godoc
// @Summary Delete a place item by ID.
// @Description delete place by ID
// @ID delete-place
// @Tags places
// @Accept */*
// @Produce json
// @Param id path string true "place ID"
// @Success 200 {object} statusResponse
// @Failure 500 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Failure 403 {object} errorResponse
// @Router /places/{id} [delete]
func (h *Handler) DeletePlace(c *gin.Context) {
	roleUser, err := getUserRole(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if roleUser != 0 {
		newErrorResponse(c, http.StatusForbidden, err.Error())
	}

	placeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	err = h.service.IPlace.Delete(placeId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// UpdatePlace godoc
// @Summary Update info about place item by ID.
// @Description update place by ID
// @ID update-place
// @Tags places
// @Accept */*
// @Produce json
// @Param id path string true "place ID"
// @Param data body modelUI.Place true "modelUI.place"
// @Success 200 {object} statusResponse
// @Failure 500 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Failure 403 {object} errorResponse
// @Router /places/{id} [put]
func (h *Handler) UpdatePlace(c *gin.Context) {
	placeId, err := strconv.Atoi(c.Param("id"))
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

	var place model.Place
	if err := c.BindJSON(&place); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.IPlace.Update(placeId, place)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
