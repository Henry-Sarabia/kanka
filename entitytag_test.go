package kanka

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

const (
	testEntityTagIndex  string = "test_data/entitytag_index.json"
	testEntityTagGet    string = "test_data/entitytag_get.json"
	testEntityTagCreate string = "test_data/entitytag_create.json"
	testEntityTagUpdate string = "test_data/entitytag_update.json"
)

func TestEntityTagService_Index(t *testing.T) {
	tags := []*EntityTag{
		&EntityTag{
			SimpleEntityTag: SimpleEntityTag{
				EntityID: 111,
				TagID:    222,
			},
		},
		&EntityTag{
			SimpleEntityTag: SimpleEntityTag{
				EntityID: 333,
				TagID:    444,
			},
		},
		&EntityTag{
			SimpleEntityTag: SimpleEntityTag{
				EntityID: 555,
				TagID:    666,
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
		want    []*EntityTag
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testEntityTagIndex,
			args:    args{campID: 5272, entID: 430214, sync: now},
			want:    tags,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testEntityTagIndex,
			args:    args{campID: -123, entID: 430214, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid entID",
			status:  http.StatusOK,
			file:    testEntityTagIndex,
			args:    args{campID: 5272, entID: -123, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testEntityTagIndex,
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

			got, err := c.EntityTags.Index(test.args.campID, test.args.entID, test.args.sync)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestEntityTagService_Get(t *testing.T) {
	tag := &EntityTag{
		SimpleEntityTag: SimpleEntityTag{
			EntityID: 430214,
			TagID:    34696,
		},
		ID: 100999,
	}

	type args struct {
		campID int
		entID  int
		tagID  int
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *EntityTag
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testEntityTagGet,
			args:    args{campID: 5272, entID: 430214, tagID: 34696},
			want:    tag,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testEntityTagGet,
			args:    args{campID: -123, entID: 430214, tagID: 34696},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid entID",
			status:  http.StatusOK,
			file:    testEntityTagGet,
			args:    args{campID: 5272, entID: -123, tagID: 34696},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid tagID",
			status:  http.StatusOK,
			file:    testEntityTagGet,
			args:    args{campID: 5272, entID: 430214, tagID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testEntityTagGet,
			args:    args{campID: -123, entID: -123, tagID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, tagID: 34696},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, entID: 430214, tagID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, tagID: 34696},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, tagID: 34696},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, tagID: 34696},
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

			got, err := c.EntityTags.Get(test.args.campID, test.args.entID, test.args.tagID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestEntityTagService_Create(t *testing.T) {
	tag := SimpleEntityTag{
		EntityID: 777,
		TagID:    888,
	}
	type args struct {
		campID int
		entID  int
		tag    SimpleEntityTag
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *EntityTag
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testEntityTagCreate,
			args:    args{campID: 5272, entID: 430214, tag: tag},
			want:    &EntityTag{SimpleEntityTag: tag},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testEntityTagCreate,
			args:    args{campID: -123, entID: 430214, tag: tag},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid entID",
			status:  http.StatusOK,
			file:    testEntityTagCreate,
			args:    args{campID: -123, entID: -123, tag: tag},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, empty tag",
			status:  http.StatusOK,
			file:    testEntityTagCreate,
			args:    args{campID: 5272, entID: 430214, tag: SimpleEntityTag{}},
			want:    &EntityTag{SimpleEntityTag: tag},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testEntityTagCreate,
			args:    args{campID: -123, entID: -123, tag: SimpleEntityTag{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, tag: tag},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, entID: -123, tag: SimpleEntityTag{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, tag: tag},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, tag: tag},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, tag: tag},
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

			got, err := c.EntityTags.Create(test.args.campID, test.args.entID, test.args.tag)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestEntityTagService_Update(t *testing.T) {
	tag := SimpleEntityTag{
		EntityID: 999,
		TagID:    101010,
	}
	type args struct {
		campID int
		entID  int
		tagID  int
		tag    SimpleEntityTag
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *EntityTag
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testEntityTagUpdate,
			args:    args{campID: 5272, entID: 430214, tagID: 111, tag: tag},
			want:    &EntityTag{SimpleEntityTag: tag, ID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testEntityTagUpdate,
			args:    args{campID: -123, entID: 430214, tagID: 111, tag: tag},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid entID",
			status:  http.StatusOK,
			file:    testEntityTagUpdate,
			args:    args{campID: -123, entID: -123, tagID: 111, tag: tag},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid tagID",
			status:  http.StatusOK,
			file:    testEntityTagUpdate,
			args:    args{campID: 5272, entID: 430214, tagID: -123, tag: tag},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, empty tag",
			status:  http.StatusOK,
			file:    testEntityTagUpdate,
			args:    args{campID: 5272, entID: 430214, tagID: 111, tag: SimpleEntityTag{}},
			want:    &EntityTag{SimpleEntityTag: tag, ID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testEntityTagUpdate,
			args:    args{campID: -123, entID: -123, tagID: -123, tag: SimpleEntityTag{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, tagID: 111, tag: tag},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, entID: -123, tagID: -123, tag: SimpleEntityTag{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, tagID: 111, tag: tag},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, tagID: 111, tag: tag},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, tagID: 111, tag: tag},
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

			got, err := c.EntityTags.Update(test.args.campID, test.args.entID, test.args.tagID, test.args.tag)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestEntityTagService_Delete(t *testing.T) {
	type args struct {
		campID int
		entID  int
		tagID  int
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
			args:    args{campID: 5272, entID: 430214, tagID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, invalid campID",
			status:  http.StatusOK,
			args:    args{campID: -123, entID: 430214, tagID: 111},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid entID",
			status:  http.StatusOK,
			args:    args{campID: -123, entID: -123, tagID: 111},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid tagID",
			status:  http.StatusOK,
			args:    args{campID: 5272, entID: 430214, tagID: -123},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid args",
			status:  http.StatusOK,
			args:    args{campID: -123, entID: -123, tagID: -123},
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			args:    args{campID: 5272, entID: 430214, tagID: 111},
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			args:    args{campID: 5272, entID: 430214, tagID: 111},
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			args:    args{campID: 5272, entID: 430214, tagID: 111},
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

			err = c.EntityTags.Delete(test.args.campID, test.args.entID, test.args.tagID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
		})
	}
}
