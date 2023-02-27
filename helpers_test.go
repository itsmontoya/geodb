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
				a: testCities["Portland, OR"].Location(),
				b: testCities["New York, NY"].Location(),
			},
			want: Meter(3.9256283176791137e+06),
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
