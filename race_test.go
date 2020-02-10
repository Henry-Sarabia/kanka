package kanka

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

const (
	testRaceIndex  string = "test_data/race_index.json"
	testRaceGet    string = "test_data/race_get.json"
	testRaceCreate string = "test_data/race_create.json"
	testRaceUpdate string = "test_data/race_update.json"
)

func TestRaceService_Index(t *testing.T) {
	races := []*Race{
		{
			SimpleRace: SimpleRace{
				Name: "Human",
				Type: "Bipedal",
			},
		},
		{
			SimpleRace: SimpleRace{
				Name: "Halfling",
				Type: "Bipedal",
			},
		},
		{
			SimpleRace: SimpleRace{
				Name: "Centaur",
				Type: "Quadrapedal",
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
		want    []*Race
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testRaceIndex,
			args:    args{campID: 5272, sync: now},
			want:    races,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testRaceIndex,
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

			got, err := c.Races.Index(test.args.campID, test.args.sync)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestRaceService_Get(t *testing.T) {
	race := &Race{
		SimpleRace: SimpleRace{
			Name:      "Locatha",
			Entry:     "\n<p>Fish people!</p>\n",
			Image:     "races/0QRnoZ5GK8aOMMMGIIRycEhODN4KPXHAInSmwed2.jpeg",
			IsPrivate: false,
			Tags:      []int{35131},
			Type:      "Amphibious",
			RaceID:    44044,
		},
		ID:             42063,
		ImageFull:      "https://kanka-user-assets.s3.eu-central-1.amazonaws.com/races/0QRnoZ5GK8aOMMMGIIRycEhODN4KPXHAInSmwed2.jpeg",
		ImageThumb:     "https://kanka-user-assets.s3.eu-central-1.amazonaws.com/races/0QRnoZ5GK8aOMMMGIIRycEhODN4KPXHAInSmwed2_thumb.jpeg",
		HasCustomImage: true,
		EntityID:       424136,
		CreatedBy:      5600,
		UpdatedBy:      5600,
	}

	type args struct {
		campID int
		raceID int
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *Race
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testRaceGet,
			args:    args{campID: 5272, raceID: 42063},
			want:    race,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testRaceGet,
			args:    args{campID: -123, raceID: 42063},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid raceID",
			status:  http.StatusOK,
			file:    testRaceGet,
			args:    args{campID: 5272, raceID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testRaceGet,
			args:    args{campID: -123, raceID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, raceID: 42063},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, raceID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, raceID: 42063},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, raceID: 42063},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, raceID: 42063},
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

			got, err := c.Races.Get(test.args.campID, test.args.raceID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestRaceService_Create(t *testing.T) {
	race := SimpleRace{
		Name: "Elf",
		Type: "Bipedal",
	}
	type args struct {
		campID int
		race   SimpleRace
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *Race
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testRaceCreate,
			args:    args{campID: 5272, race: race},
			want:    &Race{SimpleRace: race},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testRaceCreate,
			args:    args{campID: -123, race: race},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid race",
			status:  http.StatusOK,
			file:    testRaceCreate,
			args:    args{campID: 5272, race: SimpleRace{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testRaceCreate,
			args:    args{campID: -123, race: SimpleRace{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, race: race},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, race: SimpleRace{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, race: race},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, race: race},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, race: race},
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

			got, err := c.Races.Create(test.args.campID, test.args.race)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestRaceService_Update(t *testing.T) {
	race := SimpleRace{
		Name: "Orc",
		Type: "Bipedal",
	}
	type args struct {
		campID int
		raceID int
		race   SimpleRace
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *Race
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testRaceUpdate,
			args:    args{campID: 5272, raceID: 111, race: race},
			want:    &Race{SimpleRace: race, ID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testRaceUpdate,
			args:    args{campID: -123, raceID: 111, race: race},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid raceID",
			status:  http.StatusOK,
			file:    testRaceUpdate,
			args:    args{campID: 5272, raceID: -123, race: race},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid race",
			status:  http.StatusOK,
			file:    testRaceUpdate,
			args:    args{campID: 5272, raceID: 111, race: SimpleRace{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testRaceUpdate,
			args:    args{campID: -123, raceID: -123, race: SimpleRace{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, raceID: 111, race: race},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, raceID: -123, race: SimpleRace{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, raceID: 111, race: race},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, raceID: 111, race: race},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, raceID: 111, race: race},
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

			got, err := c.Races.Update(test.args.campID, test.args.raceID, test.args.race)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestRaceService_Delete(t *testing.T) {
	type args struct {
		campID int
		raceID int
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
			args:    args{campID: 5272, raceID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, invalid campID",
			status:  http.StatusOK,
			args:    args{campID: -123, raceID: 111},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid raceID",
			status:  http.StatusOK,
			args:    args{campID: 5272, raceID: -123},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid args",
			status:  http.StatusOK,
			args:    args{campID: -123, raceID: -123},
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			args:    args{campID: 5272, raceID: 111},
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			args:    args{campID: 5272, raceID: 111},
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			args:    args{campID: 5272, raceID: 111},
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

			err = c.Races.Delete(test.args.campID, test.args.raceID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
		})
	}
}
