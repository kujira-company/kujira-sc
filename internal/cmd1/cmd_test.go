package cmd1_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kujira-company/kujira-sc/internal/cmd1"
)

func TestCommand1(t *testing.T) {
	tests := []struct {
		name    string
		in      int
		wantErr error
	}{
		{
			"case zero",
			0,
			cmd1.ErrIsZero,
		},
		{
			"case odd",
			1,
			nil,
		},
		{
			"case even",
			2,
			nil,
		},
		{
			"case nabeatsu",
			3,
			nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := cmd1.Command1(tt.in)
			if !cmp.Equal(err, tt.wantErr, cmpopts.EquateErrors()) {
				t.Fatalf(cmp.Diff(err, tt.wantErr, cmpopts.EquateErrors()))
			}
		})
	}
}

func TestCommandHogeFuga(t *testing.T) {
	tests := []struct {
		name string
		in   int
		got  string
	}{
		{
			"case zero",
			0,
			"hoge",
		},
		{
			"case odd",
			1,
			"fuga",
		},
		{
			"case even",
			2,
			"hoge",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := cmd1.CommandHogeFuga(tt.in)
			if got != tt.got {
				t.Fatalf(cmp.Diff(got, tt.got))
			}
		})
	}
}
