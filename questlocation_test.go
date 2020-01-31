package kanka

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

const (
	testQuestLocationIndex  string = "test_data/questlocation_index.json"
	testQuestLocationGet    string = "test_data/questlocation_get.json"
	testQuestLocationCreate string = "test_data/questlocation_create.json"
	testQuestLocationUpdate string = "test_data/questlocation_update.json"
)

func TestQuestLocationService_Index(t *testing.T) {
	qlocs := []*QuestLocation{
		&QuestLocation{
			SimpleQuestLocation: SimpleQuestLocation{
				QuestID:     111,
				LocationID:  222,
				Description: "An inn at the crossroads.",
			},
		},
		&QuestLocation{
			SimpleQuestLocation: SimpleQuestLocation{
				QuestID:     333,
				LocationID:  444,
				Description: "A forest clearing.",
			},
		},
		&QuestLocation{
			SimpleQuestLocation: SimpleQuestLocation{
				QuestID:     555,
				LocationID:  666,
				Description: "A dank cavern...",
			},
		},
	}
	n := time.Now()
	now := &n

	type args struct {
		campID int
		qstID  int
		sync   *time.Time
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    []*QuestLocation
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testQuestLocationIndex,
			args:    args{campID: 5272, qstID: 10394, sync: now},
			want:    qlocs,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testQuestLocationIndex,
			args:    args{campID: -123, qstID: 10394, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid qstID",
			status:  http.StatusOK,
			file:    testQuestLocationIndex,
			args:    args{campID: 5272, qstID: -123, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testQuestLocationIndex,
			args:    args{campID: -123, qstID: -123, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, qstID: -123, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, sync: now},
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

			got, err := c.QuestLocations.Index(test.args.campID, test.args.qstID, test.args.sync)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestQuestLocationService_Get(t *testing.T) {
	qloc := &QuestLocation{
		SimpleQuestLocation: SimpleQuestLocation{
			LocationID:  115366,
			Description: "\n<p>Where the quest takes place</p>\n",
			Role:        "Setting",
			IsPrivate:   false,
		},
		ID:        3137,
		CreatedBy: 0,
		UpdatedBy: 0,
	}

	type args struct {
		campID int
		qstID  int
		qlocID int
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *QuestLocation
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testQuestLocationGet,
			args:    args{campID: 5272, qstID: 10394, qlocID: 3137},
			want:    qloc,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testQuestLocationGet,
			args:    args{campID: -123, qstID: 10394, qlocID: 3137},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid qstID",
			status:  http.StatusOK,
			file:    testQuestLocationGet,
			args:    args{campID: 5272, qstID: -123, qlocID: 3137},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid qlocID",
			status:  http.StatusOK,
			file:    testQuestLocationGet,
			args:    args{campID: 5272, qstID: 10394, qlocID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testQuestLocationGet,
			args:    args{campID: -123, qstID: -123, qlocID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, qlocID: 3137},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, qstID: 10394, qlocID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, qlocID: 3137},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, qlocID: 3137},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, qlocID: 3137},
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

			got, err := c.QuestLocations.Get(test.args.campID, test.args.qstID, test.args.qlocID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestQuestLocationService_Create(t *testing.T) {
	qloc := SimpleQuestLocation{
		QuestID:     777,
		LocationID:  888,
		Description: "A tunnel with a distant bright light!",
	}
	type args struct {
		campID int
		qstID  int
		qloc   SimpleQuestLocation
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *QuestLocation
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testQuestLocationCreate,
			args:    args{campID: 5272, qstID: 10394, qloc: qloc},
			want:    &QuestLocation{SimpleQuestLocation: qloc},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testQuestLocationCreate,
			args:    args{campID: -123, qstID: 10394, qloc: qloc},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid qstID",
			status:  http.StatusOK,
			file:    testQuestLocationCreate,
			args:    args{campID: 5272, qstID: -123, qloc: qloc},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, empty qloc",
			status:  http.StatusOK,
			file:    testQuestLocationCreate,
			args:    args{campID: 5272, qstID: 10394, qloc: SimpleQuestLocation{}},
			want:    &QuestLocation{SimpleQuestLocation: qloc},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testQuestLocationCreate,
			args:    args{campID: -123, qstID: -123, qloc: SimpleQuestLocation{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, qloc: qloc},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, qstID: -123, qloc: SimpleQuestLocation{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, qloc: qloc},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, qloc: qloc},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, qloc: qloc},
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

			got, err := c.QuestLocations.Create(test.args.campID, test.args.qstID, test.args.qloc)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestQuestLocationService_Update(t *testing.T) {
	qloc := SimpleQuestLocation{
		QuestID:     999,
		LocationID:  101010,
		Description: "A royal court.",
	}
	type args struct {
		campID int
		qstID  int
		qlocID int
		qloc   SimpleQuestLocation
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *QuestLocation
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testQuestLocationUpdate,
			args:    args{campID: 5272, qstID: 10394, qlocID: 111, qloc: qloc},
			want:    &QuestLocation{SimpleQuestLocation: qloc, ID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testQuestLocationUpdate,
			args:    args{campID: -123, qstID: 10394, qlocID: 111, qloc: qloc},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid qstID",
			status:  http.StatusOK,
			file:    testQuestLocationUpdate,
			args:    args{campID: 5272, qstID: -123, qlocID: 111, qloc: qloc},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid qlocID",
			status:  http.StatusOK,
			file:    testQuestLocationUpdate,
			args:    args{campID: 5272, qstID: 10394, qlocID: -123, qloc: qloc},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, empty qloc",
			status:  http.StatusOK,
			file:    testQuestLocationUpdate,
			args:    args{campID: 5272, qstID: 10394, qlocID: 111, qloc: SimpleQuestLocation{}},
			want:    &QuestLocation{SimpleQuestLocation: qloc, ID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testQuestLocationUpdate,
			args:    args{campID: -123, qstID: -123, qlocID: -123, qloc: SimpleQuestLocation{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, qlocID: 111, qloc: qloc},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, qstID: -123, qlocID: -123, qloc: SimpleQuestLocation{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, qlocID: 111, qloc: qloc},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, qlocID: 111, qloc: qloc},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, qlocID: 111, qloc: qloc},
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

			got, err := c.QuestLocations.Update(test.args.campID, test.args.qstID, test.args.qlocID, test.args.qloc)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestQuestLocationService_Delete(t *testing.T) {
	type args struct {
		campID int
		qstID  int
		qlocID int
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
			args:    args{campID: 5272, qstID: 10394, qlocID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, invalid campID",
			status:  http.StatusOK,
			args:    args{campID: -123, qstID: 10394, qlocID: 111},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid qstID",
			status:  http.StatusOK,
			args:    args{campID: 5272, qstID: -123, qlocID: 111},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid qlocID",
			status:  http.StatusOK,
			args:    args{campID: 5272, qstID: 10394, qlocID: -123},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid args",
			status:  http.StatusOK,
			args:    args{campID: -123, qstID: -123, qlocID: -123},
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			args:    args{campID: 5272, qstID: 10394, qlocID: 111},
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			args:    args{campID: 5272, qstID: 10394, qlocID: 111},
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			args:    args{campID: 5272, qstID: 10394, qlocID: 111},
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

			err = c.QuestLocations.Delete(test.args.campID, test.args.qstID, test.args.qlocID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
		})
	}
}
