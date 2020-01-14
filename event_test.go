package kanka

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

const (
	testEventIndex  string = "test_data/event_index.json"
	testEventGet    string = "test_data/event_get.json"
	testEventCreate string = "test_data/event_create.json"
	testEventUpdate string = "test_data/event_update.json"
)

func TestEventService_Index(t *testing.T) {
	evts := []*Event{
		&Event{
			SimpleEvent: SimpleEvent{
				Name: "Bloom's Peak",
				Type: "Festival",
			},
		},
		&Event{
			SimpleEvent: SimpleEvent{
				Name: "Pelor's Day",
				Type: "Holy Day",
			},
		},
		&Event{
			SimpleEvent: SimpleEvent{
				Name: "Mt. Dooom Erupts",
				Type: "Natural Disaster",
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
		want    []*Event
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testEventIndex,
			args:    args{campID: 5272, sync: now},
			want:    evts,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testEventIndex,
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

			got, err := c.Events.Index(test.args.campID, test.args.sync)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestEventService_Get(t *testing.T) {
	evt := &Event{
		SimpleEvent: SimpleEvent{
			Name:       "Winter's Crest",
			Entry:      "\n<p>A fantastic festival!</p>\n",
			Image:      "events/pKtVTVLA776DKxNq3PMquVmNGmg3rJVboEJcje8j.jpeg",
			IsPrivate:  false,
			Tags:       []int{35131},
			LocationID: 115366,
			Type:       "Festival",
		},
		ID:             15008,
		ImageFull:      "https://kanka-user-assets.s3.eu-central-1.amazonaws.com/events/pKtVTVLA776DKxNq3PMquVmNGmg3rJVboEJcje8j.jpeg",
		ImageThumb:     "https://kanka-user-assets.s3.eu-central-1.amazonaws.com/events/pKtVTVLA776DKxNq3PMquVmNGmg3rJVboEJcje8j_thumb.jpeg",
		HasCustomImage: true,
		EntityID:       442245,
		CreatedBy:      5600,
		UpdatedBy:      5600,
	}

	type args struct {
		campID int
		evtID  int
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *Event
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testEventGet,
			args:    args{campID: 5272, evtID: 15008},
			want:    evt,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testEventGet,
			args:    args{campID: -123, evtID: 15008},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid evtID",
			status:  http.StatusOK,
			file:    testEventGet,
			args:    args{campID: 5272, evtID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testEventGet,
			args:    args{campID: -123, evtID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, evtID: 15008},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, evtID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, evtID: 15008},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, evtID: 15008},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, evtID: 15008},
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

			got, err := c.Events.Get(test.args.campID, test.args.evtID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestEventService_Create(t *testing.T) {
	evt := SimpleEvent{
		Name: "Ember's Peak",
		Type: "Festival",
	}
	type args struct {
		campID int
		evt    SimpleEvent
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *Event
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testEventCreate,
			args:    args{campID: 5272, evt: evt},
			want:    &Event{SimpleEvent: evt},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testEventCreate,
			args:    args{campID: -123, evt: evt},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid event",
			status:  http.StatusOK,
			file:    testEventCreate,
			args:    args{campID: 5272, evt: SimpleEvent{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testEventCreate,
			args:    args{campID: -123, evt: SimpleEvent{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, evt: evt},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, evt: SimpleEvent{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, evt: evt},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, evt: evt},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, evt: evt},
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

			got, err := c.Events.Create(test.args.campID, test.args.evt)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestEventService_Update(t *testing.T) {
	evt := SimpleEvent{
		Name: "Lilit's Day",
		Type: "Holy Day",
	}
	type args struct {
		campID int
		evtID  int
		evt    SimpleEvent
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *Event
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testEventUpdate,
			args:    args{campID: 5272, evtID: 111, evt: evt},
			want:    &Event{SimpleEvent: evt, ID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testEventUpdate,
			args:    args{campID: -123, evtID: 111, evt: evt},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid evtID",
			status:  http.StatusOK,
			file:    testEventUpdate,
			args:    args{campID: 5272, evtID: -123, evt: evt},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid evt",
			status:  http.StatusOK,
			file:    testEventUpdate,
			args:    args{campID: 5272, evtID: 111, evt: SimpleEvent{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testEventUpdate,
			args:    args{campID: -123, evtID: -123, evt: SimpleEvent{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, evtID: 111, evt: evt},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, evtID: -123, evt: SimpleEvent{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, evtID: 111, evt: evt},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, evtID: 111, evt: evt},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, evtID: 111, evt: evt},
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

			got, err := c.Events.Update(test.args.campID, test.args.evtID, test.args.evt)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestEventService_Delete(t *testing.T) {
	type args struct {
		campID int
		evtID  int
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
			args:    args{campID: 5272, evtID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, invalid campID",
			status:  http.StatusOK,
			args:    args{campID: -123, evtID: 111},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid evtID",
			status:  http.StatusOK,
			args:    args{campID: 5272, evtID: -123},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid args",
			status:  http.StatusOK,
			args:    args{campID: -123, evtID: -123},
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			args:    args{campID: 5272, evtID: 111},
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			args:    args{campID: 5272, evtID: 111},
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			args:    args{campID: 5272, evtID: 111},
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

			err = c.Events.Delete(test.args.campID, test.args.evtID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
		})
	}
}
