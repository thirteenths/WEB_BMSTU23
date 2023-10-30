package service

import (
	"github.com/thirteenths/WEB_BMSTU23/backend/internal/modelDB"
	"github.com/thirteenths/WEB_BMSTU23/backend/internal/modelUI"
	"github.com/thirteenths/WEB_BMSTU23/backend/pkg/storage"
)

type PosterService struct {
	storage storage.IStorage
}

func NewPosterService(storage storage.IStorage) *PosterService {
	return &PosterService{storage: storage}
}

func (ps *PosterService) GetAll() ([]modelUI.Poster, error) {
	var postersUI []modelUI.Poster

	posters, err := ps.storage.Repositories()[storage.POSTER].SelectAll()
	if err != nil {
		return nil, err
	}

	for i, poster := range posters {
		imageId := int(poster.([]interface{})[1].(int32))
		image, err := ps.storage.Repositories()[storage.IMAGE].SelectById(imageId)
		if err != nil {
			return nil, err
		}
		postersUI[i] = transportPosterModel(posters, image.([]interface{})[1].(string))
	}

	return postersUI, nil
}

func (ps *PosterService) Get(id int) (modelUI.Poster, error) {
	var posterUI modelUI.Poster

	poster, err := ps.storage.Repositories()[storage.POSTER].SelectById(id)
	if err != nil {
		return modelUI.Poster{}, err
	}

	posterUI.Id = int(poster.([]interface{})[0].(int32))

	idImage := int(poster.([]interface{})[1].(int32))
	posterUI.IdEvent = int(poster.([]interface{})[2].(int32))
	posterUI.IdPlace = int(poster.([]interface{})[3].(int32))

	posterUI.Date = poster.([]interface{})[0].(string)

	image, err := ps.storage.Repositories()[storage.IMAGE].SelectById(idImage)
	if err != nil {
		return modelUI.Poster{}, err
	}
	posterUI.Image = image.([]interface{})[1].(string)

	/*event, err := ps.storage.Repositories()[storage.EVENT].SelectById(idEvent)
	if err != nil {
		return modelUI.Poster{}, err
	}
	posterUI.Place = event.([]interface{})[1].(string)
	posterUI.About = event.([]interface{})[2].(string)

	place, err := ps.storage.Repositories()[storage.PLACE].SelectById(idPlace)
	if err != nil {
		return modelUI.Poster{}, err
	}
	posterUI.Place = place.([]interface{})[1].(string)*/

	return posterUI, nil
}

func (ps *PosterService) Add(poster modelUI.Poster) (int, error) {
	IdImage, err := ps.storage.Repositories()[storage.IMAGE].Insert(modelDB.Image{Path: poster.Image})
	if err != nil {
		return 0, err
	}

	posterDB := PosterUIToDB(poster, IdImage)

	posterId, err := ps.storage.Repositories()[storage.POSTER].Insert(posterDB)
	if err != nil {
		return 0, err
	}

	return posterId, nil
}

func (ps *PosterService) Delete(id int) error {
	return ps.storage.Repositories()[storage.PLACE].DeleteById(id)
}

func (ps *PosterService) Update(id int, poster modelUI.Poster) error {
	IdImage, err := ps.storage.Repositories()[storage.IMAGE].Insert(modelDB.Image{Path: poster.Image})
	if err != nil {
		return err
	}

	posterDB := PosterUIToDB(poster, IdImage)

	err = ps.storage.Repositories()[storage.POSTER].UpdateById(id, posterDB)
	if err != nil {
		return err
	}
	return nil
}

func (ps *PosterService) CheckIn(idPoster, idPerson int) error {
	_, err := ps.storage.Repositories()[storage.TICKET].Insert(modelDB.Ticket{IdPerson: idPerson, IdPoster: idPoster})
	if err != nil {
		return err
	}
	return nil
}

func transportPosterModel(fields storage.Model, image string) modelUI.Poster {
	return modelUI.Poster{
		Id:      int(fields.([]interface{})[0].(int32)),
		Image:   image,
		IdEvent: int(fields.([]interface{})[2].(int32)),
		IdPlace: int(fields.([]interface{})[3].(int32)),
		Date:    fields.([]interface{})[4].(string),
	}
}

func PosterUIToDB(posterUI modelUI.Poster, imageId int) modelDB.Poster {
	return modelDB.Poster{
		Id:      posterUI.Id,
		IdImage: imageId,
		IdPlace: posterUI.IdPlace,
		IdEvent: posterUI.IdEvent,
		Date:    posterUI.Date}
}
