package brdocs

import (
	"testing"
)

func TestCNH(t *testing.T) {
	t.Run("Format", func(t *testing.T) {
		cnh := CNH("53197826392")

		if string(cnh) != cnh.Format() {
			t.Errorf("Invalid CNH format")
		}
	})

	t.Run("Check if the CNH number is valid", func(t *testing.T) {
		cnh := CNH("53197826392")

		if !cnh.IsValid() {
			t.Errorf("Invalid CNH number")
		}
	})
}
