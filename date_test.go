package date_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pieterclaerhout/go-date"
)

func TestUnixRoundToHour(t *testing.T) {

	type test struct {
		name     string
		input    int64
		expected int64
	}

	var tests = []test{
		{"0", 0, 0},
		{"1553862120", 1553862120, 1553860800},
		{"1553862150", 1553862150, 1553860800},
		{"1553862179", 1553862179, 1553860800},
		{"1553862181", 1553862181, 1553860800},
		{"1553860799", 1553860799, 1553857200},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := date.UnixRoundToHour(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}

}
