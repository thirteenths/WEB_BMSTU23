package postgres

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
	"time"
)

type Student struct {
	Id    int    `json:"_"`
	Name  string `json:"name"`
	Class string `json:"class"` // В скрипте grope но я так понимаю это очепятка
	Role  string `json:"role"`  // = "student"
}

type TestModel struct {
	IntField    int       `json:"id"`
	StringField string    `json:"string_field"`
	BoolField   bool      `json:"bool_field"`
	TimeField   time.Time `json:"time_field"`
}

func NewRequestMakerStudentMock() *RequestMaker {
	return &RequestMaker{
		tableName:   "Student",
		tableFields: []string{"Id", "Name", "Class", "Role"},
	}
}

func TestRequestMaker_InsertStudent(t *testing.T) {
	requestMaker := NewRequestMakerStudentMock()

	var tests = []struct {
		input  any
		output string
	}{
		{Student{Name: "Sam", Class: "IU7-51B", Role: "student"}, "INSERT INTO STUDENT (NAME,CLASS,ROLE) VALUES ('Sam','IU7-51B','student') RETURNING ID"},
		{Student{Name: "Sam", Class: "IU7-51B"}, "INSERT INTO STUDENT (NAME,CLASS) VALUES ('Sam','IU7-51B') RETURNING ID"},
		{Student{Class: "IU7-51B"}, "INSERT INTO STUDENT (CLASS) VALUES ('IU7-51B') RETURNING ID"},
	}
	for _, test := range tests {
		require.Equal(t, test.output, requestMaker.Insert(test.input))
	}
}

func TestRequestMaker_UpdateByIdStudent(t *testing.T) {
	requestMaker := NewRequestMakerStudentMock()

	var tests = []struct {
		input  any
		output string
	}{
		{Student{Name: "Sam", Class: "IU7-51B", Role: "student"}, "UPDATE STUDENT SET NAME='Sam',CLASS='IU7-51B',ROLE='student' WHERE ID=$1"},
		{Student{Name: "Sam", Class: "IU7-51B"}, "UPDATE STUDENT SET NAME='Sam',CLASS='IU7-51B' WHERE ID=$1"},
		{Student{Name: "Sam"}, "UPDATE STUDENT SET NAME='Sam' WHERE ID=$1"},
		{Student{Class: "IU7-51B"}, "UPDATE STUDENT SET CLASS='IU7-51B' WHERE ID=$1"},
		{Student{Name: "Sam", Role: "student"}, "UPDATE STUDENT SET NAME='Sam',ROLE='student' WHERE ID=$1"},
	}
	for _, test := range tests {
		require.Equal(t, test.output, requestMaker.UpdateById(test.input))
	}
}

func TestRequestMaker_DeleteByIdStudent(t *testing.T) {
	requestMaker := NewRequestMakerStudentMock()

	var tests = []struct {
		input  any
		output string
	}{
		{Student{Name: "Sam", Class: "IU7-51B", Role: "student"}, "DELETE FROM STUDENT WHERE ID=$1"},
		{Student{Name: "Sam", Class: "IU7-51B"}, "DELETE FROM STUDENT WHERE ID=$1"},
		{Student{Name: "Sam"}, "DELETE FROM STUDENT WHERE ID=$1"},
		{Student{Class: "IU7-51B"}, "DELETE FROM STUDENT WHERE ID=$1"},
		{Student{Name: "Sam", Role: "student"}, "DELETE FROM STUDENT WHERE ID=$1"},
	}
	for _, test := range tests {
		require.Equal(t, test.output, requestMaker.DeleteById())
	}
}

func TestRequestMaker_SelectByIdStudent(t *testing.T) {
	requestMaker := NewRequestMakerStudentMock()

	var tests = []struct {
		input  any
		output string
	}{
		{Student{Name: "Sam", Class: "IU7-51B", Role: "student"}, "SELECT * FROM STUDENT WHERE ID=$1"},
		{Student{Name: "Sam", Class: "IU7-51B"}, "SELECT * FROM STUDENT WHERE ID=$1"},
		{Student{Name: "Sam"}, "SELECT * FROM STUDENT WHERE ID=$1"},
		{Student{Class: "IU7-51B"}, "SELECT * FROM STUDENT WHERE ID=$1"},
		{Student{Name: "Sam", Role: "student"}, "SELECT * FROM STUDENT WHERE ID=$1"},
	}
	for _, test := range tests {
		require.Equal(t, test.output, requestMaker.SelectById())
	}
}

func TestRequestMaker_SelectAll(t *testing.T) {
	requestMaker := NewRequestMakerStudentMock()

	var tests = []struct {
		input  any
		output string
	}{
		{Student{Name: "Sam", Class: "IU7-51B", Role: "student"}, "SELECT * FROM STUDENT"},
		{Student{Name: "Sam", Class: "IU7-51B"}, "SELECT * FROM STUDENT"},
		{Student{Name: "Sam"}, "SELECT * FROM STUDENT"},
		{Student{Class: "IU7-51B"}, "SELECT * FROM STUDENT"},
		{Student{Name: "Sam", Role: "student"}, "SELECT * FROM STUDENT"},
	}
	for _, test := range tests {
		require.Equal(t, test.output, requestMaker.SelectAll())
	}
}

func TestRequestMaker_SelectByField(t *testing.T) {
	requestMaker := NewRequestMakerStudentMock()

	var tests = []struct {
		inputField string
		output     string
	}{
		{"Id", "SELECT * FROM STUDENT WHERE ID=$1"},
		{"Name", "SELECT * FROM STUDENT WHERE NAME=$1"},
		{"Class", "SELECT * FROM STUDENT WHERE CLASS=$1"},
		{"Role", "SELECT * FROM STUDENT WHERE ROLE=$1"},
	}
	for _, test := range tests {
		require.Equal(t, test.output, requestMaker.SelectByField(strings.ToUpper(test.inputField)))
	}
}
