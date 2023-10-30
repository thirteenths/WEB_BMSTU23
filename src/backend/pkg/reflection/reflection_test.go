package reflection

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetType(t *testing.T) {
	var tests = []struct {
		input  any
		output string
	}{
		{Student{}, "Student"},
		{Schedule{}, "Schedule"},
		{Public{}, "Public"},
		{Queue{}, "Queue"},
		{Subscription{}, "Subscription"},
	}
	for _, test := range tests {
		require.Equal(t, test.output, GetType(test.input))
	}
}

func TestGetArrayOfFields(t *testing.T) {
	var tests = []struct {
		input  any
		output []string
	}{
		{Student{}, []string{"Id", "Name", "Class", "Role"}},
		{Schedule{}, []string{"Id", "Subject", "Date"}},
		{Public{}, []string{"Id", "Name", "Link"}},
		{Queue{}, []string{"Id", "IdStudent", "IdSchedule", "Status"}},
		{Subscription{}, []string{"Id", "IdStudent", "IdPublic"}},
	}
	for _, test := range tests {
		require.Equal(t, test.output, GetArrayOfFields(test.input))
	}
}

func TestGetValueOfStruct_Student(t *testing.T) {
	models := Student{Name: "Ivan", Class: "IU7-61B", Role: "Stud"}
	var tests = []struct {
		input  string
		output string
	}{
		{"Name", "'Ivan'"},
		{"Class", "'IU7-61B'"},
		{"Role", "'Stud'"},
	}
	for _, test := range tests {
		require.Equal(t, test.output, GetValueOfStruct(models, test.input))
	}
}

func TestGetValueOfStruct_Public(t *testing.T) {
	models := Public{Name: "Stud", Link: "www.go.com"}
	var tests = []struct {
		input  string
		output string
	}{
		{"Name", "'Stud'"},
		{"Link", "'www.go.com'"},
	}
	for _, test := range tests {
		require.Equal(t, test.output, GetValueOfStruct(models, test.input))
	}
}

func TestGetValueOfStruct_Queue(t *testing.T) {
	models := Queue{IdSchedule: 1, IdStudent: 2, Status: false}
	var tests = []struct {
		input  string
		output any
	}{
		{"IdSchedule", "1"},
		{"IdStudent", "2"},
		{"Status", "false"},
	}
	for _, test := range tests {
		require.Equal(t, test.output, GetValueOfStruct(models, test.input))
	}
}

func TestGetValueOfStruct_Schedule(t *testing.T) {
	models := Schedule{Subject: "OOP", Date: "2015-09-10 15:08:55.000000"}
	var tests = []struct {
		input  string
		output any
	}{
		{"Subject", "'OOP'"},
		{"Date", "'2015-09-10 15:08:55.000000'"},
	}
	for _, test := range tests {
		require.Equal(t, test.output, GetValueOfStruct(models, test.input))
	}
}

func TestGetValueOfStruct_Subscription(t *testing.T) {
	models := Subscription{IdStudent: 1, IdPublic: 2}
	var tests = []struct {
		input  string
		output any
	}{
		{"IdStudent", "1"},
		{"IdPublic", "2"},
	}
	for _, test := range tests {
		require.Equal(t, test.output, GetValueOfStruct(models, test.input))
	}
}
