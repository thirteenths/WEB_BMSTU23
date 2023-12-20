package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/thirteenths/WEB_BMSTU23/backend/internal/model"
	"net/http"
	"strconv"
)

// GetAllEvent godoc
// @Summary Show All info about events.
// @Description get events.
// @ID get-all-event
// @Tags events
// @Accept */*
// @Produce json
// @Success 200 {object} getAllEventsResponse
// @Failure 500 {object} errorResponse
// @Router /events [get]
func (h *Handler) GetAllEvent(c *gin.Context) {
	events, err := h.service.IEvent.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllEventsResponse{
		Data: events,
	})
}

// GetEvent godoc
// @Summary Show info about event.
// @Description get event by ID.
// @ID get-event
// @Tags events
// @Accept */*
// @Produce json
// @Param id path string true "event ID"
// @Success 200 {object} modelUI.Event
// @Failure 500 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Router /events/{id} [get]
func (h *Handler) GetEvent(c *gin.Context) {
	eventsId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	event, err := h.service.IEvent.Get(eventsId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, event)
}

// CreateEvent godoc
// @Summary Create new event.
// @Description post event
// @ID create-event
// @Tags events
// @Accept */*
// @Produce json
// @Param data body modelUI.Event true "modelUI.event"
// @Success 201 {object} int
// @Failure 500 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Failure 403 {object} errorResponse
// @Router /events [post]
func (h *Handler) CreateEvent(c *gin.Context) {
	var event model.Event

	if err := c.BindJSON(&event); err != nil {
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

	eventId, err := h.service.IEvent.Add(event)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id_event": eventId,
	})
}

// DeleteEvent godoc
// @Summary Delete a event item by ID.
// @Description delete event by ID
// @ID delete-event
// @Tags events
// @Accept */*
// @Produce json
// @Param id path string true "event ID"
// @Success 200 {object} statusResponse
// @Failure 500 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Failure 403 {object} errorResponse
// @Router /events/{id} [delete]
func (h *Handler) DeleteEvent(c *gin.Context) {
	roleUser, err := getUserRole(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if roleUser != 0 {
		newErrorResponse(c, http.StatusForbidden, err.Error())
	}

	eventId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	err = h.service.IEvent.Delete(eventId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// UpdateEvent godoc
// @Summary Update info about event item by ID.
// @Description update poster by ID
// @ID update-event
// @Tags events
// @Accept */*
// @Produce json
// @Param id path string true "event ID"
// @Param data body modelUI.Event true "modelUI.event"
// @Success 200 {object} statusResponse
// @Failure 500 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Failure 403 {object} errorResponse
// @Router /events/{id} [put]
func (h *Handler) UpdateEvent(c *gin.Context) {
	eventId, err := strconv.Atoi(c.Param("id"))
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

	var event model.Event
	if err := c.BindJSON(&event); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.IEvent.Update(eventId, event)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
