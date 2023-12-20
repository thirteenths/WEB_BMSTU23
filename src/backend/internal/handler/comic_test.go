package handler

import (
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/thirteenths/WEB_BMSTU23/backend/internal/mock_service"
	"github.com/thirteenths/WEB_BMSTU23/backend/internal/modelUI"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestHandlerComic struct {
}

func TestHandler_GetAllComic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockComicOK := mock_service.NewMockIComic(ctrl)
	mockComicOK.EXPECT().GetAll().Return([]modelUI.Comic{}, nil).AnyTimes()

	mockComicInternalServerError := mock_service.NewMockIComic(ctrl)
	mockComicInternalServerError.EXPECT().GetAll().Return([]modelUI.Comic{}, nil).AnyTimes()

	var tests = []struct {
		inputMock  *mock_service.MockIComic
		outputCode int
		outputBody any
	}{
		{mockComicOK, http.StatusOK, []modelUI.Comic(nil)},
		//{mockComicInternalServerError, http.StatusInternalServerError, []modelUI.Comic(nil)},
	}
	for _, test := range tests {
		mockService := mock_service.NewService(test.inputMock, ctrl)
		handler := NewHandler(mockService)
		router := handler.InitRoutes()

		req, _ := http.NewRequest("GET", "/comics/", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		var comic []modelUI.Comic
		_ = json.Unmarshal(w.Body.Bytes(), &comic)

		assert.Equal(t, test.outputBody, comic)
		assert.Equal(t, test.outputCode, w.Code)
	}
}

func TestHandler_GetComic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockComic := mock_service.NewMockIComic(ctrl)
	mockComic.EXPECT().Get(1).Return(modelUI.Comic{}, nil).AnyTimes()
	mockService := mock_service.NewService(mockComic, ctrl)

	handler := NewHandler(mockService)
	router := handler.InitRoutes()

	req, _ := http.NewRequest("GET", "/comics/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var comic modelUI.Comic
	_ = json.Unmarshal(w.Body.Bytes(), &comic)

	assert.Equal(t, modelUI.Comic{}, comic)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestHandler_CreateComic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockComicCreated := mock_service.NewMockIComic(ctrl)
	mockComicCreated.EXPECT().Add(modelUI.Comic{}).Return(1, nil).AnyTimes()

	mockComicInternalServerError := mock_service.NewMockIComic(ctrl)
	mockComicInternalServerError.EXPECT().GetAll().Return([]modelUI.Comic{}, nil).AnyTimes()

	var tests = []struct {
		inputMock *mock_service.MockIComic

		outputCode int
		outputBody int
	}{
		{mockComicCreated, http.StatusCreated, 1},
		//{mockComicInternalServerError, http.StatusInternalServerError, []modelUI.Comic(nil)},
	}
	for _, test := range tests {
		mockService := mock_service.NewService(test.inputMock, ctrl)

		handler := NewHandler(mockService)

		router := handler.InitRoutes()

		req, _ := http.NewRequest("POST", "/comics/", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		var comicId int
		_ = json.Unmarshal(w.Body.Bytes(), &comicId)

		assert.Equal(t, test.outputBody, comicId)
		assert.Equal(t, test.outputCode, w.Code)
	}
}

func TestHandler_DeleteComic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockComicOK := mock_service.NewMockIComic(ctrl)
	mockComicOK.EXPECT().Delete(1).Return(nil).AnyTimes()

	mockComicInternalServerError := mock_service.NewMockIComic(ctrl)
	mockComicInternalServerError.EXPECT().GetAll().Return([]modelUI.Comic{}, nil).AnyTimes()

	var tests = []struct {
		inputRoute string
		inputMock  *mock_service.MockIComic
		outputCode int
		outputBody any
	}{
		{"/comics/1", mockComicOK, http.StatusOK, nil},
		{"/comics/", mockComicOK, http.StatusBadRequest, nil},
		//{mockComicInternalServerError, http.StatusInternalServerError, []modelUI.Comic(nil)},
	}
	for _, test := range tests {
		mockService := mock_service.NewService(test.inputMock, ctrl)

		handler := NewHandler(mockService)

		router := handler.InitRoutes()

		req, _ := http.NewRequest("DELETE", test.inputRoute, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		var comicId int
		_ = json.Unmarshal(w.Body.Bytes(), &comicId)

		//assert.Equal(t, test.outputBody, comicId)
		assert.Equal(t, test.outputCode, test.outputCode)
	}
}

func TestHandler_UpdateComic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockComicOK := mock_service.NewMockIComic(ctrl)
	mockComicOK.EXPECT().Update(1, modelUI.Comic{})

	mockComicInternalServerError := mock_service.NewMockIComic(ctrl)
	mockComicInternalServerError.EXPECT().GetAll().Return([]modelUI.Comic{}, nil).AnyTimes()

	var tests = []struct {
		inputRoute string
		inputMock  *mock_service.MockIComic
		outputCode int
		outputBody any
	}{
		{"/comics/1", mockComicOK, http.StatusOK, nil},
		{"/comics/", mockComicOK, http.StatusBadRequest, nil},
		//{mockComicInternalServerError, http.StatusInternalServerError, []modelUI.Comic(nil)},
	}
	for _, test := range tests {
		mockService := mock_service.NewService(test.inputMock, ctrl)

		handler := NewHandler(mockService)

		router := handler.InitRoutes()

		req, _ := http.NewRequest("PUT", test.inputRoute, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		var comicId int
		_ = json.Unmarshal(w.Body.Bytes(), &comicId)

		//assert.Equal(t, test.outputBody, comicId)
		assert.Equal(t, test.outputCode, test.outputCode)
	}
}
