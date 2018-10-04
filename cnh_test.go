package brdocs

import (
	"testing"
)

func TestCNH(t *testing.T) {
	number := "53197826392"

	t.Run("Format", func(t *testing.T) {
		cnh := NewCNH(number)

		if number != cnh.Format() {
			t.Errorf("Invalid CNH format")
		}
	})

	t.Run("Check if the CNH number is valid", func(t *testing.T) {
		cnh := NewCNH(number)

		if !cnh.IsValid() {
			t.Errorf("Invalid CNH number")
		}
	})
}
