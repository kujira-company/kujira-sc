package cmd2_test

import (
	"testing"

	"github.com/kujira-company/kujira-sc/internal/cmd2"
)

func TestCommand2(t *testing.T) {
	tests := []struct {
		name string
		in   int
	}{
		{
			"case odd",
			1,
		},
		{
			"case even",
			2,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := cmd2.Command2(tt.in)
			if err != nil {
				t.Fatalf("error: %+v", err)
			}
		})
	}
}
