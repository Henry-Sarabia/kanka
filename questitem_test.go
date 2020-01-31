package kanka

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

const (
	testQuestItemIndex  string = "test_data/questitem_index.json"
	testQuestItemGet    string = "test_data/questitem_get.json"
	testQuestItemCreate string = "test_data/questitem_create.json"
	testQuestItemUpdate string = "test_data/questitem_update.json"
)

func TestQuestItemService_Index(t *testing.T) {
	items := []*QuestItem{
		&QuestItem{
			SimpleQuestItem: SimpleQuestItem{
				QuestID:     111,
				ItemID:      222,
				Description: "A crystal flask with a glowing blue liquid; golden flakes float throughout.",
			},
		},
		&QuestItem{
			SimpleQuestItem: SimpleQuestItem{
				QuestID:     333,
				ItemID:      444,
				Description: "A towering bronze shield emblazoned with a silver starburst.",
			},
		},
		&QuestItem{
			SimpleQuestItem: SimpleQuestItem{
				QuestID:     555,
				ItemID:      666,
				Description: "A wooden gnarled staff adorned with a massive emerald encased in ivy.",
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
		want    []*QuestItem
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testQuestItemIndex,
			args:    args{campID: 5272, qstID: 10394, sync: now},
			want:    items,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testQuestItemIndex,
			args:    args{campID: -123, qstID: 10394, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid qstID",
			status:  http.StatusOK,
			file:    testQuestItemIndex,
			args:    args{campID: 5272, qstID: -123, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testQuestItemIndex,
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

			got, err := c.QuestItems.Index(test.args.campID, test.args.qstID, test.args.sync)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestQuestItemService_Get(t *testing.T) {
	item := &QuestItem{
		SimpleQuestItem: SimpleQuestItem{
			ItemID:      123,
			Description: "<!DOCTYPE html>\r\n<html>\r\n<head>\r\n</head>\r\n<body>\r\n<p>Priceless item</p>\r\n</body>\r\n</html>",
			Role:        "Reward",
			IsPrivate:   false,
		},
		ID:        445,
		CreatedBy: 0,
		UpdatedBy: 0,
	}

	type args struct {
		campID int
		qstID  int
		itemID int
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *QuestItem
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testQuestItemGet,
			args:    args{campID: 5272, qstID: 10394, itemID: 445},
			want:    item,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testQuestItemGet,
			args:    args{campID: -123, qstID: 10394, itemID: 445},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid qstID",
			status:  http.StatusOK,
			file:    testQuestItemGet,
			args:    args{campID: 5272, qstID: -123, itemID: 445},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid itemID",
			status:  http.StatusOK,
			file:    testQuestItemGet,
			args:    args{campID: 5272, qstID: 10394, itemID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testQuestItemGet,
			args:    args{campID: -123, qstID: -123, itemID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, itemID: 445},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, qstID: 10394, itemID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, itemID: 445},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, itemID: 445},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, itemID: 445},
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

			got, err := c.QuestItems.Get(test.args.campID, test.args.qstID, test.args.itemID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestQuestItemService_Create(t *testing.T) {
	item := SimpleQuestItem{
		QuestID:     777,
		ItemID:      888,
		Description: "A crimson steel pin in the shape of a burning heart.",
	}
	type args struct {
		campID int
		qstID  int
		item   SimpleQuestItem
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *QuestItem
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testQuestItemCreate,
			args:    args{campID: 5272, qstID: 10394, item: item},
			want:    &QuestItem{SimpleQuestItem: item},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testQuestItemCreate,
			args:    args{campID: -123, qstID: 10394, item: item},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid qstID",
			status:  http.StatusOK,
			file:    testQuestItemCreate,
			args:    args{campID: 5272, qstID: -123, item: item},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, empty item",
			status:  http.StatusOK,
			file:    testQuestItemCreate,
			args:    args{campID: 5272, qstID: 10394, item: SimpleQuestItem{}},
			want:    &QuestItem{SimpleQuestItem: item},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testQuestItemCreate,
			args:    args{campID: -123, qstID: -123, item: SimpleQuestItem{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, item: item},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, qstID: -123, item: SimpleQuestItem{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, item: item},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, item: item},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, item: item},
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

			got, err := c.QuestItems.Create(test.args.campID, test.args.qstID, test.args.item)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestQuestItemService_Update(t *testing.T) {
	item := SimpleQuestItem{
		QuestID:     999,
		ItemID:      101010,
		Description: "A bolt of dark fabric that seems to drink the surrounding light.",
	}
	type args struct {
		campID int
		qstID  int
		itemID int
		item   SimpleQuestItem
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *QuestItem
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testQuestItemUpdate,
			args:    args{campID: 5272, qstID: 10394, itemID: 111, item: item},
			want:    &QuestItem{SimpleQuestItem: item, ID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testQuestItemUpdate,
			args:    args{campID: -123, qstID: 10394, itemID: 111, item: item},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid qstID",
			status:  http.StatusOK,
			file:    testQuestItemUpdate,
			args:    args{campID: 5272, qstID: -123, itemID: 111, item: item},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid itemID",
			status:  http.StatusOK,
			file:    testQuestItemUpdate,
			args:    args{campID: 5272, qstID: 10394, itemID: -123, item: item},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, empty item",
			status:  http.StatusOK,
			file:    testQuestItemUpdate,
			args:    args{campID: 5272, qstID: 10394, itemID: 111, item: SimpleQuestItem{}},
			want:    &QuestItem{SimpleQuestItem: item, ID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testQuestItemUpdate,
			args:    args{campID: -123, qstID: -123, itemID: -123, item: SimpleQuestItem{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, itemID: 111, item: item},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, qstID: -123, itemID: -123, item: SimpleQuestItem{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, itemID: 111, item: item},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, itemID: 111, item: item},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, itemID: 111, item: item},
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

			got, err := c.QuestItems.Update(test.args.campID, test.args.qstID, test.args.itemID, test.args.item)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestQuestItemService_Delete(t *testing.T) {
	type args struct {
		campID int
		qstID  int
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
			args:    args{campID: 5272, qstID: 10394, itemID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, invalid campID",
			status:  http.StatusOK,
			args:    args{campID: -123, qstID: 10394, itemID: 111},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid qstID",
			status:  http.StatusOK,
			args:    args{campID: 5272, qstID: -123, itemID: 111},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid itemID",
			status:  http.StatusOK,
			args:    args{campID: 5272, qstID: 10394, itemID: -123},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid args",
			status:  http.StatusOK,
			args:    args{campID: -123, qstID: -123, itemID: -123},
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			args:    args{campID: 5272, qstID: 10394, itemID: 111},
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			args:    args{campID: 5272, qstID: 10394, itemID: 111},
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			args:    args{campID: 5272, qstID: 10394, itemID: 111},
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

			err = c.QuestItems.Delete(test.args.campID, test.args.qstID, test.args.itemID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
		})
	}
}
