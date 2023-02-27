package geodb

import (
	"reflect"
	"testing"
)

func TestRadius_IsWithin(t *testing.T) {
	type fields struct {
		radius Meter
		center Location
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
			name: "direct match",
			fields: fields{
				radius: 3000,
				center: MakeLocation(45.526387, -122.671891),
			},
			args: args{
				l: MakeLocation(45.526387, -122.671891),
			},
			wantWithin: true,
		},
		{
			name: "match",
			fields: fields{
				radius: 3000,
				center: MakeLocation(45.526387, -122.671891),
			},
			args: args{
				l: MakeLocation(45.526442, -122.673248),
			},
			wantWithin: true,
		},
		{
			name: "no match",
			fields: fields{
				radius: 3000,
				center: MakeLocation(45.526387, -122.671891),
			},
			args: args{
				l: MakeLocation(45.577992, -122.598070),
			},
			wantWithin: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewRadius(tt.fields.radius, tt.fields.center)
			if gotWithin := p.IsWithin(tt.args.l); gotWithin != tt.wantWithin {
				t.Errorf("Radius.IsWithin() = %v, want %v", gotWithin, tt.wantWithin)
			}
		})
	}
}

func TestRadius_Center(t *testing.T) {
	type fields struct {
		radius Meter
		center Location
	}

	tests := []struct {
		name   string
		fields fields
		want   Location
	}{
		{
			name: "basic",
			fields: fields{
				center: MakeLocation(2, 2),
			},
			want: MakeLocation(2, 2),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewRadius(tt.fields.radius, tt.fields.center)
			if got := p.Center(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Radius.Center() = %v, want %v", got, tt.want)
			}
		})
	}
}
