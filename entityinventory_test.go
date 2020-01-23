package kanka

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

const (
	testEntityInventoryIndex  string = "test_data/entityinventory_index.json"
	testEntityInventoryCreate string = "test_data/entityinventory_create.json"
	testEntityInventoryUpdate string = "test_data/entityinventory_update.json"
)

func TestEntityInventoryService_Index(t *testing.T) {
	invs := []*EntityInventory{
		&EntityInventory{
			SimpleEntityInventory: SimpleEntityInventory{
				EntityID: 111,
				ItemID:   222,
				Amount:   1,
			},
		},
		&EntityInventory{
			SimpleEntityInventory: SimpleEntityInventory{
				EntityID: 333,
				ItemID:   444,
				Amount:   1000,
			},
		},
		&EntityInventory{
			SimpleEntityInventory: SimpleEntityInventory{
				EntityID: 555,
				ItemID:   666,
				Amount:   999999,
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
		want    []*EntityInventory
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testEntityInventoryIndex,
			args:    args{campID: 5272, entID: 430214, sync: now},
			want:    invs,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testEntityInventoryIndex,
			args:    args{campID: -123, entID: 430214, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid entID",
			status:  http.StatusOK,
			file:    testEntityInventoryIndex,
			args:    args{campID: 5272, entID: -123, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testEntityInventoryIndex,
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

			got, err := c.EntityInventories.Index(test.args.campID, test.args.entID, test.args.sync)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestEntityInventoryService_Create(t *testing.T) {
	inv := SimpleEntityInventory{
		EntityID: 777,
		ItemID:   888,
		Amount:   1,
	}
	type args struct {
		campID int
		entID  int
		inv    SimpleEntityInventory
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *EntityInventory
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testEntityInventoryCreate,
			args:    args{campID: 5272, entID: 430214, inv: inv},
			want:    &EntityInventory{SimpleEntityInventory: inv},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testEntityInventoryCreate,
			args:    args{campID: -123, entID: 430214, inv: inv},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid entID",
			status:  http.StatusOK,
			file:    testEntityInventoryCreate,
			args:    args{campID: -123, entID: -123, inv: inv},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, empty entity inventory",
			status:  http.StatusOK,
			file:    testEntityInventoryCreate,
			args:    args{campID: 5272, entID: 430214, inv: SimpleEntityInventory{}},
			want:    &EntityInventory{SimpleEntityInventory: inv},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testEntityInventoryCreate,
			args:    args{campID: -123, entID: -123, inv: SimpleEntityInventory{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, inv: inv},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, entID: -123, inv: SimpleEntityInventory{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, inv: inv},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, inv: inv},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, inv: inv},
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

			got, err := c.EntityInventories.Create(test.args.campID, test.args.entID, test.args.inv)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestEntityInventoryService_Update(t *testing.T) {
	inv := SimpleEntityInventory{
		EntityID: 999,
		ItemID:   101010,
		Amount:   1,
	}
	type args struct {
		campID int
		entID  int
		invID  int
		inv    SimpleEntityInventory
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *EntityInventory
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testEntityInventoryUpdate,
			args:    args{campID: 5272, entID: 430214, invID: 111, inv: inv},
			want:    &EntityInventory{SimpleEntityInventory: inv, ID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testEntityInventoryUpdate,
			args:    args{campID: -123, entID: 430214, invID: 111, inv: inv},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid entID",
			status:  http.StatusOK,
			file:    testEntityInventoryUpdate,
			args:    args{campID: -123, entID: -123, invID: 111, inv: inv},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid invID",
			status:  http.StatusOK,
			file:    testEntityInventoryUpdate,
			args:    args{campID: 5272, entID: 430214, invID: -123, inv: inv},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, empty inv",
			status:  http.StatusOK,
			file:    testEntityInventoryUpdate,
			args:    args{campID: 5272, entID: 430214, invID: 111, inv: SimpleEntityInventory{}},
			want:    &EntityInventory{SimpleEntityInventory: inv, ID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testEntityInventoryUpdate,
			args:    args{campID: -123, entID: -123, invID: -123, inv: SimpleEntityInventory{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, invID: 111, inv: inv},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, entID: -123, invID: -123, inv: SimpleEntityInventory{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, invID: 111, inv: inv},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, invID: 111, inv: inv},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, invID: 111, inv: inv},
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

			got, err := c.EntityInventories.Update(test.args.campID, test.args.entID, test.args.invID, test.args.inv)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestEntityInventoryService_Delete(t *testing.T) {
	type args struct {
		campID int
		entID  int
		invID  int
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
			args:    args{campID: 5272, entID: 430214, invID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, invalid campID",
			status:  http.StatusOK,
			args:    args{campID: -123, entID: 430214, invID: 111},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid entID",
			status:  http.StatusOK,
			args:    args{campID: -123, entID: -123, invID: 111},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid invID",
			status:  http.StatusOK,
			args:    args{campID: 5272, entID: 430214, invID: -123},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid args",
			status:  http.StatusOK,
			args:    args{campID: -123, entID: -123, invID: -123},
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			args:    args{campID: 5272, entID: 430214, invID: 111},
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			args:    args{campID: 5272, entID: 430214, invID: 111},
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			args:    args{campID: 5272, entID: 430214, invID: 111},
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

			err = c.EntityInventories.Delete(test.args.campID, test.args.entID, test.args.invID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
		})
	}
}
