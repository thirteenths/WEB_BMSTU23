package postgres

import (
	"fmt"
	"log"
)

type DifQManager struct {
	conn Connect
}

func NewDifQManager(conn Connect) *DifQManager {
	return &DifQManager{conn: conn}
}

func (m *DifQManager) InsertDataViewer(idPerson, idSchedule int) error {
	query := fmt.Sprintf("CALL insert_data_viewer( %d, %d);", idPerson, idSchedule)

	log.Print(query)

	_, err := m.conn.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (m *DifQManager) DeleteDataViewer(idPerson, idSchedule int) error {
	query := fmt.Sprintf("CALL delete_data_viewer( %d, %d);", idPerson, idSchedule)

	log.Print(query)

	_, err := m.conn.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (m *DifQManager) InsertDataComic(idPerson, idSchedule int) error {
	query := fmt.Sprintf("CALL insert_data_comic( %d, %d);", idPerson, idSchedule)

	log.Print(query)

	_, err := m.conn.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (m *DifQManager) DeleteDataComic(idPerson, idSchedule int) error {
	query := fmt.Sprintf("CALL delete_data_comic( %d, %d);", idPerson, idSchedule)

	log.Print(query)

	_, err := m.conn.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
