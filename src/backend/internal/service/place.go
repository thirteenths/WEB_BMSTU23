package service

import (
	"github.com/thirteenths/WEB_BMSTU23/backend/internal/modelDB"
	"github.com/thirteenths/WEB_BMSTU23/backend/internal/modelUI"
	"github.com/thirteenths/WEB_BMSTU23/backend/pkg/storage"
)

type PlaceService struct {
	storage storage.IStorage
}

func NewPlaceService(storage storage.IStorage) *PlaceService {
	return &PlaceService{storage: storage}
}

func (s *PlaceService) GetAll() ([]modelUI.Place, error) {
	var placesUI []modelUI.Place

	places, err := s.storage.Repositories()[storage.PLACE].SelectAll()

	if err != nil {
		return nil, err
	}

	for i, place := range places {
		placesUI[i] = transportPlaceModel(place)
	}
	return placesUI, nil
}
func (s *PlaceService) Get(id int) (modelUI.Place, error) {
	place, err := s.storage.Repositories()[storage.PLACE].SelectById(id)

	if err != nil {
		return modelUI.Place{}, err
	}

	placeUI := transportPlaceModel(place)

	return placeUI, nil
}
func (s *PlaceService) Add(place modelUI.Place) (int, error) {
	placeDB := PlaceUIToDB(place)

	placeId, err := s.storage.Repositories()[storage.PLACE].Insert(placeDB)
	if err != nil {
		return 0, err
	}
	return placeId, nil
}
func (s *PlaceService) Delete(id int) error {
	return s.storage.Repositories()[storage.PLACE].DeleteById(id)
}
func (s *PlaceService) Update(id int, place modelUI.Place) error {
	placeDB := PlaceUIToDB(place)
	return s.storage.Repositories()[storage.PLACE].UpdateById(id, placeDB)
}

func PlaceUIToDB(placeUI modelUI.Place) modelDB.Place {
	return modelDB.Place{
		Id:          placeUI.Id,
		Name:        placeUI.Name,
		CountTicket: placeUI.CountTicket,
	}
}

func transportPlaceModel(fields storage.Model) modelUI.Place {
	return modelUI.Place{
		Id:          int(fields.([]interface{})[0].(int32)),
		Name:        fields.([]interface{})[1].(string),
		CountTicket: int(fields.([]interface{})[2].(int32)),
	}
}
