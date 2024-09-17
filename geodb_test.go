package geodb

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/gdbu/stringset"
)

var testCities map[string]*Coordinates

func TestMain(m *testing.M) {
	if err := populateTestCities(); err != nil {
		log.Fatal(err)
		return
	}

	os.Exit(m.Run())
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *GeoDB
	}{
		{
			name: "basic",
			want: &GeoDB{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGeoDB_Insert(t *testing.T) {
	type args struct {
		es []entry
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "basic",
			args: args{
				es: []entry{
					{
						key:   "test_key",
						shape: NewRadius(3000, testCities["Portland, OR"].Location()),
					},
				},
			},
		},

		{
			name: "multiple",
			args: args{
				es: []entry{
					{
						key:   "portland",
						shape: NewRadius(3000, testCities["Portland, OR"].Location()),
					},
					{
						key:   "seattle",
						shape: NewRadius(3000, testCities["Seattle, WA"].Location()),
					},
					{
						key:   "chicago",
						shape: NewRadius(3000, testCities["Chicago, IL"].Location()),
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New()
			for _, e := range tt.args.es {
				g.Insert(e.key, e.shape)
			}
		})
	}
}

func TestGeoDB_GetMatches(t *testing.T) {
	type fields struct {
		es []*entry
	}

	type args struct {
		c Coordinates
	}

	tests := []struct {
		name        string
		fields      fields
		args        args
		wantMatches []string
		wantErr     bool
	}{
		{
			name: "with matches",
			fields: fields{
				es: []*entry{
					{
						key:   "1",
						shape: NewRadius(100, testCities["Portland, OR"].Location()),
					},
					{
						key:   "2",
						shape: NewRadius(200, testCities["Portland, OR"].Location()),
					},
					{
						key:   "3",
						shape: NewRadius(300, testCities["Portland, OR"].Location()),
					},
				},
			},
			args: args{
				c: *testCities["Portland, OR"],
			},
			wantMatches: []string{"1", "2", "3"},
			wantErr:     false,
		},
		{
			name: "with mixed matches and non-matches",
			fields: fields{
				es: []*entry{
					{
						key:   "1",
						shape: NewRadius(100, testCities["Portland, OR"].Location()),
					},
					{
						key:   "2",
						shape: NewRadius(200, testCities["Portland, OR"].Location()),
					},
					{
						key:   "3",
						shape: NewRadius(300, testCities["Portland, OR"].Location()),
					},
					{
						key:   "4",
						shape: NewRadius(100, NewCoordinates(45.50921710781631, -122.68495391699925).Location()),
					},
				},
			},
			args: args{
				c: *testCities["Portland, OR"],
			},
			wantMatches: []string{"1", "2", "3"},
			wantErr:     false,
		},
		{
			name: "outside of region - no matches",
			fields: fields{
				es: []*entry{
					{
						key:   "1",
						shape: NewRadius(100, testCities["Portland, OR"].Location()),
					},
					{
						key:   "2",
						shape: NewRadius(200, testCities["Portland, OR"].Location()),
					},
					{
						key:   "3",
						shape: NewRadius(300, testCities["Portland, OR"].Location()),
					},
					{
						key:   "4",
						shape: NewRadius(100, NewCoordinates(45.50921710781631, -122.68495391699925).Location()),
					},
				},
			},
			args: args{
				c: *testCities["Seattle, WA"],
			},
			wantMatches: []string{},
			wantErr:     false,
		},
		{
			name: "900 points, -90_-180",
			fields: fields{
				es: makeTestPolys(30, 30),
			},
			args: args{
				c: MakeCoordinates(-90, -180),
			},
			wantMatches: []string{"-90.000000_-180.000000"},
			wantErr:     false,
		},
		{
			name: "900 points, -84_-168 (bordered)",
			fields: fields{
				es: makeTestPolys(30, 30),
			},
			args: args{
				c: MakeCoordinates(-84, -168),
			},
			wantMatches: []string{"-90.000000_-180.000000", "-84.000000_-168.000000"},
			wantErr:     false,
		},
		{
			name: "10_000 points, -90_-180",
			fields: fields{
				es: makeTestPolys(100, 100),
			},
			args: args{
				c: MakeCoordinates(-90, -180),
			},
			wantMatches: []string{"-90.000000_-180.000000"},
			wantErr:     false,
		},
		{
			name: "10_000 points, -89_-177 (bordered)",
			fields: fields{
				es: makeTestPolys(100, 100),
			},
			args: args{
				c: MakeCoordinates(-89, -177),
			},
			wantMatches: []string{"-90.000000_-180.000000", "-89.000000_-177.000000"},
			wantErr:     false,
		},
		{
			name: "100_000 points, -90.000000_-180.000000",
			fields: fields{
				es: makeTestPolys(1000, 100),
			},
			args: args{
				c: MakeCoordinates(-90, -180),
			},
			wantMatches: []string{"-90.000000_-180.000000"},
			wantErr:     false,
		},
		{
			name: "100_000 points, -89_-177 (bordered)",
			fields: fields{
				es: makeTestPolys(1000, 100),
			},
			args: args{
				c: MakeCoordinates(-90, -177),
			},
			wantMatches: []string{"-90.000000_-180.000000", "-90.000000_-177.000000"},
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New()
			for _, e := range tt.fields.es {
				g.Insert(e.key, e.shape)
			}

			for _, e := range tt.fields.es {
				rect := e.Rect()
				loc := rect.Center()
				ms := g.GetMatches(loc.Coordinates())
				if len(ms) == 0 {
					t.Fatalf("No matches for %s", e.key)
				}

				set := stringset.MakeMap(ms...)
				if !set.Has(e.key) {
					t.Fatalf("Invalid match, want to contain %s and received %s", e.key, ms)
				}
			}

			gotMatches := g.GetMatches(tt.args.c)
			if !reflect.DeepEqual(gotMatches, tt.wantMatches) {
				t.Errorf("GeoDB.GetMatches() = %v, want %v", gotMatches, tt.wantMatches)
			}
		})
	}
}

func makeTestPolys(rows, columns int) (out []*entry) {
	minLat := -90
	maxLat := 90
	minLon := -180
	maxLon := 180
	latPerEntry := Degree((maxLat - minLat) / rows)
	lonPerEntry := Degree((maxLon - minLon) / columns)

	curLat := Degree(minLat)
	curLon := Degree(minLon)
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			var e entry
			e.key = fmt.Sprintf("%f_%f", curLat, curLon)
			e.shape, _ = NewPolygon([]Coordinates{
				MakeCoordinates(curLat, curLon),
				MakeCoordinates(curLat, curLon+lonPerEntry),
				MakeCoordinates(curLat+latPerEntry, curLon+lonPerEntry),
				MakeCoordinates(curLat+latPerEntry, curLon),
			})

			curLat += latPerEntry
			curLon += lonPerEntry
			out = append(out, &e)
		}
	}

	return
}

func TestGeoDB_GetMatches_poly(t *testing.T) {
	type fields struct {
		coords []Coordinates
	}

	type args struct {
		c Coordinates
	}

	tests := []struct {
		name   string
		fields fields
		args   args

		wantMatch bool
		wantErr   bool
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
			wantMatch: true,
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
			wantMatch: false,
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
			wantMatch: true,
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
			wantMatch: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := New()
			p, err := NewPolygon(tt.fields.coords)
			if err != nil {
				t.Fatal(err)
			}

			db.Insert("test_0", p)

			gotMatches := db.GetMatches(tt.args.c)
			if tt.wantMatch && len(gotMatches) == 0 {
				t.Errorf("wanted match and got no match")
			} else if !tt.wantMatch && len(gotMatches) > 0 {
				t.Errorf("wanted no match and got match")
			}
		})
	}
}

func TestGeoDB_Len(t *testing.T) {
	type fields struct {
		es []*entry
	}

	tests := []struct {
		name    string
		fields  fields
		wantN   int
		wantErr bool
	}{
		{
			name: "basic",
			fields: fields{
				es: []*entry{
					{
						key:   "1",
						shape: NewRadius(100, testCities["Portland, OR"].Location()),
					},
					{
						key:   "2",
						shape: NewRadius(200, testCities["Portland, OR"].Location()),
					},
					{
						key:   "3",
						shape: NewRadius(300, testCities["Portland, OR"].Location()),
					},
				},
			},
			wantN:   3,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New()
			for _, e := range tt.fields.es {
				g.Insert(e.key, e.shape)
			}

			gotN := g.Len()
			if gotN != tt.wantN {
				t.Errorf("GeoDB.EntriesLen() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func populateTestCities() (err error) {
	var f *os.File
	if f, err = os.Open("./testing/testCities.json"); err != nil {
		return
	}
	defer f.Close()
	return json.NewDecoder(f).Decode(&testCities)
}
