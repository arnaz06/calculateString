package calculateString

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculateString(t *testing.T) {
	tests := []struct {
		testName       string
		input          string
		expectedResult int
		expectedError  string
	}{
		{
			testName:       "valid string",
			input:          "5 + 2 - 1",
			expectedResult: 5 + 2 - 1,
		},
		{
			testName:       "valid string",
			input:          "512 - 20 + 100",
			expectedResult: 512 - 20 + 100,
		},
		{
			testName:       "valid string",
			input:          "100 - 81 + 18",
			expectedResult: 100 - 81 + 18,
		},
		{
			testName:       "invalid string",
			input:          "uasda5 + 2 - 1",
			expectedResult: 0,
			expectedError:  "validation error",
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			res, err := Calculate(test.input)
			require.Equal(t, test.expectedResult, res)
			require.Equal(t, test.expectedError, err)
		})
	}
}
