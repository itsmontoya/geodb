package geodb

import "testing"

func Test_region_insert(t *testing.T) {
	type fields struct {
		shape  Shape
		radius Meter
	}

	type args struct {
		e entry
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		wantOk bool
	}{
		{
			fields: fields{
				shape:  NewRadius(3000, testCities["Portland, OR"].Location()),
				radius: 10000,
			},
			args: args{
				e: entry{
					shape: NewRadius(3000, testCities["Portland, OR"].Location()),
				},
			},
			wantOk: true,
		},
		{
			fields: fields{
				shape:  NewRadius(3000, testCities["Portland, OR"].Location()),
				radius: 10000,
			},
			args: args{
				e: entry{
					shape: NewRadius(3000, testCities["Seattle, WA"].Location()),
				},
			},
			wantOk: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := newRegion(tt.fields.shape, tt.fields.radius)
			if gotOk := r.insert(tt.args.e); gotOk != tt.wantOk {
				t.Errorf("region.insert() = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
