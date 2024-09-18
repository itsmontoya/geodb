package geodb

import "testing"

func TestRect_DoesNotOverlap(t *testing.T) {
	type fields struct {
		Min Coordinates
		Max Coordinates
	}

	type args struct {
		in *Rect
	}
	/*
	   {{34.11641744750447 -118.08129048380394} {34.13086055249552 -118.06384351619607}}

	   {{33.87660474394087 -118.14813020764912} {-117.9653451221156 -117.97666923823937}}
	   {{33.87660474394087 -117.97666923823937} {34.14722755249552 -117.80520826882963}}
	*/

	tests := []struct {
		name             string
		fields           fields
		args             args
		wantNotContained bool
	}{
		{
			name: "basic",
			fields: fields{
				Min: MakeCoordinates(33.87660474394087, -118.14813020764912),
				Max: MakeCoordinates(33.87660474394087, -118.14813020764912),
			},
			args:             args{},
			wantNotContained: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rect{
				Min: tt.fields.Min,
				Max: tt.fields.Max,
			}
			if gotNotContained := r.DoesNotOverlap(tt.args.in); gotNotContained != tt.wantNotContained {
				t.Errorf("Rect.DoesNotOverlap() = %v, want %v", gotNotContained, tt.wantNotContained)
			}
		})
	}
}
