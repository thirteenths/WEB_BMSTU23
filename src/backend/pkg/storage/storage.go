package storage

type Query string // Запрос
type Model any
type FieldValue any // Запрос

type Repository interface {
	Create() error // Создаsть таблицу (?)
	Drop() error   // Удалить таблицу (?)

	Insert(model any) (int, error)                                          // Добавить строчку
	UpdateById(id int, model Model) error                                   // Изменить строчку
	DeleteById(id int) error                                                // Удалить строчку
	SelectById(id int) (Model, error)                                       // Найти строчку по id
	SelectByField(fieldName string, fieldValue FieldValue) ([]Model, error) // Найти строчку по полю
	SelectAll() ([]Model, error)
}

type RequestMaker interface {
	Create() Query // Создать таблицу (?)
	Drop() Query   // Удалить таблицу (?)

	Insert(model any) Query                               // Добавить строчку
	UpdateById(id int) Query                              // Изменить строчку
	DeleteById(id int) Query                              // Удалить строчку
	SelectById(id int) Query                              // Найти строчку по id
	SelectByField(fieldName string, fieldValue any) Query // Найти строчку по полю
}

type DifQManager interface {
	InsertDataViewer(idPerson, idSchedule int) error
	DeleteDataViewer(idPerson, idSchedule int) error
	InsertDataComic(idPerson, idSchedule int) error
	DeleteDataComic(idPerson, idSchedule int) error
}

type TableName string

const (
	COMIC  TableName = "comic"
	PERSON TableName = "person"
	EVENT  TableName = "event"
	PLACE  TableName = "place"
	IMAGE  TableName = "image"
	POSTER TableName = "poster"
	TICKET TableName = "ticket"
)

type IStorage interface {
	Repositories() map[TableName]Repository
	QueryManager() DifQManager
}
