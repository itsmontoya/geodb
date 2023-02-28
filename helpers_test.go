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

func Test_getCorners(t *testing.T) {
	type args struct {
		locs []Location
	}
	tests := []struct {
		name        string
		args        args
		wantLowLat  Degree
		wantHighLat Degree
		wantLowLon  Degree
		wantHighLon Degree
	}{
		{
			name: "low / high",
			args: args{
				locs: []Location{
					MakeLocation(4, 4),
					MakeLocation(6, 6),
				},
			},
			wantLowLat:  4,
			wantHighLat: 6,
			wantLowLon:  4,
			wantHighLon: 6,
		},
		{
			name: "high / low",
			args: args{
				locs: []Location{
					MakeLocation(6, 6),
					MakeLocation(4, 4),
				},
			},
			wantLowLat:  4,
			wantHighLat: 6,
			wantLowLon:  4,
			wantHighLon: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLowLat, gotHighLat, gotLowLon, gotHighLon := getCorners(tt.args.locs)
			switch {
			case gotLowLat != tt.wantLowLat:
				t.Errorf("getCorners() gotLowLat = %v, want %v", gotLowLat, tt.wantLowLat)
			case gotHighLat != tt.wantHighLat:
				t.Errorf("getCorners() gotHighLat = %v, want %v", gotHighLat, tt.wantHighLat)
			case gotLowLon != tt.wantLowLon:
				t.Errorf("getCorners() gotLowLon = %v, want %v", gotLowLon, tt.wantLowLon)
			case gotHighLon != tt.wantHighLon:
				t.Errorf("getCorners() gotHighLon = %v, want %v", gotHighLon, tt.wantHighLon)
			}
		})
	}
}
