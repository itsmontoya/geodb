package geodb

import (
	"reflect"
	"testing"
)

func TestPolygon_IsWithin(t *testing.T) {
	type fields struct {
		locs []Location
	}

	type args struct {
		l Location
	}

	tests := []struct {
		name       string
		fields     fields
		args       args
		wantWithin bool
	}{
		{
			name: "match",
			fields: fields{
				locs: []Location{
					MakeLocation(0, 0),
					MakeLocation(0, 4),
					MakeLocation(4, 4),
					MakeLocation(4, 0),
				},
			},
			args: args{
				l: MakeLocation(2, 2),
			},
			wantWithin: true,
		},
		{
			name: "no match",
			fields: fields{
				locs: []Location{
					MakeLocation(0, 0),
					MakeLocation(0, 4),
					MakeLocation(4, 4),
					MakeLocation(4, 0),
				},
			},
			args: args{
				l: MakeLocation(6, 2),
			},
			wantWithin: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := NewPolygon(tt.fields.locs)
			if err != nil {
				t.Fatal(err)
			}

			if gotWithin := p.IsWithin(tt.args.l); gotWithin != tt.wantWithin {
				t.Errorf("Polygon.IsWithin() = %v, want %v", gotWithin, tt.wantWithin)
			}
		})
	}
}

func TestPolygon_Center(t *testing.T) {
	type fields struct {
		locations []Location
	}

	tests := []struct {
		name   string
		fields fields
		want   Location
	}{
		{
			name: "basic",
			fields: fields{
				locations: []Location{
					MakeLocation(0, 0),
					MakeLocation(0, 4),
					MakeLocation(4, 4),
					MakeLocation(4, 0),
				},
			},
			want: MakeLocation(2, 2),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := NewPolygon(tt.fields.locations)
			if err != nil {
				t.Fatal(err)
			}

			if got := p.Center(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Polygon.Center() = %v, want %v", got.String(), tt.want.String())
			}
		})
	}
}
