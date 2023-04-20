package guesser_test

import (
	"fmt"
	"testing"

	"github.com/robherley/guesslang-go/pkg/guesser"
)

func TestIsReliable(t *testing.T) {
	cases := []struct {
		confidences []float64
		expected    bool
	}{
		{[]float64{0.2, 0.1, 0.1, 0.1, 0.1}, false},
		{[]float64{0.9, 0.1, 0.1, 0.1, 0.1}, true},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("confidences: %v", tc.confidences), func(t *testing.T) {
			actual := guesser.IsReliable(tc.confidences)
			if actual != tc.expected {
				t.Errorf("want %t, got %t", tc.expected, actual)
			}
		})
	}
}
