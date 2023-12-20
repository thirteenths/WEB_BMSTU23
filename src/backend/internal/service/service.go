package service

import (
	"github.com/thirteenths/WEB_BMSTU23/backend/internal/model"
	"github.com/thirteenths/WEB_BMSTU23/backend/pkg/storage"
)

type IPerson interface {
	GetAll() ([]model.Person, error)
	Get(id int) (model.Person, error)
	Add(person model.Person) (int, error)
	Delete(id int) error
	Update(id int, person model.Person) error
}

type IPoster interface {
	GetAll() ([]model.Poster, error)
	Get(id int) (model.Poster, error)
	Add(poster model.Poster) (int, error)
	Delete(id int) error
	Update(id int, poster model.Poster) error
	CheckIn(idPoster, idPerson int) error
}

type IComic interface {
	GetAll() ([]model.Comic, error)
	Get(id int) (model.Comic, error)
	Add(comic model.Comic) (int, error)
	Delete(id int) error
	Update(id int, comic model.Comic) error
}

type IPlace interface {
	GetAll() ([]model.Place, error)
	Get(id int) (model.Place, error)
	Add(comic model.Place) (int, error)
	Delete(id int) error
	Update(id int, comic model.Place) error
}

type IEvent interface {
	GetAll() ([]model.Event, error)
	Get(id int) (model.Event, error)
	Add(comic model.Event) (int, error)
	Delete(id int) error
	Update(id int, comic model.Event) error
}

type IAuthorization interface {
	ParseToken(accessToken string) (int, int, error)
	GenerationToken(login string, password string) (string, error)
	CreateUser(person model.Person) (int, error)
}

type Service struct {
	IComic
	IPoster
	IPlace
	IEvent
	IAuthorization
}

func NewService(storage storage.IStorage) *Service {
	return &Service{
		IComic:         NewComicService(storage),
		IPoster:        NewPosterService(storage),
		IPlace:         NewPlaceService(storage),
		IEvent:         NewEventService(storage),
		IAuthorization: NewAuthService(storage),
	}
}
