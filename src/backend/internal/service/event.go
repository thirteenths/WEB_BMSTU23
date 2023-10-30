package service

import (
	"github.com/thirteenths/WEB_BMSTU23/backend/internal/modelDB"
	"github.com/thirteenths/WEB_BMSTU23/backend/internal/modelUI"
	"github.com/thirteenths/WEB_BMSTU23/backend/pkg/storage"
)

type EventService struct {
	storage storage.IStorage
}

func NewEventService(storage storage.IStorage) *EventService {
	return &EventService{storage: storage}
}

func (s *EventService) GetAll() ([]modelUI.Event, error) {
	var eventsUI []modelUI.Event

	events, err := s.storage.Repositories()[storage.EVENT].SelectAll()

	if err != nil {
		return nil, err
	}

	for i, event := range events {
		eventsUI[i] = transportEventModel(event)
	}

	return eventsUI, nil
}
func (s *EventService) Get(id int) (modelUI.Event, error) {
	event, err := s.storage.Repositories()[storage.EVENT].SelectById(id)

	if err != nil {
		return modelUI.Event{}, err
	}

	eventUI := transportEventModel(event)

	return eventUI, nil
}
func (s *EventService) Add(event modelUI.Event) (int, error) {
	var eventDB modelDB.Event

	eventDB = EventUIToDB(event)

	eventId, err := s.storage.Repositories()[storage.EVENT].Insert(eventDB)
	if err != nil {
		return 0, err
	}

	return eventId, nil
}
func (s *EventService) Delete(id int) error {
	return s.storage.Repositories()[storage.EVENT].DeleteById(id)
}
func (s *EventService) Update(id int, event modelUI.Event) error {
	eventDB := EventUIToDB(event)
	return s.storage.Repositories()[storage.EVENT].UpdateById(id, eventDB)
}

func transportEventModel(fields storage.Model) modelUI.Event {
	var event modelUI.Event

	event.Id = int(fields.([]interface{})[0].(int32))
	event.Name = fields.([]interface{})[1].(string)
	event.About = fields.([]interface{})[2].(string)

	return event
}

func EventUIToDB(eventUI modelUI.Event) modelDB.Event {
	return modelDB.Event{
		Id:    eventUI.Id,
		Name:  eventUI.Name,
		About: eventUI.About}
}
