package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thirteenths/WEB_BMSTU23/backend/internal/model"
)

// GetAllComic godoc
// @Summary Show All info about comics.
// @Description get comics.
// @ID get-all-comic
// @Tags comics
// @Accept */*
// @Produce json
// @success 200 {object} getAllComicsResponse
// @Failure 500 {object} errorResponse{message=error}
// @Router /comics [get]
func (h *Handler) GetAllComic(c *gin.Context) {
	comics, err := h.service.IComic.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllComicsResponse{
		Data: comics,
	})
}

// GetComic godoc
// @Summary Show info about comic.
// @Description get comic by ID.
// @ID get-comic
// @Tags comics
// @Accept */*
// @Produce json
// @Param id path string true "comic ID"
// @Success 200 {object} model.Comic
// @Failure 500 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Router /comics/{id} [get]
func (h *Handler) GetComic(c *gin.Context) {
	comicId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	comic, err := h.service.IComic.Get(comicId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, comic)
}

// CreateComic godoc
// @Summary Create new comic.
// @Description post comic
// @ID create-comic
// @Tags comics
// @Accept */*
// @Produce json
// @Param data body model.Comic true "model.Comic"
// @Success 201 {object} statusResponse
// @Failure 500 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Failure 403 {object} errorResponse
// @Router /comics [post]
func (h *Handler) CreateComic(c *gin.Context) {
	var comic model.Comic
	if err := c.BindJSON(&comic); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	roleUser, err := getUserRole(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if roleUser != 0 {
		newErrorResponse(c, http.StatusForbidden, "Access is denied")
		return
	}

	comicId, err := h.service.IComic.Add(comic)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id_comic": comicId,
	})
}

// DeleteComic godoc
// @Summary Delete a comic item by ID.
// @Description delete comic by ID
// @ID delete-comic
// @Tags comics
// @Accept */*
// @Produce json
// @Param id path string true "comic ID"
// @Success 200 {object} statusResponse
// @Failure 500 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Failure 403 {object} errorResponse
// @Router /comics/{id} [delete]
func (h *Handler) DeleteComic(c *gin.Context) {
	roleUser, err := getUserRole(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if roleUser != 0 {
		newErrorResponse(c, http.StatusForbidden, "Access is denied")
		return
	}

	comicId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	err = h.service.IComic.Delete(comicId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// UpdateComic godoc
// @Summary Update info about comic item by ID.
// @Description update comic by ID
// @ID update-comic
// @Tags comics
// @Accept */*
// @Produce json
// @Param id path string true "comic ID"
// @Param data body model.Comic true "model.comic"
// @Success 200 {object} statusResponse
// @Failure 500 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Failure 403 {object} errorResponse
// @Router /comics/{id} [put]
func (h *Handler) UpdateComic(c *gin.Context) {
	comicId, err := strconv.Atoi(c.Param("id"))
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
		newErrorResponse(c, http.StatusForbidden, "Access is denied")
		return
	}

	var comic model.Comic
	if err := c.BindJSON(&comic); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.IComic.Update(comicId, comic)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// UpdateComicKek godoc
// @Summary Update count of kek comic item by ID.
// @Description update comic by ID
// @ID update-comic
// @Tags comics
// @Accept */*
// @Produce json
// @Param id path string true "comic ID"
// @Param data body string true "string"
// @Success 200 {object} statusResponse
// @Failure 500 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Router /comics/{id} [patch]
func (h *Handler) UpdateComicKek(c *gin.Context) {
	comicId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	roleUser, err := getUserRole(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if roleUser != 1 {
		newErrorResponse(c, http.StatusForbidden, "Access is denied")
		return
	}

	var comic model.Comic
	if err := c.BindJSON(&comic); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.IComic.Update(comicId, comic)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
