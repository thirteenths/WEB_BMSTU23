package postgres

import (
	"github.com/thirteenths/WEB_BMSTU23/backend/pkg/storage"
	"log"
)

type Repository struct {
	reqMaker *RequestMaker
	conn     Connect
}

func NewRepository(model any, conn Connect) *Repository {
	return &Repository{
		reqMaker: NewRequestMaker(model),
		conn:     conn}
}

func (r *Repository) Create() error {
	return nil
}

func (r *Repository) Drop() error {
	return nil
}

func (r *Repository) Insert(model any) (int, error) {
	var id int
	query := r.reqMaker.Insert(model)

	log.Println(query, model)

	err := r.conn.db.QueryRow(string(query)).Scan(&id)

	return id, err
}

func (r *Repository) UpdateById(id int, model storage.Model) error {
	query := r.reqMaker.UpdateById(model)

	log.Println(query, id, model)

	_, err := r.conn.db.Exec(string(query), id)

	return err
}

func (r *Repository) DeleteById(id int) error {
	query := r.reqMaker.DeleteById()

	log.Println(query, id)

	_, err := r.conn.db.Exec(string(query), id)

	return err
}

func (r *Repository) SelectById(id int) (storage.Model, error) {
	var model storage.Model

	query := r.reqMaker.SelectById()

	log.Println(query, id)

	row, err := r.conn.db.Query(string(query), id)

	if err != nil {
		return model, err
	}

	row.Next()
	model, err = row.Values()

	return model, err
}

func (r *Repository) SelectByField(fieldName string, fieldValue storage.FieldValue) ([]storage.Model, error) {
	var models []storage.Model

	query := r.reqMaker.SelectByField(fieldName)

	log.Println(query, fieldValue)

	rows, err := r.conn.db.Query(string(query), fieldValue)

	if err != nil {
		return models, err
	}

	for rows.Next() {
		var model storage.Model
		model, err = rows.Values()
		models = append(models, model)
	}

	log.Println("storage: ", models)
	return models, err
}

func (r *Repository) SelectAll() ([]storage.Model, error) {
	var models []storage.Model

	query := r.reqMaker.SelectAll()

	log.Println(query)

	rows, err := r.conn.db.Query(string(query))

	for rows.Next() {
		var model storage.Model
		model, err = rows.Values()
		models = append(models, model)
	}
	return models, err
}
