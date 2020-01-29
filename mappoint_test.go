package kanka

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

const (
	testMapPointIndex  string = "test_data/mappoint_index.json"
	testMapPointCreate string = "test_data/mappoint_create.json"
	testMapPointUpdate string = "test_data/mappoint_update.json"
)

func TestMapPointService_Index(t *testing.T) {
	mps := []*MapPoint{
		&MapPoint{
			SimpleMapPoint: SimpleMapPoint{
				TargetEntityID: 0,
				Name:           "Courtyard",
				AxisX:          608,
				AxisY:          471,
				Color:          "#2e7e35",
				Icon:           "castle-emblem",
				Shape:          "circle",
				Size:           "standard",
			},
		},
		&MapPoint{
			SimpleMapPoint: SimpleMapPoint{
				TargetEntityID: 0,
				Name:           "Smithy",
				AxisX:          557,
				AxisY:          689,
				Color:          "#000000",
				Icon:           "anvil",
				Shape:          "square",
				Size:           "standard",
			},
		},
		&MapPoint{
			SimpleMapPoint: SimpleMapPoint{
				TargetEntityID: 0,
				Name:           "Armory",
				AxisX:          593,
				AxisY:          355,
				Color:          "#0000ff",
				Icon:           "shield",
				Shape:          "square",
				Size:           "standard",
			},
		},
	}
	n := time.Now()
	now := &n

	type args struct {
		campID int
		locID  int
		sync   *time.Time
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    []*MapPoint
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testMapPointIndex,
			args:    args{campID: 5272, locID: 115368, sync: now},
			want:    mps,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testMapPointIndex,
			args:    args{campID: -123, locID: 115368, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid locID",
			status:  http.StatusOK,
			file:    testMapPointIndex,
			args:    args{campID: 5272, locID: -123, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testMapPointIndex,
			args:    args{campID: -123, locID: -123, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, locID: 115368, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, locID: -123, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, locID: 115368, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, locID: 115368, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, locID: 115368, sync: now},
			want:    nil,
			wantErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f, err := os.Open(test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()

			c, _ := testClient(test.status, f)

			got, err := c.MapPoints.Index(test.args.campID, test.args.locID, test.args.sync)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestMapPointService_Create(t *testing.T) {
	mp := SimpleMapPoint{
		TargetEntityID: 0,
		Name:           "Tower",
		AxisX:          300,
		AxisY:          300,
		Color:          "white",
		Icon:           "tower",
		Shape:          "circle",
		Size:           "standard",
	}
	type args struct {
		campID int
		locID  int
		mp     SimpleMapPoint
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *MapPoint
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testMapPointCreate,
			args:    args{campID: 5272, locID: 115368, mp: mp},
			want:    &MapPoint{SimpleMapPoint: mp},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testMapPointCreate,
			args:    args{campID: -123, locID: 115368, mp: mp},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid locID",
			status:  http.StatusOK,
			file:    testMapPointCreate,
			args:    args{campID: 5272, locID: -123, mp: mp},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "Status OK, valid response, invalid mp (missing Color)",
			status: http.StatusOK,
			file:   testMapPointCreate,
			args: args{campID: 5272, locID: 115368, mp: SimpleMapPoint{
				TargetEntityID: 0,
				Name:           "Tower",
				AxisX:          300,
				AxisY:          300,
				Color:          "",
				Icon:           "tower",
				Shape:          "circle",
				Size:           "standard",
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "Status OK, valid response, invalid mp (missing Icon)",
			status: http.StatusOK,
			file:   testMapPointCreate,
			args: args{campID: 5272, locID: 115368, mp: SimpleMapPoint{
				TargetEntityID: 0,
				Name:           "Tower",
				AxisX:          300,
				AxisY:          300,
				Color:          "white",
				Icon:           "",
				Shape:          "circle",
				Size:           "standard",
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "Status OK, valid response, invalid mp (missing Shape)",
			status: http.StatusOK,
			file:   testMapPointCreate,
			args: args{campID: 5272, locID: 115368, mp: SimpleMapPoint{
				TargetEntityID: 0,
				Name:           "Tower",
				AxisX:          300,
				AxisY:          300,
				Color:          "white",
				Icon:           "tower",
				Shape:          "",
				Size:           "standard",
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "Status OK, valid response, invalid mp (missing Size)",
			status: http.StatusOK,
			file:   testMapPointCreate,
			args: args{campID: 5272, locID: 115368, mp: SimpleMapPoint{
				TargetEntityID: 0,
				Name:           "Tower",
				AxisX:          300,
				AxisY:          300,
				Color:          "white",
				Icon:           "tower",
				Shape:          "circle",
				Size:           "",
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testMapPointCreate,
			args:    args{campID: -123, locID: -123, mp: SimpleMapPoint{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, locID: 115368, mp: mp},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, locID: -123, mp: SimpleMapPoint{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, locID: 115368, mp: mp},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, locID: 115368, mp: mp},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, locID: 115368, mp: mp},
			want:    nil,
			wantErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f, err := os.Open(test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()

			c, _ := testClient(test.status, f)

			got, err := c.MapPoints.Create(test.args.campID, test.args.locID, test.args.mp)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
