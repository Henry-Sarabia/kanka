package kanka

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

const (
	testLocationIndex  string = "test_data/location_index.json"
	testLocationGet    string = "test_data/location_get.json"
	testLocationCreate string = "test_data/location_create.json"
	testLocationUpdate string = "test_data/location_update.json"
)

func TestLocationService_Index(t *testing.T) {
	locs := []*Location{
		{
			SimpleLocation: SimpleLocation{
				Name: "King's Landing",
				Type: "Capital",
			},
		},
		{
			SimpleLocation: SimpleLocation{
				Name: "Dragonstone",
				Type: "Castle",
			},
		},
		{
			SimpleLocation: SimpleLocation{
				Name: "Crossroads Inn",
				Type: "Inn",
			},
		},
	}
	n := time.Now()
	now := &n

	type args struct {
		campID int
		sync   *time.Time
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    []*Location
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testLocationIndex,
			args:    args{campID: 5272, sync: now},
			want:    locs,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testLocationIndex,
			args:    args{campID: -123, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, sync: now},
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

			got, err := c.Locations.Index(test.args.campID, test.args.sync)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestLocationService_Get(t *testing.T) {
	loc := &Location{
		SimpleLocation: SimpleLocation{
			Name:             "Winterfell",
			Entry:            "\n<p>The jewel of the North.</p>\n",
			Image:            "locations/ox87nkFQWMn9tTpLuXNk56fq0Du2V3HjocFl9ROY.jpeg",
			IsPrivate:        false,
			Tags:             []int{35115},
			Type:             "Castle",
			Map:              "locations/7MHptkcOx4MpyAxPokfZCB4bGVDLEXq9nCHe2Tex.jpeg",
			ParentLocationID: 115366,
		},
		ID:             115368,
		ImageFull:      "https://kanka-user-assets.s3.eu-central-1.amazonaws.com/locations/ox87nkFQWMn9tTpLuXNk56fq0Du2V3HjocFl9ROY.jpeg",
		ImageThumb:     "https://kanka-user-assets.s3.eu-central-1.amazonaws.com/locations/ox87nkFQWMn9tTpLuXNk56fq0Du2V3HjocFl9ROY_thumb.jpeg",
		HasCustomImage: true,
		EntityID:       436726,
		CreatedBy:      5600,
		UpdatedBy:      5600,
		IsMapPrivate:   0,
	}

	type args struct {
		campID int
		locID  int
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *Location
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testLocationGet,
			args:    args{campID: 5272, locID: 116623},
			want:    loc,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testLocationGet,
			args:    args{campID: -123, locID: 116623},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid locID",
			status:  http.StatusOK,
			file:    testLocationGet,
			args:    args{campID: 5272, locID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testLocationGet,
			args:    args{campID: -123, locID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, locID: 116623},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, locID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, locID: 116623},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, locID: 116623},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, locID: 116623},
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

			got, err := c.Locations.Get(test.args.campID, test.args.locID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestLocationService_Create(t *testing.T) {
	sl := SimpleLocation{
		Name: "Dorne",
		Type: "Kingdom",
	}
	type args struct {
		campID int
		loc    SimpleLocation
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *Location
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testLocationCreate,
			args:    args{campID: 5272, loc: sl},
			want:    &Location{SimpleLocation: sl},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testLocationCreate,
			args:    args{campID: -123, loc: sl},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid location",
			status:  http.StatusOK,
			file:    testLocationCreate,
			args:    args{campID: 5272, loc: SimpleLocation{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testLocationCreate,
			args:    args{campID: -123, loc: SimpleLocation{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, loc: sl},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, loc: SimpleLocation{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, loc: sl},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, loc: sl},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, loc: sl},
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

			got, err := c.Locations.Create(test.args.campID, test.args.loc)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestLocationService_Update(t *testing.T) {
	loc := SimpleLocation{
		Name: "Iron Isles",
		Type: "Kingdom",
	}
	type args struct {
		campID int
		locID  int
		loc    SimpleLocation
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *Location
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testLocationUpdate,
			args:    args{campID: 5272, locID: 111, loc: loc},
			want:    &Location{SimpleLocation: loc, ID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testLocationUpdate,
			args:    args{campID: -123, locID: 111, loc: loc},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid locID",
			status:  http.StatusOK,
			file:    testLocationUpdate,
			args:    args{campID: 5272, locID: -123, loc: loc},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid loc",
			status:  http.StatusOK,
			file:    testLocationUpdate,
			args:    args{campID: 5272, locID: 111, loc: SimpleLocation{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testLocationUpdate,
			args:    args{campID: -123, locID: -123, loc: SimpleLocation{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, locID: 111, loc: loc},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, locID: -123, loc: SimpleLocation{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, locID: 111, loc: loc},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, locID: 111, loc: loc},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, locID: 111, loc: loc},
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

			got, err := c.Locations.Update(test.args.campID, test.args.locID, test.args.loc)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestLocationService_Delete(t *testing.T) {
	type args struct {
		campID int
		locID  int
	}
	tests := []struct {
		name    string
		status  int
		args    args
		wantErr bool
	}{
		{
			name:    "StatusOK, valid args",
			status:  http.StatusOK,
			args:    args{campID: 5272, locID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, invalid campID",
			status:  http.StatusOK,
			args:    args{campID: -123, locID: 111},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid locID",
			status:  http.StatusOK,
			args:    args{campID: 5272, locID: -123},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid args",
			status:  http.StatusOK,
			args:    args{campID: -123, locID: -123},
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			args:    args{campID: 5272, locID: 111},
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			args:    args{campID: 5272, locID: 111},
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			args:    args{campID: 5272, locID: 111},
			wantErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f, err := os.Open(testFileEmpty)
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()

			c, _ := testClient(test.status, f)

			err = c.Locations.Delete(test.args.campID, test.args.locID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
		})
	}
}
