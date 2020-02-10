package kanka

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

const (
	testRelationIndex  string = "test_data/relation_index.json"
	testRelationGet    string = "test_data/relation_get.json"
	testRelationCreate string = "test_data/relation_create.json"
	testRelationUpdate string = "test_data/relation_update.json"
)

func TestRelationService_Index(t *testing.T) {
	rels := []*Relation{
		{
			SimpleRelation: SimpleRelation{
				OwnerID:  111,
				TargetID: 222,
				Relation: "Liege-lord",
				Attitude: 50,
			},
		},
		{
			SimpleRelation: SimpleRelation{
				OwnerID:  333,
				TargetID: 444,
				Relation: "Married",
				Attitude: 80,
			},
		},
		{
			SimpleRelation: SimpleRelation{
				OwnerID:  555,
				TargetID: 666,
				Relation: "Friend",
				Attitude: 20,
			},
		},
	}
	n := time.Now()
	now := &n

	type args struct {
		campID int
		entID  int
		sync   *time.Time
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    []*Relation
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testRelationIndex,
			args:    args{campID: 5272, entID: 430214, sync: now},
			want:    rels,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testRelationIndex,
			args:    args{campID: -123, entID: 430214, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid entID",
			status:  http.StatusOK,
			file:    testRelationIndex,
			args:    args{campID: 5272, entID: -123, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testRelationIndex,
			args:    args{campID: -123, entID: -123, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, entID: -123, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, sync: now},
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

			got, err := c.Relations.Index(test.args.campID, test.args.entID, test.args.sync)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestRelationService_Get(t *testing.T) {
	rel := &Relation{
		SimpleRelation: SimpleRelation{
			OwnerID:   430214,
			TargetID:  420176,
			Relation:  "Rival",
			Attitude:  -50,
			IsPrivate: false,
		},
	}

	type args struct {
		campID int
		entID  int
		relID  int
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *Relation
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testRelationGet,
			args:    args{campID: 5272, entID: 430214, relID: 123},
			want:    rel,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testRelationGet,
			args:    args{campID: -123, entID: 430214, relID: 123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid entID",
			status:  http.StatusOK,
			file:    testRelationGet,
			args:    args{campID: 5272, entID: -123, relID: 123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid relID",
			status:  http.StatusOK,
			file:    testRelationGet,
			args:    args{campID: 5272, entID: 430214, relID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testRelationGet,
			args:    args{campID: -123, entID: -123, relID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, relID: 123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, entID: 430214, relID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, relID: 123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, relID: 123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, relID: 123},
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

			got, err := c.Relations.Get(test.args.campID, test.args.entID, test.args.relID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestRelationService_Create(t *testing.T) {
	rel := SimpleRelation{
		OwnerID:  777,
		TargetID: 888,
		Relation: "Lover",
		Attitude: 100,
	}
	type args struct {
		campID int
		entID  int
		rel    SimpleRelation
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *Relation
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testRelationCreate,
			args:    args{campID: 5272, entID: 430214, rel: rel},
			want:    &Relation{SimpleRelation: rel},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testRelationCreate,
			args:    args{campID: -123, entID: 430214, rel: rel},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid entID",
			status:  http.StatusOK,
			file:    testRelationCreate,
			args:    args{campID: 5272, entID: -123, rel: rel},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "Status OK, valid response, invalid rel (missing Relation)",
			status: http.StatusOK,
			file:   testRelationCreate,
			args: args{campID: 5272, entID: 430214, rel: SimpleRelation{
				OwnerID:  777,
				TargetID: 888,
				Relation: "",
				Attitude: 100,
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "Status OK, valid response, invalid rel (exceeding max Relation length)",
			status: http.StatusOK,
			file:   testRelationCreate,
			args: args{campID: 5272, entID: 430214, rel: SimpleRelation{
				OwnerID:  777,
				TargetID: 888,
				Relation: "0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789123456",
				Attitude: 100,
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "Status OK, valid response, invalid rel (exceeding min Attitude value)",
			status: http.StatusOK,
			file:   testRelationCreate,
			args: args{campID: 5272, entID: 430214, rel: SimpleRelation{
				OwnerID:  777,
				TargetID: 888,
				Relation: "Lover",
				Attitude: -101,
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "Status OK, valid response, invalid rel (exceeding max Attitude value)",
			status: http.StatusOK,
			file:   testRelationCreate,
			args: args{campID: 5272, entID: 430214, rel: SimpleRelation{
				OwnerID:  777,
				TargetID: 888,
				Relation: "Lover",
				Attitude: 101,
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testRelationCreate,
			args:    args{campID: -123, entID: -123, rel: SimpleRelation{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, rel: rel},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, entID: -123, rel: SimpleRelation{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, rel: rel},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, rel: rel},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, rel: rel},
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

			got, err := c.Relations.Create(test.args.campID, test.args.entID, test.args.rel)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestRelationService_Update(t *testing.T) {
	rel := SimpleRelation{
		OwnerID:  999,
		TargetID: 101010,
		Relation: "Destined",
		Attitude: 0,
	}
	type args struct {
		campID int
		entID  int
		relID  int
		rel    SimpleRelation
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *Relation
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testRelationUpdate,
			args:    args{campID: 5272, entID: 430214, relID: 111, rel: rel},
			want:    &Relation{SimpleRelation: rel, ID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testRelationUpdate,
			args:    args{campID: -123, entID: 430214, relID: 111, rel: rel},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid entID",
			status:  http.StatusOK,
			file:    testRelationUpdate,
			args:    args{campID: 5272, entID: -123, relID: 111, rel: rel},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid relID",
			status:  http.StatusOK,
			file:    testRelationUpdate,
			args:    args{campID: 5272, entID: 430214, relID: -123, rel: rel},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid rel",
			status:  http.StatusOK,
			file:    testRelationUpdate,
			args:    args{campID: 5272, entID: 430214, relID: 111, rel: SimpleRelation{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "Status OK, valid response, invalid rel (missing Relation)",
			status: http.StatusOK,
			file:   testRelationUpdate,
			args: args{campID: 5272, entID: 430214, rel: SimpleRelation{
				OwnerID:  999,
				TargetID: 101010,
				Relation: "",
				Attitude: 100,
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "Status OK, valid response, invalid rel (exceeding max Relation length)",
			status: http.StatusOK,
			file:   testRelationUpdate,
			args: args{campID: 5272, entID: 430214, rel: SimpleRelation{
				OwnerID:  999,
				TargetID: 101010,
				Relation: "0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789123456",
				Attitude: 100,
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "Status OK, valid response, invalid rel (exceeding min Attitude value)",
			status: http.StatusOK,
			file:   testRelationUpdate,
			args: args{campID: 5272, entID: 430214, rel: SimpleRelation{
				OwnerID:  999,
				TargetID: 101010,
				Relation: "Destined",
				Attitude: -101,
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "Status OK, valid response, invalid rel (exceeding max Attitude value)",
			status: http.StatusOK,
			file:   testRelationUpdate,
			args: args{campID: 5272, entID: 430214, rel: SimpleRelation{
				OwnerID:  999,
				TargetID: 101010,
				Relation: "Destined",
				Attitude: 101,
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testRelationUpdate,
			args:    args{campID: -123, entID: -123, relID: -123, rel: SimpleRelation{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, relID: 111, rel: rel},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, entID: -123, relID: -123, rel: SimpleRelation{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, relID: 111, rel: rel},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, relID: 111, rel: rel},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, relID: 111, rel: rel},
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

			got, err := c.Relations.Update(test.args.campID, test.args.entID, test.args.relID, test.args.rel)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestRelationService_Delete(t *testing.T) {
	type args struct {
		campID int
		entID  int
		relID  int
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
			args:    args{campID: 5272, entID: 430214, relID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, invalid campID",
			status:  http.StatusOK,
			args:    args{campID: -123, entID: 430214, relID: 111},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid entID",
			status:  http.StatusOK,
			args:    args{campID: -123, entID: -123, relID: 111},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid relID",
			status:  http.StatusOK,
			args:    args{campID: 5272, entID: 430214, relID: -123},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid args",
			status:  http.StatusOK,
			args:    args{campID: -123, entID: -123, relID: -123},
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			args:    args{campID: 5272, entID: 430214, relID: 111},
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			args:    args{campID: 5272, entID: 430214, relID: 111},
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			args:    args{campID: 5272, entID: 430214, relID: 111},
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

			err = c.Relations.Delete(test.args.campID, test.args.entID, test.args.relID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
		})
	}
}
