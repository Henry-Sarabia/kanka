package kanka

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

const (
	testTagIndex  string = "test_data/tag_index.json"
	testTagGet    string = "test_data/tag_get.json"
	testTagCreate string = "test_data/tag_create.json"
	testTagUpdate string = "test_data/tag_update.json"
)

func TestTagService_Index(t *testing.T) {
	tags := []*Tag{
		&Tag{
			SimpleTag: SimpleTag{
				Name:  "Flora",
				Color: "green",
			},
		},
		&Tag{
			SimpleTag: SimpleTag{
				Name:  "Fauna",
				Color: "red",
			},
		},
		&Tag{
			SimpleTag: SimpleTag{
				Name:  "Dungeon",
				Color: "black",
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
		want    []*Tag
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testTagIndex,
			args:    args{campID: 5272, sync: now},
			want:    tags,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testTagIndex,
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

			got, err := c.Tags.Index(test.args.campID, test.args.sync)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestTagService_Get(t *testing.T) {
	tag := &Tag{
		SimpleTag: SimpleTag{
			Name:      "Seven Kingdoms",
			Entry:     "\n<p>Just some references</p>\n",
			Image:     "tags/clthzBLiqZTET8tKUKibQZeKcOjQosz6nNWj5C77.jpeg",
			IsPrivate: false,
			Tags:      []int{},
			Type:      "Lore",
			Color:     "red",
		},
		ID:             35131,
		ImageFull:      "https://kanka-user-assets.s3.eu-central-1.amazonaws.com/tags/clthzBLiqZTET8tKUKibQZeKcOjQosz6nNWj5C77.jpeg",
		ImageThumb:     "https://kanka-user-assets.s3.eu-central-1.amazonaws.com/tags/clthzBLiqZTET8tKUKibQZeKcOjQosz6nNWj5C77_thumb.jpeg",
		HasCustomImage: true,
		EntityID:       436885,
		CreatedBy:      5600,
		UpdatedBy:      5600,
		Entities:       []int{436884, 436892, 437024, 442245, 442249, 424136, 443498, 443499},
	}

	type args struct {
		campID int
		tagID  int
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *Tag
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testTagGet,
			args:    args{campID: 5272, tagID: 35131},
			want:    tag,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testTagGet,
			args:    args{campID: -123, tagID: 35131},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid tagID",
			status:  http.StatusOK,
			file:    testTagGet,
			args:    args{campID: 5272, tagID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testTagGet,
			args:    args{campID: -123, tagID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, tagID: 35131},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, tagID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, tagID: 35131},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, tagID: 35131},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, tagID: 35131},
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

			got, err := c.Tags.Get(test.args.campID, test.args.tagID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestTagService_Create(t *testing.T) {
	tag := SimpleTag{
		Name:  "Important",
		Color: "yellow",
	}
	type args struct {
		campID int
		tag    SimpleTag
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *Tag
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testTagCreate,
			args:    args{campID: 5272, tag: tag},
			want:    &Tag{SimpleTag: tag},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testTagCreate,
			args:    args{campID: -123, tag: tag},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid tag",
			status:  http.StatusOK,
			file:    testTagCreate,
			args:    args{campID: 5272, tag: SimpleTag{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testTagCreate,
			args:    args{campID: -123, tag: SimpleTag{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, tag: tag},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, tag: SimpleTag{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, tag: tag},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, tag: tag},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, tag: tag},
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

			got, err := c.Tags.Create(test.args.campID, test.args.tag)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestTagService_Update(t *testing.T) {
	tag := SimpleTag{
		Name:  "Worldbuilding",
		Color: "purple",
	}
	type args struct {
		campID int
		tagID  int
		tag    SimpleTag
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *Tag
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testTagUpdate,
			args:    args{campID: 5272, tagID: 111, tag: tag},
			want:    &Tag{SimpleTag: tag, ID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testTagUpdate,
			args:    args{campID: -123, tagID: 111, tag: tag},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid tagID",
			status:  http.StatusOK,
			file:    testTagUpdate,
			args:    args{campID: 5272, tagID: -123, tag: tag},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid tag",
			status:  http.StatusOK,
			file:    testTagUpdate,
			args:    args{campID: 5272, tagID: 111, tag: SimpleTag{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testTagUpdate,
			args:    args{campID: -123, tagID: -123, tag: SimpleTag{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, tagID: 111, tag: tag},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, tagID: -123, tag: SimpleTag{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, tagID: 111, tag: tag},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, tagID: 111, tag: tag},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, tagID: 111, tag: tag},
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

			got, err := c.Tags.Update(test.args.campID, test.args.tagID, test.args.tag)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestTagService_Delete(t *testing.T) {
	type args struct {
		campID int
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
			args:    args{campID: 5272, tagID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, invalid campID",
			status:  http.StatusOK,
			args:    args{campID: -123, tagID: 111},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid tagID",
			status:  http.StatusOK,
			args:    args{campID: 5272, tagID: -123},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid args",
			status:  http.StatusOK,
			args:    args{campID: -123, tagID: -123},
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			args:    args{campID: 5272, tagID: 111},
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			args:    args{campID: 5272, tagID: 111},
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			args:    args{campID: 5272, tagID: 111},
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

			err = c.Tags.Delete(test.args.campID, test.args.tagID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
		})
	}
}
