package postgres

import (
	"fmt"
	"github.com/thirteenths/WEB_BMSTU23/backend/pkg/reflection"
	"github.com/thirteenths/WEB_BMSTU23/backend/pkg/strcase"
)

type RequestMaker struct {
	tableName   string
	tableFields []string
}

func NewRequestMaker(model any) *RequestMaker {
	return &RequestMaker{
		tableName:   reflection.GetType(model),
		tableFields: reflection.GetArrayOfFields(model)}
}

func (m *RequestMaker) Create() string {
	return ""
}

func (m *RequestMaker) Drop() string {
	return ""
}

func (m *RequestMaker) Insert(model any) string {
	query := fmt.Sprintf("INSERT INTO %s (", strcase.UpperSnakeCase(m.tableName))

	var queryFields, queryValues string
	var flag bool

	for i := 1; i < len(m.tableFields); i++ {
		data := reflection.GetValueOfStruct(model, m.tableFields[i])
		if data != "''" {
			if flag {
				queryFields += ","
				queryValues += ","
			}
			flag = true
			queryFields += m.tableFields[i]
			queryValues += fmt.Sprintf("%s", data)
		}

	}

	query += strcase.UpperSnakeCase(queryFields) + ") VALUES (" + queryValues + ") RETURNING ID"

	return query
}

func (m *RequestMaker) UpdateById(model any) string {
	query := "UPDATE " + strcase.UpperSnakeCase(m.tableName) + " SET "
	var flag bool

	for i := 1; i < len(m.tableFields); i++ {
		data := reflection.GetValueOfStruct(model, m.tableFields[i])

		if data != "''" {

			if flag {
				query += ","
			}
			flag = true
			data := reflection.GetValueOfStruct(model, m.tableFields[i])
			query += fmt.Sprintf("%s=%s", strcase.UpperSnakeCase(m.tableFields[i]), data)
		}
	}

	query += " WHERE ID=$1"
	return query
}

func (m *RequestMaker) DeleteById() string {
	query := "DELETE FROM " + strcase.UpperSnakeCase(m.tableName) + " WHERE ID=$1"

	return query
}

func (m *RequestMaker) SelectById() string {
	query := "SELECT * FROM " + strcase.UpperSnakeCase(m.tableName) + " WHERE ID=$1"

	return query
}

func (m *RequestMaker) SelectByField(fieldName string) string {
	query := "SELECT * FROM " + strcase.UpperSnakeCase(m.tableName) + " WHERE " + fieldName + "=$1"

	return query
}

func (m *RequestMaker) SelectAll() string {
	query := "SELECT * FROM " + strcase.UpperSnakeCase(m.tableName)

	return query
}
