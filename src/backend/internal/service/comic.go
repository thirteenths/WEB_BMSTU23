package service

import (
	"github.com/thirteenths/WEB_BMSTU23/backend/internal/modelDB"
	"github.com/thirteenths/WEB_BMSTU23/backend/internal/modelUI"
	"github.com/thirteenths/WEB_BMSTU23/backend/pkg/storage"
)

type ComicService struct {
	storage storage.IStorage
}

func NewComicService(storage storage.IStorage) *ComicService {
	return &ComicService{storage: storage}
}

func (cs *ComicService) GetAll() ([]modelUI.Comic, error) {
	var comicUI []modelUI.Comic

	comics, err := cs.storage.Repositories()[storage.COMIC].SelectAll()
	if err != nil {
		return nil, err
	}

	for i, comic := range comics {
		comicUI[i].Id = int(comic.([]interface{})[0].(int32))

		idImage := int(comic.([]interface{})[1].(int32))
		output, err := cs.storage.Repositories()[storage.IMAGE].SelectById(idImage)
		if err != nil {
			return nil, err
		}

		comicUI[i].Image = output.([]interface{})[1].(string)
		comicUI[i].Name = comic.([]interface{})[2].(string)
		comicUI[i].City = comic.([]interface{})[3].(string)
		comicUI[i].Sentence = comic.([]interface{})[4].(string)
		comicUI[i].CountOfKek = int(comic.([]interface{})[5].(int32))
	}

	return comicUI, nil
}

func (cs *ComicService) Get(id int) (modelUI.Comic, error) {
	var comicUI modelUI.Comic

	comic, err := cs.storage.Repositories()[storage.COMIC].SelectById(id)
	if err != nil {
		return modelUI.Comic{}, err
	}

	comicUI.Id = int(comic.([]interface{})[0].(int32))

	idImage := int(comic.([]interface{})[1].(int32))
	output, err := cs.storage.Repositories()[storage.IMAGE].SelectById(idImage)
	if err != nil {
		return modelUI.Comic{}, err
	}

	comicUI.Image = output.([]interface{})[1].(string)
	comicUI.Name = comic.([]interface{})[2].(string)
	comicUI.City = comic.([]interface{})[3].(string)
	comicUI.Sentence = comic.([]interface{})[4].(string)
	comicUI.CountOfKek = int(comic.([]interface{})[5].(int32))

	return comicUI, nil
}

func (cs *ComicService) Add(comic modelUI.Comic) (int, error) {
	var comicDB modelDB.Comic

	comicDB.Name = comic.Name
	comicDB.City = comic.City
	comicDB.Sentence = comic.Sentence
	comicDB.CountKek = comic.CountOfKek

	imageId, err := cs.storage.Repositories()[storage.IMAGE].Insert(modelDB.Image{Path: comic.Image})
	if err != nil {
		return 0, err
	}

	comicDB.IdImage = imageId

	comicId, err := cs.storage.Repositories()[storage.COMIC].Insert(comicDB)
	if err != nil {
		return 0, err
	}

	return comicId, nil
}

func (cs *ComicService) Delete(id int) error {
	err := cs.storage.Repositories()[storage.COMIC].DeleteById(id)
	if err != nil {
		return err
	}
	return nil
}

func (cs *ComicService) Update(id int, comic modelUI.Comic) error {
	var comicDB modelDB.Comic

	comicDB.Name = comic.Name
	comicDB.City = comic.City
	comicDB.Sentence = comic.Sentence
	comicDB.CountKek = comic.CountOfKek

	image, err := cs.storage.Repositories()[storage.IMAGE].SelectByField("path", comic.Image)
	if err != nil {
		return err
	}

	comicDB.IdImage = int(image[0].([]interface{})[0].(int32))

	err = cs.storage.Repositories()[storage.COMIC].UpdateById(id, comicDB)
	if err != nil {
		return err
	}
	return nil
}

func transport(fields storage.Model) modelDB.Comic {
	var comic modelDB.Comic

	comic.Id = int(fields.([]interface{})[0].(int32))
	comic.IdImage = int(fields.([]interface{})[1].(int32))
	comic.Name = fields.([]interface{})[2].(string)
	comic.City = fields.([]interface{})[3].(string)
	comic.Sentence = fields.([]interface{})[4].(string)
	comic.CountKek = int(fields.([]interface{})[5].(int32))

	return comic
}
