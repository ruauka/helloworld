package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	TestTable := []struct {
		num1, num2 int
		expected   int
		testName   string
	}{
		{
			num1:     10,
			num2:     20,
			expected: 301,
			testName: "test ok",
		},
	}

	for _, testCase := range TestTable {
		t.Run(testCase.testName, func(t *testing.T) {
			res := add(testCase.num1, testCase.num2)
			require.Equal(t, testCase.expected, res)
		})
	}
}
