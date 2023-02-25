package geodb

import (
	"testing"
)

func TestMeter_ToMeters(t *testing.T) {
	tests := []struct {
		name string
		m    Meter
		want Meter
	}{
		{
			name: "basic",
			m:    12,
			want: 12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.ToMeters(); got != tt.want {
				t.Errorf("Meter.ToMeters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMeter_ToFeet(t *testing.T) {
	tests := []struct {
		name string
		m    Meter
		want Foot
	}{
		{
			name: "basic",
			m:    20,
			want: 20 / 0.3048,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.ToFeet(); got != tt.want {
				t.Errorf("Meter.ToFeet() = %v, want %v", got, tt.want)
			}
		})
	}
}
