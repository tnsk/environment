package environment

import (
	"testing"
)

func TestChecks(t *testing.T) {
	// Valid checks "true test"
	t.Run("Check Valid", func(t *testing.T) {
		if err := Checks([]string{"GOPATH", "SHELL"}); err != nil {
			t.Fatal(err)
		}
	})

	// Invalid checks "false test"
	t.Run("Check Invalid", func(t *testing.T) {
		if err := Checks([]string{"GOPATH", "SHELL", "FISTIKCISAHAP"}); err == nil {
			t.Fatal("Expected error")
		}
	})
}
