package geodb

import (
	"testing"
)

func TestNewPolygon(t *testing.T) {
	type args struct {
		coords []Coordinates
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "four locations",
			args: args{
				coords: []Coordinates{
					MakeCoordinates(0, 0),
					MakeCoordinates(0, 4),
					MakeCoordinates(4, 4),
					MakeCoordinates(4, 0),
				},
			},
			wantErr: false,
		},
		{
			name: "three locations",
			args: args{
				coords: []Coordinates{
					MakeCoordinates(0, 0),
					MakeCoordinates(0, 4),
					MakeCoordinates(4, 4),
				},
			},
			wantErr: false,
		},
		{
			name: "two locations",
			args: args{
				coords: []Coordinates{
					MakeCoordinates(0, 0),
					MakeCoordinates(0, 4),
				},
			},
			wantErr: true,
		},
		{
			name: "one location",
			args: args{
				coords: []Coordinates{
					MakeCoordinates(0, 0),
				},
			},
			wantErr: true,
		},
		{
			name: "empty",
			args: args{
				coords: []Coordinates{},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewPolygon(tt.args.coords)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPolygon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestPolygon_IsWithin(t *testing.T) {
	type fields struct {
		coords []Coordinates
	}

	type args struct {
		c Coordinates
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
				coords: []Coordinates{
					MakeCoordinates(0, 0),
					MakeCoordinates(0, 4),
					MakeCoordinates(4, 4),
					MakeCoordinates(4, 0),
				},
			},
			args: args{
				c: MakeCoordinates(2, 2),
			},
			wantWithin: true,
		},
		{
			name: "no match",
			fields: fields{
				coords: []Coordinates{
					MakeCoordinates(0, 0),
					MakeCoordinates(0, 4),
					MakeCoordinates(4, 4),
					MakeCoordinates(4, 0),
				},
			},
			args: args{
				c: MakeCoordinates(6, 2),
			},
			wantWithin: false,
		},
		{
			name: "Triangle - Match",
			fields: fields{
				coords: []Coordinates{
					MakeCoordinates(32.7933, -97.1566),
					MakeCoordinates(32.6540, -97.2686),
					MakeCoordinates(32.7848, -97.4444),
				},
			},
			args: args{
				c: MakeCoordinates(32.7450, -97.3582),
			},
			wantWithin: true,
		},
		{
			name: "Triangle - No match",
			fields: fields{
				coords: []Coordinates{
					MakeCoordinates(32.7933, -97.1566),
					MakeCoordinates(32.6540, -97.2686),
					MakeCoordinates(32.7848, -97.4444),
				},
			},
			args: args{
				c: MakeCoordinates(32.6714, -97.3193),
			},
			wantWithin: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := NewPolygon(tt.fields.coords)
			if err != nil {
				t.Fatal(err)
			}

			if gotWithin := p.IsWithin(tt.args.c); gotWithin != tt.wantWithin {
				t.Errorf("Polygon.IsWithin() = %v, want %v", gotWithin, tt.wantWithin)
			}
		})
	}
}
