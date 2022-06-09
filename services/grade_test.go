package services_test

import (
	"basic/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckGrade(t *testing.T) {
	type testCases struct {
		name     string
		score    int
		expected string
	}
	cases := []testCases{
		{name: "success grade A", score: 80, expected: "A"},
		{name: "success grade B", score: 70, expected: "B"},
		{name: "success grade C", score: 60, expected: "C"},
		{name: "success grade D", score: 50, expected: "D"},
		{name: "success grade F", score: 30, expected: "F"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := services.CheckGrade(tc.score)
			assert.Equal(t, tc.expected, actual)
		})
	}

}

func BenchmarkCheckGrade(b *testing.B) {
	for i := 0; i < b.N; i++ {
		services.CheckGrade(80)
	}
}

func TestHello(t *testing.T) {

}
