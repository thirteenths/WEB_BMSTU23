package postgres

import (
	"github.com/thirteenths/WEB_BMSTU23/backend/internal/modelDB"
	"github.com/thirteenths/WEB_BMSTU23/backend/pkg/storage"
)

type Storage struct {
	Reps     map[storage.TableName]storage.Repository
	QManager storage.DifQManager
}

func NewPostgresStorage(user, password, database string) (*Storage, error) {

	conn, err := NewConnect(user, password, database)
	if err != nil {
		return &Storage{}, err
	}

	Reps := make(map[storage.TableName]storage.Repository)

	Reps[storage.TICKET] = NewRepository(modelDB.Ticket{}, *conn)
	Reps[storage.PERSON] = NewRepository(modelDB.Person{}, *conn)
	Reps[storage.EVENT] = NewRepository(modelDB.Event{}, *conn)
	Reps[storage.COMIC] = NewRepository(modelDB.Comic{}, *conn)
	Reps[storage.PLACE] = NewRepository(modelDB.Place{}, *conn)
	Reps[storage.IMAGE] = NewRepository(modelDB.Image{}, *conn)
	Reps[storage.POSTER] = NewRepository(modelDB.Poster{}, *conn)

	return &Storage{Reps: Reps, QManager: NewDifQManager(*conn)}, err
}

func (s *Storage) QueryManager() storage.DifQManager {
	return s.QManager
}

func (s *Storage) Repositories() map[storage.TableName]storage.Repository {
	return s.Reps
}
