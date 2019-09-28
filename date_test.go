package yddate_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pieterclaerhout/go-yddate"
)

func TestUnixRoundToHourTest(t *testing.T) {

	type unixRoundToHourTest struct {
		input    int64
		expected int64
	}

	var unixRoundToHourTests = []unixRoundToHourTest{
		{0, 0},
		{1553862120, 1553860800},
		{1553862150, 1553860800},
		{1553862179, 1553860800},
		{1553862181, 1553860800},
		{1553860799, 1553857200},
	}

	for _, test := range unixRoundToHourTests {
		actual := yddate.UnixRoundToHour(test.input)
		assert.Equal(t, test.expected, actual)
	}

}
