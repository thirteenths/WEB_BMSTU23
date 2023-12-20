package service

import (
	"github.com/thirteenths/WEB_BMSTU23/backend/internal/model"
	"github.com/thirteenths/WEB_BMSTU23/backend/pkg/storage"
)

type ComicService struct {
	storage storage.IStorage
}

func NewComicService(storage storage.IStorage) *ComicService {
	return &ComicService{storage: storage}
}

func (cs *ComicService) GetAll() ([]model.Comic, error) {
	var comicUI []model.Comic

	comics, err := cs.storage.Repositories()[storage.COMIC].SelectAll()
	if err != nil {
		return nil, err
	}

	for i, comic := range comics {
		comicUI[i] = transportComicModel(comic)
	}

	return comicUI, nil
}

func (cs *ComicService) Get(id int) (model.Comic, error) {
	comic, err := cs.storage.Repositories()[storage.COMIC].SelectById(id)
	if err != nil {
		return model.Comic{}, err
	}

	return transportComicModel(comic), nil
}

func (cs *ComicService) Add(comic model.Comic) (int, error) {
	return cs.storage.Repositories()[storage.COMIC].Insert(comic)
}

func (cs *ComicService) Delete(id int) error {
	return cs.storage.Repositories()[storage.COMIC].DeleteById(id)
}

func (cs *ComicService) Update(id int, comic model.Comic) error {
	return cs.storage.Repositories()[storage.COMIC].UpdateById(id, comic)
}

func transportComicModel(fields storage.Model) model.Comic {
	return model.Comic{
		Id:       int(fields.([]interface{})[0].(int32)),
		Image:    fields.([]interface{})[1].(string),
		Name:     fields.([]interface{})[2].(string),
		City:     fields.([]interface{})[3].(string),
		Sentence: fields.([]interface{})[4].(string),
		CountKek: int(fields.([]interface{})[5].(int32)),
	}
}
