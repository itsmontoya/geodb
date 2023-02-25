package geodb

import (
	"testing"
)

func TestFoot_ToMeters(t *testing.T) {
	tests := []struct {
		name string
		f    Foot
		want Meter
	}{
		{
			name: "basic",
			f:    20 / 0.3048,
			want: 20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.ToMeters(); got != tt.want {
				t.Errorf("Foot.ToMeters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFoot_ToFeet(t *testing.T) {
	tests := []struct {
		name string
		f    Foot
		want Foot
	}{
		{
			name: "basic",
			f:    13,
			want: 13,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.ToFeet(); got != tt.want {
				t.Errorf("Foot.ToFeet() = %v, want %v", got, tt.want)
			}
		})
	}
}
