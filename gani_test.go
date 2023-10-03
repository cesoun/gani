package gani

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed test/default.gani
var defaultGani []byte

//go:embed test/example.gani
var exampleGani []byte

//go:embed test/example-singledir.gani
var exampleSingleDirGani []byte

func TestGani_Parse(t *testing.T) {
	type args struct {
		bs []byte
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Parse Default Gani",
			args: args{
				bs: defaultGani,
			},
			wantErr: false,
		},
		{
			name: "Parse Example Gani",
			args: args{
				bs: exampleGani,
			},
			wantErr: false,
		},
		{
			name: "Parse ExampleSingleDir Gani",
			args: args{
				bs: exampleSingleDirGani,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGani()
			if err := g.Parse(tt.args.bs); (err != nil) != tt.wantErr {
				t.Errorf("Parse() with error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Ignore windows formatting here.
			betterWant := strings.ReplaceAll(string(tt.args.bs), "\r\n", "\n")
			if !strings.EqualFold(betterWant, g.String()) {
				t.Errorf("String() mismatch:\n%q\n%q\n", betterWant, g.String())
			}
		})
	}
}
