package convertdata

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestStringToDate(t *testing.T) {
	var tests = []struct {
		input  string
		output time.Time
	}{
		{"2015-09-10 15:08:55", time.Date(2015, time.September, 10, 15, 8, 55, 0, time.UTC)},
		{"2015-09-10", time.Date(2015, time.September, 10, 0, 8, 0, 0, time.UTC)},
	}
	for _, test := range tests {
		require.Equal(t, test.output, StringToDate(test.input))
	}
}

func TestCreateDate(t *testing.T) {
	/*var tests = []struct {
		input struct {
			year  int
			month int
			day   int
			hour  int
			min   int
		}
		output time.Time
	}{
		//{int{2015, 9, 10, 15, 8, 55}, time.Date(2015, time.September, 10, 15, 8, 55, 0, time.UTC)},
	}
	for _, test := range tests {
		require.Equal(t, test.output, CreateDate(test.input))
	}*/
}
