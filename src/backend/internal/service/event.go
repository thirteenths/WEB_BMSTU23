package service

import (
	"github.com/thirteenths/WEB_BMSTU23/backend/internal/model"
	"github.com/thirteenths/WEB_BMSTU23/backend/pkg/storage"
)

type EventService struct {
	storage storage.IStorage
}

func NewEventService(storage storage.IStorage) *EventService {
	return &EventService{storage: storage}
}

func (s *EventService) GetAll() ([]model.Event, error) {
	var eventsUI []model.Event

	events, err := s.storage.Repositories()[storage.EVENT].SelectAll()

	if err != nil {
		return nil, err
	}

	for i, event := range events {
		eventsUI[i] = transportEventModel(event)
	}

	return eventsUI, nil
}
func (s *EventService) Get(id int) (model.Event, error) {
	event, err := s.storage.Repositories()[storage.EVENT].SelectById(id)

	if err != nil {
		return model.Event{}, err
	}

	return transportEventModel(event), nil
}
func (s *EventService) Add(event model.Event) (int, error) {
	return s.storage.Repositories()[storage.EVENT].Insert(event)
}
func (s *EventService) Delete(id int) error {
	return s.storage.Repositories()[storage.EVENT].DeleteById(id)
}
func (s *EventService) Update(id int, event model.Event) error {
	return s.storage.Repositories()[storage.EVENT].UpdateById(id, event)
}

func transportEventModel(fields storage.Model) model.Event {
	return model.Event{
		Id:    int(fields.([]interface{})[0].(int32)),
		Name:  fields.([]interface{})[1].(string),
		About: fields.([]interface{})[2].(string),
	}
}
