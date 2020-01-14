package kanka

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

const (
	testFamilyIndex  string = "test_data/family_index.json"
	testFamilyGet    string = "test_data/family_get.json"
	testFamilyCreate string = "test_data/family_create.json"
	testFamilyUpdate string = "test_data/family_update.json"
)

func TestFamilyService_Index(t *testing.T) {
	fams := []*Family{
		&Family{
			SimpleFamily: SimpleFamily{
				Name:       "Baratheon",
				LocationID: 111,
			},
		},
		&Family{
			SimpleFamily: SimpleFamily{
				Name:       "Lannister",
				LocationID: 222,
			},
		},
		&Family{
			SimpleFamily: SimpleFamily{
				Name:       "Reed",
				LocationID: 333,
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
		want    []*Family
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testFamilyIndex,
			args:    args{campID: 5272, sync: now},
			want:    fams,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testFamilyIndex,
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

			got, err := c.Families.Index(test.args.campID, test.args.sync)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFamilyService_Get(t *testing.T) {
	fam := &Family{
		SimpleFamily: SimpleFamily{
			Name:       "Stark",
			Entry:      "\n<p>House Stark</p>\n",
			Image:      "families/EIwz3WTvqcbVVcUYBue9O4DQ9dETmlI6JYbkOGx1.png",
			IsPrivate:  false,
			Tags:       []int{35131},
			Type:       "Royal",
			LocationID: 115368,
		},
		ID:             16439,
		ImageFull:      "https://kanka-user-assets.s3.eu-central-1.amazonaws.com/families/EIwz3WTvqcbVVcUYBue9O4DQ9dETmlI6JYbkOGx1.png",
		ImageThumb:     "https://kanka-user-assets.s3.eu-central-1.amazonaws.com/families/EIwz3WTvqcbVVcUYBue9O4DQ9dETmlI6JYbkOGx1_thumb.png",
		HasCustomImage: true,
		EntityID:       436884,
		CreatedBy:      5600,
		UpdatedBy:      5600,
		Members:        []int{118427},
	}

	type args struct {
		campID int
		famID  int
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *Family
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testFamilyGet,
			args:    args{campID: 5272, famID: 16439},
			want:    fam,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testFamilyGet,
			args:    args{campID: -123, famID: 16439},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid famID",
			status:  http.StatusOK,
			file:    testFamilyGet,
			args:    args{campID: 5272, famID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testFamilyGet,
			args:    args{campID: -123, famID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, famID: 16439},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, famID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, famID: 16439},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, famID: 16439},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, famID: 16439},
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

			got, err := c.Families.Get(test.args.campID, test.args.famID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFamilyService_Create(t *testing.T) {
	fam := SimpleFamily{
		Name:       "Greyjoy",
		LocationID: 111,
	}
	type args struct {
		campID int
		f      SimpleFamily
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *Family
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testFamilyCreate,
			args:    args{campID: 5272, f: fam},
			want:    &Family{SimpleFamily: fam},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testFamilyCreate,
			args:    args{campID: -123, f: fam},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid family",
			status:  http.StatusOK,
			file:    testFamilyCreate,
			args:    args{campID: 5272, f: SimpleFamily{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testFamilyCreate,
			args:    args{campID: -123, f: SimpleFamily{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, f: fam},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, f: SimpleFamily{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, f: fam},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, f: fam},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, f: fam},
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

			got, err := c.Families.Create(test.args.campID, test.args.f)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFamilyService_Update(t *testing.T) {
	fam := SimpleFamily{
		Name:       "Martell",
		LocationID: 222,
	}
	type args struct {
		campID int
		famID  int
		f      SimpleFamily
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *Family
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testFamilyUpdate,
			args:    args{campID: 5272, famID: 111, f: fam},
			want:    &Family{SimpleFamily: fam, ID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testFamilyUpdate,
			args:    args{campID: -123, famID: 111, f: fam},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid famID",
			status:  http.StatusOK,
			file:    testFamilyUpdate,
			args:    args{campID: 5272, famID: -123, f: fam},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid fam",
			status:  http.StatusOK,
			file:    testFamilyUpdate,
			args:    args{campID: 5272, famID: 111, f: SimpleFamily{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testFamilyUpdate,
			args:    args{campID: -123, famID: -123, f: SimpleFamily{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, famID: 111, f: fam},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, famID: -123, f: SimpleFamily{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, famID: 111, f: fam},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, famID: 111, f: fam},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, famID: 111, f: fam},
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

			got, err := c.Families.Update(test.args.campID, test.args.famID, test.args.f)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFamilyService_Delete(t *testing.T) {
	type args struct {
		campID int
		famID  int
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
			args:    args{campID: 5272, famID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, invalid campID",
			status:  http.StatusOK,
			args:    args{campID: -123, famID: 111},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid famID",
			status:  http.StatusOK,
			args:    args{campID: 5272, famID: -123},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid args",
			status:  http.StatusOK,
			args:    args{campID: -123, famID: -123},
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			args:    args{campID: 5272, famID: 111},
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			args:    args{campID: 5272, famID: 111},
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			args:    args{campID: 5272, famID: 111},
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

			err = c.Families.Delete(test.args.campID, test.args.famID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
		})
	}
}
