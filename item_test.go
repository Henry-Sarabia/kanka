package kanka

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

const (
	testItemIndex  string = "test_data/item_index.json"
	testItemGet    string = "test_data/item_get.json"
	testItemCreate string = "test_data/item_create.json"
	testItemUpdate string = "test_data/item_update.json"
)

func TestItemService_Index(t *testing.T) {
	items := []*Item{
		&Item{
			SimpleItem: SimpleItem{
				Name: "Sword",
				Type: "Weapon",
			},
		},
		&Item{
			SimpleItem: SimpleItem{
				Name: "Bow",
				Type: "Weapon",
			},
		},
		&Item{
			SimpleItem: SimpleItem{
				Name: "Saddlebag",
				Type: "Equipment",
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
		want    []*Item
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testItemIndex,
			args:    args{campID: 5272, sync: now},
			want:    items,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testItemIndex,
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

			got, err := c.Items.Index(test.args.campID, test.args.sync)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestItemService_Get(t *testing.T) {
	item := &Item{
		SimpleItem: SimpleItem{
			Name:        "Wind-up Chronicle",
			Entry:       "\n<p>Time dilation</p>\n",
			Image:       "items/vyh5Fjzeij0gtdgsqjlJJOyTO4it6LG5l3ajTDNA.jpeg",
			IsPrivate:   false,
			Tags:        []int{34696},
			LocationID:  111,
			CharacterID: 222,
			Type:        "Timepiece",
			Price:       "500",
			Size:        "Diminutive",
		},
		ID:             34780,
		ImageFull:      "https://kanka-user-assets.s3.eu-central-1.amazonaws.com/items/vyh5Fjzeij0gtdgsqjlJJOyTO4it6LG5l3ajTDNA.jpeg",
		ImageThumb:     "https://kanka-user-assets.s3.eu-central-1.amazonaws.com/items/vyh5Fjzeij0gtdgsqjlJJOyTO4it6LG5l3ajTDNA_thumb.jpeg",
		HasCustomImage: true,
		EntityID:       434400,
		CreatedBy:      5600,
		UpdatedBy:      5600,
	}

	type args struct {
		campID int
		itemID int
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *Item
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testItemGet,
			args:    args{campID: 5272, itemID: 34780},
			want:    item,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testItemGet,
			args:    args{campID: -123, itemID: 34780},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid itemID",
			status:  http.StatusOK,
			file:    testItemGet,
			args:    args{campID: 5272, itemID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testItemGet,
			args:    args{campID: -123, itemID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, itemID: 34780},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, itemID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, itemID: 34780},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, itemID: 34780},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, itemID: 34780},
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

			got, err := c.Items.Get(test.args.campID, test.args.itemID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestItemService_Create(t *testing.T) {
	item := SimpleItem{
		Name: "Torch",
		Type: "Equipment",
	}
	type args struct {
		campID int
		item   SimpleItem
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *Item
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testItemCreate,
			args:    args{campID: 5272, item: item},
			want:    &Item{SimpleItem: item},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testItemCreate,
			args:    args{campID: -123, item: item},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid item",
			status:  http.StatusOK,
			file:    testItemCreate,
			args:    args{campID: 5272, item: SimpleItem{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testItemCreate,
			args:    args{campID: -123, item: SimpleItem{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, item: item},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, item: SimpleItem{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, item: item},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, item: item},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, item: item},
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

			got, err := c.Items.Create(test.args.campID, test.args.item)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestItemService_Update(t *testing.T) {
	item := SimpleItem{
		Name: "Rope",
		Type: "Equipment",
	}
	type args struct {
		campID int
		itemID int
		item   SimpleItem
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *Item
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testItemUpdate,
			args:    args{campID: 5272, itemID: 111, item: item},
			want:    &Item{SimpleItem: item, ID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testItemUpdate,
			args:    args{campID: -123, itemID: 111, item: item},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid itemID",
			status:  http.StatusOK,
			file:    testItemUpdate,
			args:    args{campID: 5272, itemID: -123, item: item},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid item",
			status:  http.StatusOK,
			file:    testItemUpdate,
			args:    args{campID: 5272, itemID: 111, item: SimpleItem{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testItemUpdate,
			args:    args{campID: -123, itemID: -123, item: SimpleItem{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, itemID: 111, item: item},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, itemID: -123, item: SimpleItem{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, itemID: 111, item: item},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, itemID: 111, item: item},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, itemID: 111, item: item},
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

			got, err := c.Items.Update(test.args.campID, test.args.itemID, test.args.item)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestItemService_Delete(t *testing.T) {
	type args struct {
		campID int
		itemID int
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
			args:    args{campID: 5272, itemID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, invalid campID",
			status:  http.StatusOK,
			args:    args{campID: -123, itemID: 111},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid itemID",
			status:  http.StatusOK,
			args:    args{campID: 5272, itemID: -123},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid args",
			status:  http.StatusOK,
			args:    args{campID: -123, itemID: -123},
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			args:    args{campID: 5272, itemID: 111},
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			args:    args{campID: 5272, itemID: 111},
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			args:    args{campID: 5272, itemID: 111},
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

			err = c.Items.Delete(test.args.campID, test.args.itemID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
		})
	}
}
