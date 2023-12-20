package service

import (
	"github.com/thirteenths/WEB_BMSTU23/backend/internal/model"
	"github.com/thirteenths/WEB_BMSTU23/backend/pkg/storage"
)

type PosterService struct {
	storage storage.IStorage
}

func NewPosterService(storage storage.IStorage) *PosterService {
	return &PosterService{storage: storage}
}

func (ps *PosterService) GetAll() ([]model.Poster, error) {
	var postersUI []model.Poster

	posters, err := ps.storage.Repositories()[storage.POSTER].SelectAll()
	if err != nil {
		return nil, err
	}

	for i, poster := range posters {
		postersUI[i] = transportPosterModel(poster)
	}

	return postersUI, nil
}

func (ps *PosterService) Get(id int) (model.Poster, error) {
	poster, err := ps.storage.Repositories()[storage.POSTER].SelectById(id)
	if err != nil {
		return model.Poster{}, err
	}
	return transportPosterModel(poster), nil
}

func (ps *PosterService) Add(poster model.Poster) (int, error) {
	return ps.storage.Repositories()[storage.POSTER].Insert(poster)
}

func (ps *PosterService) Delete(id int) error {
	return ps.storage.Repositories()[storage.PLACE].DeleteById(id)
}

func (ps *PosterService) Update(id int, poster model.Poster) error {
	return ps.storage.Repositories()[storage.POSTER].UpdateById(id, poster)
}

func (ps *PosterService) CheckIn(idPoster, idPerson int) error {
	_, err := ps.storage.Repositories()[storage.TICKET].Insert(model.Ticket{IdPerson: idPerson, IdPoster: idPoster})
	if err != nil {
		return err
	}
	return nil
}

func transportPosterModel(fields storage.Model) model.Poster {
	return model.Poster{
		Id:      int(fields.([]interface{})[0].(int32)),
		Image:   fields.([]interface{})[1].(string),
		IdEvent: int(fields.([]interface{})[2].(int32)),
		IdPlace: int(fields.([]interface{})[3].(int32)),
		Date:    fields.([]interface{})[4].(string),
	}
}
