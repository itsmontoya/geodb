package geodb

import (
	"reflect"
	"testing"
)

func TestGetDistance(t *testing.T) {
	type args struct {
		a Location
		b Location
	}
	tests := []struct {
		name string
		args args
		want Unit
	}{
		{
			name: "Portland to New York",
			args: args{
				a: testMakeLocation(testCities["Portland, OR"]),
				b: testMakeLocation(testCities["New York, NY"]),
			},
			want: Meter(3.925129057503389e+06),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDistance(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
