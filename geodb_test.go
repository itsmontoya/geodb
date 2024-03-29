package geodb

import (
	"encoding/json"
	"log"
	"os"
	"reflect"
	"testing"
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
	type args struct {
		regionSize Meter
	}

	tests := []struct {
		name string
		args args
		want *GeoDB
	}{
		{
			name: "basic",
			args: args{
				regionSize: 3000,
			},
			want: &GeoDB{
				s: store{
					regionSize: 3000,
					rs:         nil,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.regionSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGeoDB_Insert(t *testing.T) {
	type fields struct {
		regionSize Meter
		rs         []*region
	}

	type args struct {
		key   string
		shape Shape
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "basic",
			fields: fields{
				regionSize: 3000,
			},
			args: args{
				key:   "test_key",
				shape: NewRadius(3000, testCities["Portland, OR"].Location()),
			},
		},
		{
			name: "with existing region",
			fields: fields{
				regionSize: 3000,
				rs: []*region{
					{
						radius: 10000,
						center: testCities["Portland, OR"].Location(),
					},
				},
			},
			args: args{
				key:   "test_key",
				shape: NewRadius(3000, testCities["Portland, OR"].Location()),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GeoDB{
				s: store{
					regionSize: tt.fields.regionSize,
					rs:         tt.fields.rs,
				},
			}

			err := g.Insert(tt.args.key, tt.args.shape)
			if (err != nil) != tt.wantErr {
				t.Errorf("GeoDB.RegionsLen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGeoDB_GetMatches(t *testing.T) {
	type fields struct {
		regionSize Meter
		rs         []*region
	}

	type args struct {
		l Location
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
				regionSize: 10000,
				rs: []*region{
					{
						radius: 3000,
						center: testCities["Portland, OR"].Location(),
						ts: []entry{
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
				},
			},
			args: args{
				testCities["Portland, OR"].Location(),
			},
			wantMatches: []string{"1", "2", "3"},
			wantErr:     false,
		},
		{
			name: "with mixed matches and non-matches",
			fields: fields{
				regionSize: 10000,
				rs: []*region{
					{
						radius: 10000,
						center: testCities["Portland, OR"].Location(),
						ts: []entry{
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
								key: "4",
								shape: NewRadius(
									100,
									MakeLocation(
										45.50921710781631,
										-122.68495391699925),
								),
							},
						},
					},
				},
			},
			args: args{
				testCities["Portland, OR"].Location(),
			},
			wantMatches: []string{"1", "2", "3"},
			wantErr:     false,
		},
		{
			name: "outside of region - no matches",
			fields: fields{
				regionSize: 10000,
				rs: []*region{
					{
						radius: 10000,
						center: testCities["Portland, OR"].Location(),
						ts: []entry{
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
								key: "4",
								shape: NewRadius(
									100,
									MakeLocation(
										45.50921710781631,
										-122.68495391699925),
								),
							},
						},
					},
				},
			},
			args: args{
				testCities["Seattle, WA"].Location(),
			},
			wantMatches: nil,
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GeoDB{
				s: store{
					regionSize: tt.fields.regionSize,
					rs:         tt.fields.rs,
				},
			}

			gotMatches, err := g.GetMatches(tt.args.l)
			if (err != nil) != tt.wantErr {
				t.Errorf("GeoDB.GetMatches() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(gotMatches, tt.wantMatches) {
				t.Errorf("GeoDB.GetMatches() = %v, want %v", gotMatches, tt.wantMatches)
			}
		})
	}
}

func TestGeoDB_GetMatches_poly(t *testing.T) {
	type fields struct {
		locs []Location
	}

	type args struct {
		l Location
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
			wantMatch: true,
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
			wantMatch: false,
		},
		{
			name: "Triangle - Match",
			fields: fields{
				locs: []Location{
					MakeLocation(32.7933, -97.1566),
					MakeLocation(32.6540, -97.2686),
					MakeLocation(32.7848, -97.4444),
				},
			},
			args: args{
				l: MakeLocation(32.7450, -97.3582),
			},
			wantMatch: true,
		},
		{
			name: "Triangle - No match",
			fields: fields{
				locs: []Location{
					MakeLocation(32.7933, -97.1566),
					MakeLocation(32.6540, -97.2686),
					MakeLocation(32.7848, -97.4444),
				},
			},
			args: args{
				l: MakeLocation(32.6714, -97.3193),
			},
			wantMatch: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := New(3000)
			p, err := NewPolygon(tt.fields.locs)
			if err != nil {
				t.Fatal(err)
			}

			if err := db.Insert("test_0", p); err != nil {
				t.Fatal(err)
			}

			gotMatches, err := db.GetMatches(tt.args.l)
			if (err != nil) != tt.wantErr {
				t.Errorf("GeoDB.GetMatches() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantMatch && len(gotMatches) == 0 {
				t.Errorf("wanted match and got no match")
			} else if !tt.wantMatch && len(gotMatches) > 0 {
				t.Errorf("wanted no match and got match")
			}
		})
	}
}

func TestGeoDB_RegionsLen(t *testing.T) {
	type fields struct {
		// region size
		regionSize Meter
		// Internal region store
		rs []*region
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
				regionSize: 3000,
				rs: []*region{
					newRegion(NewRadius(3000, testCities["Portland, OR"].Location()), 3000),
				},
			},
			wantN:   1,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GeoDB{
				s: store{
					regionSize: tt.fields.regionSize,
					rs:         tt.fields.rs,
				},
			}

			gotN, err := g.RegionsLen()
			if (err != nil) != tt.wantErr {
				t.Errorf("GeoDB.RegionsLen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotN != tt.wantN {
				t.Errorf("GeoDB.RegionsLen() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestGeoDB_EntriesLen(t *testing.T) {
	type fields struct {
		// region size
		regionSize Meter
		// Internal region store
		rs []*region
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
				regionSize: 3000,
				rs: []*region{
					{
						radius: 3000,
						center: testCities["Portland, OR"].Location(),
						ts: []entry{
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
				},
			},
			wantN:   3,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GeoDB{
				s: store{
					regionSize: tt.fields.regionSize,
					rs:         tt.fields.rs,
				},
			}

			gotN, err := g.EntriesLen()
			if (err != nil) != tt.wantErr {
				t.Errorf("GeoDB.EntriesLen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotN != tt.wantN {
				t.Errorf("GeoDB.EntriesLen() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

/*
func BenchmarkgeodbInserting(b *testing.B) {
	gt := New(Meter(1000 * 1000))
	j := 0
	n := len(testEntries)
	for i := 0; i < b.N; i++ {
		e := testEntries[j]
		if j++; j == n {
			j = 0
		}

		// Set radius at 3 km
		gt.Insert(getTestKey(e), e.Lat, e.Lon, 3000)
	}

	b.ReportAllocs()
}

func BenchmarkgeodbPoly_Square(b *testing.B) {
	var locs []*gmath.Point
	locs = append(locs, &gmath.Point{Y: 1, X: 1})
	locs = append(locs, &gmath.Point{Y: 3, X: 1})
	locs = append(locs, &gmath.Point{Y: 3, X: 3})
	locs = append(locs, &gmath.Point{Y: 1, X: 3})
	tgt := NewPolygon(locs)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if tgt.IsWithin(2, 2) {
			testCnt++
		}
	}

	b.ReportAllocs()
}

func BenchmarkgeodbPoly_Square_20(b *testing.B) {
	var locs []*gmath.Point
	locs = append(locs, &gmath.Point{Y: 1, X: 1})
	locs = append(locs, &gmath.Point{Y: 2, X: 1})
	locs = append(locs, &gmath.Point{Y: 2, X: 2})
	locs = append(locs, &gmath.Point{Y: 1, X: 2})
	tgt := NewPolygon(locs)

	var tis []*gmath.Point
	tis = append(tis,
		&gmath.Point{X: 0, Y: 0},
		&gmath.Point{X: 1, Y: 0},
		&gmath.Point{X: 2, Y: 0},
		&gmath.Point{X: 3, Y: 0},
		&gmath.Point{X: 4, Y: 0},
		&gmath.Point{X: 5, Y: 0},
		&gmath.Point{X: 0, Y: 1},
		&gmath.Point{X: 1, Y: 1},
		&gmath.Point{X: 2, Y: 1},
		&gmath.Point{X: 3, Y: 1},
		&gmath.Point{X: 4, Y: 1},
		&gmath.Point{X: 5, Y: 1},
		&gmath.Point{X: 0, Y: 2},
		&gmath.Point{X: 1, Y: 2},
		&gmath.Point{X: 2, Y: 2},
		&gmath.Point{X: 3, Y: 2},
		&gmath.Point{X: 4, Y: 2},
		&gmath.Point{X: 5, Y: 2},
		&gmath.Point{X: 0, Y: 3},
		&gmath.Point{X: 1, Y: 3},
		&gmath.Point{X: 2, Y: 3},
		&gmath.Point{X: 3, Y: 3},
		&gmath.Point{X: 4, Y: 3},
		&gmath.Point{X: 5, Y: 3},
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, pt := range tis {
			if tgt.IsWithin(pt.Y, pt.X) {
				testCnt++
			}
		}
	}

	b.ReportAllocs()
}

func BenchmarkgeodbLookup_Single(b *testing.B) {
	portland := testCities["Portland, OR"]
	tgt := newTarget(getTestKey(portland), 300, MakeLocation(portland.Lat, portland.Lon))
	latR := Degree(portland.Lat).toRadians()
	lonR := Degree(portland.Lon).toRadians()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if tgt.isWithinRadius(latR, lonR) {
			testCnt++
		}
	}

	b.ReportAllocs()
}

func BenchmarkgeodbLookup1_000(b *testing.B) {
	benchLookup(b, 1000)
}

func BenchmarkgeodbLookup10_000(b *testing.B) {
	benchLookup(b, 10000)
}

func BenchmarkgeodbLookup100_000(b *testing.B) {
	benchLookup(b, 100000)
}

func BenchmarkgeodbLookup_USA(b *testing.B) {
	benchLookup(b, len(testEntries))
}

func BenchmarkgeodbListLookup1_000(b *testing.B) {
	benchListLookup(b, 1000)
}

func BenchmarkgeodbListLookup10_000(b *testing.B) {
	benchListLookup(b, 10000)
}

func BenchmarkgeodbListLookup100_000(b *testing.B) {
	benchListLookup(b, 100000)
}

func BenchmarkgeodbListLookup_USA(b *testing.B) {
	benchListLookup(b, len(testEntries))
}

func BenchmarkRTreeInserting(b *testing.B) {
	rt := rtreego.NewTree(2, 25, 50)
	j := 0
	n := len(testEntries)
	for i := 0; i < b.N; i++ {
		e := testEntries[j]
		if j++; j == n {
			j = 0
		}

		// Set radius at 3 km
		rt.Insert(getRTreeLoc(getTestKey(e), e.Lat, e.Lon, 3000))
	}

	b.ReportAllocs()
}

func BenchmarkRTreeLookup1_000(b *testing.B) {
	benchRTreeLookup(b, 1000)
}

func BenchmarkRTreeLookup10_000(b *testing.B) {
	benchRTreeLookup(b, 10000)
}

func BenchmarkRTreeLookup100_000(b *testing.B) {
	benchRTreeLookup(b, 100000)
}

func BenchmarkRTreeLookup_USA(b *testing.B) {
	benchRTreeLookup(b, len(testEntries))
}

func benchLookup(b *testing.B, n int) {
	// 1200 km radius
	gt := New(Meter(1000 * 1200))
	b.StopTimer()
	portland := testCities["Portland, OR"]

	if n > len(testEntries) {
		n = len(testEntries)
	}

	for i := 0; i < n; i++ {
		e := testEntries[i]
		// Set radius at 3 km
		gt.Insert(getTestKey(e), e.Lat, e.Lon, 3000)
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		testMatches = gt.GetMatches(portland.Lat, portland.Lon)
	}

	b.ReportAllocs()
}

func benchRTreeLookup(b *testing.B, n int) {
	rt := rtreego.NewTree(2, 25, 50)
	b.StopTimer()
	portland := testCities["Portland, OR"]

	if n > len(testEntries) {
		n = len(testEntries)
	}

	for i := 0; i < n; i++ {
		e := testEntries[i]
		// Set radius at 3 km
		rt.Insert(getRTreeLoc(getTestKey(e), e.Lat, e.Lon, 3000))
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		rect, _ := rtreego.NewRect(rtreego.Point{portland.Lon, portland.Lat}, []float64{0.008, 0.008})
		testSpacials = rt.SearchIntersect(rect)
	}

	b.ReportAllocs()
}

*/

/*
type rtreeloc struct {
	Key      string
	Location *rtreego.Rect
}

func (r *rtreeloc) Bounds() *rtreego.Rect {
	return r.Location
}

func getRTreeLoc(key string, lat, lon, radius float64) *rtreeloc {
	dist := radius / 69.172 // Dist is now in Degrees
	adjustedDist := dist * math.Cos(lat*(math.Pi/180))

	x := lon - dist
	y := lat - adjustedDist

	rect, _ := rtreego.NewRect(
		rtreego.Point{x, y},
		[]float64{2 * dist, 2 * adjustedDist},
	)

	return &rtreeloc{
		Key:      key,
		Location: rect,
	}
}


*/

func populateTestCities() (err error) {
	var f *os.File
	if f, err = os.Open("./testing/testCities.json"); err != nil {
		return
	}
	defer f.Close()
	return json.NewDecoder(f).Decode(&testCities)
}
