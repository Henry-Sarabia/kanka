package kanka

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

const (
	testAttributeIndex  string = "test_data/attribute_index.json"
	testAttributeGet    string = "test_data/attribute_get.json"
	testAttributeCreate string = "test_data/attribute_create.json"
	testAttributeUpdate string = "test_data/attribute_update.json"
)

func TestAttributeService_Index(t *testing.T) {
	atrs := []*Attribute{
		&Attribute{
			SimpleAttribute: SimpleAttribute{
				Name:  "Troops",
				Value: "500",
			},
		},
		&Attribute{
			SimpleAttribute: SimpleAttribute{
				Name:  "Population",
				Value: "2000",
			},
		},
		&Attribute{
			SimpleAttribute: SimpleAttribute{
				Name:  "Title",
				Value: "King",
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
		want    []*Attribute
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testAttributeIndex,
			args:    args{campID: 5272, entID: 430214, sync: now},
			want:    atrs,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testAttributeIndex,
			args:    args{campID: -123, entID: 430214, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid entID",
			status:  http.StatusOK,
			file:    testAttributeIndex,
			args:    args{campID: 5272, entID: -123, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testAttributeIndex,
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

			got, err := c.Attributes.Index(test.args.campID, test.args.entID, test.args.sync)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestAttributeService_Get(t *testing.T) {
	atr := &Attribute{
		SimpleAttribute: SimpleAttribute{
			Name:         "Item",
			Value:        "Pocket Watch (Wind-up Chronicle)",
			IsPrivate:    false,
			Type:         "",
			APIKey:       "",
			DefaultOrder: 0,
		},
		ID:        318053,
		EntityID:  430214,
		CreatedBy: 5600,
		UpdatedBy: 5600,
	}

	type args struct {
		campID int
		entID  int
		atrID  int
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *Attribute
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testAttributeGet,
			args:    args{campID: 5272, entID: 430214, atrID: 318053},
			want:    atr,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testAttributeGet,
			args:    args{campID: -123, entID: 430214, atrID: 318053},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid entID",
			status:  http.StatusOK,
			file:    testAttributeGet,
			args:    args{campID: 5272, entID: -123, atrID: 318053},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid atrID",
			status:  http.StatusOK,
			file:    testAttributeGet,
			args:    args{campID: 5272, entID: 430214, atrID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testAttributeGet,
			args:    args{campID: -123, entID: -123, atrID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, atrID: 318053},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, entID: 430214, atrID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, atrID: 318053},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, atrID: 318053},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, atrID: 318053},
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

			got, err := c.Attributes.Get(test.args.campID, test.args.entID, test.args.atrID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestAttributeService_Create(t *testing.T) {
	atr := SimpleAttribute{
		Name:  "Race",
		Value: "elf",
	}
	type args struct {
		campID int
		entID  int
		atr    SimpleAttribute
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *Attribute
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testAttributeCreate,
			args:    args{campID: 5272, entID: 430214, atr: atr},
			want:    &Attribute{SimpleAttribute: atr},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testAttributeCreate,
			args:    args{campID: -123, entID: 430214, atr: atr},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid entID",
			status:  http.StatusOK,
			file:    testAttributeCreate,
			args:    args{campID: -123, entID: -123, atr: atr},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid attribute",
			status:  http.StatusOK,
			file:    testAttributeCreate,
			args:    args{campID: 5272, entID: 430214, atr: SimpleAttribute{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testAttributeCreate,
			args:    args{campID: -123, entID: -123, atr: SimpleAttribute{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, atr: atr},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, entID: -123, atr: SimpleAttribute{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, atr: atr},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, atr: atr},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, atr: atr},
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

			got, err := c.Attributes.Create(test.args.campID, test.args.entID, test.args.atr)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestAttributeService_Update(t *testing.T) {
	atr := SimpleAttribute{
		Name:  "Conquests",
		Value: "2",
	}
	type args struct {
		campID int
		entID  int
		atrID  int
		atr    SimpleAttribute
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *Attribute
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testAttributeUpdate,
			args:    args{campID: 5272, entID: 430214, atrID: 111, atr: atr},
			want:    &Attribute{SimpleAttribute: atr, ID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testAttributeUpdate,
			args:    args{campID: -123, entID: 430214, atrID: 111, atr: atr},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid entID",
			status:  http.StatusOK,
			file:    testAttributeUpdate,
			args:    args{campID: -123, entID: -123, atrID: 111, atr: atr},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid atrID",
			status:  http.StatusOK,
			file:    testAttributeUpdate,
			args:    args{campID: 5272, entID: 430214, atrID: -123, atr: atr},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid atr",
			status:  http.StatusOK,
			file:    testAttributeUpdate,
			args:    args{campID: 5272, entID: 430214, atrID: 111, atr: SimpleAttribute{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testAttributeUpdate,
			args:    args{campID: -123, entID: -123, atrID: -123, atr: SimpleAttribute{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, atrID: 111, atr: atr},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, entID: -123, atrID: -123, atr: SimpleAttribute{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, atrID: 111, atr: atr},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, atrID: 111, atr: atr},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, atrID: 111, atr: atr},
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

			got, err := c.Attributes.Update(test.args.campID, test.args.entID, test.args.atrID, test.args.atr)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestAttributeService_Delete(t *testing.T) {
	type args struct {
		campID int
		entID  int
		atrID  int
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
			args:    args{campID: 5272, entID: 430214, atrID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, invalid campID",
			status:  http.StatusOK,
			args:    args{campID: -123, entID: 430214, atrID: 111},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid entID",
			status:  http.StatusOK,
			args:    args{campID: -123, entID: -123, atrID: 111},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid atrID",
			status:  http.StatusOK,
			args:    args{campID: 5272, entID: 430214, atrID: -123},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid args",
			status:  http.StatusOK,
			args:    args{campID: -123, entID: -123, atrID: -123},
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			args:    args{campID: 5272, entID: 430214, atrID: 111},
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			args:    args{campID: 5272, entID: 430214, atrID: 111},
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			args:    args{campID: 5272, entID: 430214, atrID: 111},
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

			err = c.Attributes.Delete(test.args.campID, test.args.entID, test.args.atrID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
		})
	}
}
