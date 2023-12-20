package postgres

import (
	"github.com/thirteenths/WEB_BMSTU23/backend/internal/model"
	"github.com/thirteenths/WEB_BMSTU23/backend/pkg/storage"
)

type Storage struct {
	Reps     map[storage.TableName]storage.Repository
	QManager storage.DifQManager
}

func NewPostgresStorage(url string) (*Storage, error) {

	conn, err := NewConnect(url)
	if err != nil {
		return &Storage{}, err
	}

	Reps := make(map[storage.TableName]storage.Repository)

	Reps[storage.TICKET] = NewRepository(model.Ticket{}, *conn)
	Reps[storage.PERSON] = NewRepository(model.Person{}, *conn)
	Reps[storage.EVENT] = NewRepository(model.Event{}, *conn)
	Reps[storage.COMIC] = NewRepository(model.Comic{}, *conn)
	Reps[storage.PLACE] = NewRepository(model.Place{}, *conn)
	Reps[storage.IMAGE] = NewRepository(model.Image{}, *conn)
	Reps[storage.POSTER] = NewRepository(model.Poster{}, *conn)

	return &Storage{Reps: Reps, QManager: NewDifQManager(*conn)}, err
}

func (s *Storage) QueryManager() storage.DifQManager {
	return s.QManager
}

func (s *Storage) Repositories() map[storage.TableName]storage.Repository {
	return s.Reps
}
