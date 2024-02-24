package service

import (
	"github.com/thirteenths/WEB_BMSTU23/backend/internal/model"
	"github.com/thirteenths/WEB_BMSTU23/backend/pkg/storage"
)

type PlaceService struct {
	storage storage.IStorage
}

func NewPlaceService(storage storage.IStorage) *PlaceService {
	return &PlaceService{storage: storage}
}

func (s *PlaceService) GetAll() ([]model.Place, error) {
	var placesUI []model.Place

	places, err := s.storage.Repositories()[storage.PLACE].SelectAll()

	if err != nil {
		return nil, err
	}

	for _, place := range places {
		placesUI = append(placesUI, transportPlaceModel(place))
	}
	return placesUI, nil
}
func (s *PlaceService) Get(id int) (model.Place, error) {
	place, err := s.storage.Repositories()[storage.PLACE].SelectById(id)

	if err != nil {
		return model.Place{}, err
	}

	return transportPlaceModel(place), nil
}
func (s *PlaceService) Add(place model.Place) (int, error) {
	placeId, err := s.storage.Repositories()[storage.PLACE].Insert(place)
	if err != nil {
		return 0, err
	}
	return placeId, nil
}
func (s *PlaceService) Delete(id int) error {
	return s.storage.Repositories()[storage.PLACE].DeleteById(id)
}
func (s *PlaceService) Update(id int, place model.Place) error {
	return s.storage.Repositories()[storage.PLACE].UpdateById(id, place)
}

func transportPlaceModel(fields storage.Model) model.Place {
	return model.Place{
		Id:          int(fields.([]interface{})[0].(int32)),
		Name:        fields.([]interface{})[1].(string),
		CountTicket: int(fields.([]interface{})[2].(int32)),
	}
}
