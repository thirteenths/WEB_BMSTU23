package service

import (
	"github.com/thirteenths/WEB_BMSTU23/backend/internal/modelUI"
	"github.com/thirteenths/WEB_BMSTU23/backend/pkg/storage"
)

type IPerson interface {
	GetAll() ([]modelUI.Person, error)
	Get(id int) (modelUI.Person, error)
	Add(person modelUI.Person) (int, error)
	Delete(id int) error
	Update(id int, person modelUI.Person) error
}

type IPoster interface {
	GetAll() ([]modelUI.Poster, error)
	Get(id int) (modelUI.Poster, error)
	Add(poster modelUI.Poster) (int, error)
	Delete(id int) error
	Update(id int, poster modelUI.Poster) error
	CheckIn(idPoster, idPerson int) error
}

type IComic interface {
	GetAll() ([]modelUI.Comic, error)
	Get(id int) (modelUI.Comic, error)
	Add(comic modelUI.Comic) (int, error)
	Delete(id int) error
	Update(id int, comic modelUI.Comic) error
}

type IPlace interface {
	GetAll() ([]modelUI.Place, error)
	Get(id int) (modelUI.Place, error)
	Add(comic modelUI.Place) (int, error)
	Delete(id int) error
	Update(id int, comic modelUI.Place) error
}

type IEvent interface {
	GetAll() ([]modelUI.Event, error)
	Get(id int) (modelUI.Event, error)
	Add(comic modelUI.Event) (int, error)
	Delete(id int) error
	Update(id int, comic modelUI.Event) error
}

type IAuthorization interface {
	ParseToken(accessToken string) (int, int, error)
	GenerationToken(login string, password string) (string, error)
	CreateUser(person modelUI.Person) (int, error)
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
