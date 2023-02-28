package geodb

import "testing"

func TestLocation_String(t *testing.T) {
	type fields struct {
		lat Radian
		lon Radian
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "basic",
			fields: fields{
				lat: Degree(3).toRadians(),
				lon: Degree(6).toRadians(),
			},
			want: "Lat: 3.000000 / Lon: 6.000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Location{
				lat: tt.fields.lat,
				lon: tt.fields.lon,
			}

			if got := l.String(); got != tt.want {
				t.Errorf("Location.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
