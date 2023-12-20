package mock_service

import (
	"github.com/golang/mock/gomock"
	"github.com/thirteenths/WEB_BMSTU23/backend/internal/service"
)

func NewService(comic service.IComic, ctrl *gomock.Controller) *service.Service {
	return &service.Service{
		IComic:         comic,
		IPoster:        NewMockIPoster(ctrl),
		IAuthorization: NewMockIAuthorization(ctrl),
		IEvent:         NewMockIEvent(ctrl),
		IPlace:         NewMockIPlace(ctrl),
	}
}
