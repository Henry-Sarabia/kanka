package kanka

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

const (
	testEntityEventIndex  string = "test_data/entityevent_index.json"
	testEntityEventGet    string = "test_data/entityevent_get.json"
	testEntityEventCreate string = "test_data/entityevent_create.json"
	testEntityEventUpdate string = "test_data/entityevent_update.json"
)

func TestEntityEventService_Index(t *testing.T) {
	evts := []*EntityEvent{
		{
			SimpleEntityEvent: SimpleEntityEvent{
				Comment: "Battle of Halefort",
				Year:    745,
			},
		},
		{
			SimpleEntityEvent: SimpleEntityEvent{
				Comment: "Razing of Shale's End",
				Year:    728,
			},
		},
		{
			SimpleEntityEvent: SimpleEntityEvent{
				Comment: "Last stand at Halefort",
				Year:    747,
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
		want    []*EntityEvent
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testEntityEventIndex,
			args:    args{campID: 5272, entID: 430214, sync: now},
			want:    evts,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testEntityEventIndex,
			args:    args{campID: -123, entID: 430214, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid entID",
			status:  http.StatusOK,
			file:    testEntityEventIndex,
			args:    args{campID: 5272, entID: -123, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testEntityEventIndex,
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

			got, err := c.EntityEvents.Index(test.args.campID, test.args.entID, test.args.sync)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestEntityEventService_Get(t *testing.T) {
	evt := &EntityEvent{
		SimpleEntityEvent: SimpleEntityEvent{
			Colour:         "light-blue",
			Comment:        "Birthday",
			Day:            1,
			EntityID:       430214,
			IsPrivate:      false,
			IsRecurring:    true,
			Length:         1,
			Month:          7,
			RecurringUntil: 0,
			Year:           776,
		},
		CalendarID: 436,
		Date:       "",
		ID:         17492,
		CreatedBy:  5600,
		UpdatedBy:  5600,
	}

	type args struct {
		campID int
		entID  int
		evtID  int
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *EntityEvent
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testEntityEventGet,
			args:    args{campID: 5272, entID: 430214, evtID: 17492},
			want:    evt,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testEntityEventGet,
			args:    args{campID: -123, entID: 430214, evtID: 17492},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid entID",
			status:  http.StatusOK,
			file:    testEntityEventGet,
			args:    args{campID: 5272, entID: -123, evtID: 17492},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid evtID",
			status:  http.StatusOK,
			file:    testEntityEventGet,
			args:    args{campID: 5272, entID: 430214, evtID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testEntityEventGet,
			args:    args{campID: -123, entID: -123, evtID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, evtID: 17492},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, entID: 430214, evtID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, evtID: 17492},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, evtID: 17492},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, evtID: 17492},
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

			got, err := c.EntityEvents.Get(test.args.campID, test.args.entID, test.args.evtID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestEntityEventService_Create(t *testing.T) {
	evt := SimpleEntityEvent{
		Comment: "Invasion of Northwind",
		Year:    689,
	}
	type args struct {
		campID int
		entID  int
		evt    SimpleEntityEvent
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *EntityEvent
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testEntityEventCreate,
			args:    args{campID: 5272, entID: 430214, evt: evt},
			want:    &EntityEvent{SimpleEntityEvent: evt},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testEntityEventCreate,
			args:    args{campID: -123, entID: 430214, evt: evt},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid entID",
			status:  http.StatusOK,
			file:    testEntityEventCreate,
			args:    args{campID: 5272, entID: -123, evt: evt},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, empty entityevent",
			status:  http.StatusOK,
			file:    testEntityEventCreate,
			args:    args{campID: 5272, entID: 430214, evt: SimpleEntityEvent{}},
			want:    &EntityEvent{SimpleEntityEvent: evt},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testEntityEventCreate,
			args:    args{campID: -123, entID: -123, evt: SimpleEntityEvent{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, evt: evt},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, entID: -123, evt: SimpleEntityEvent{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, evt: evt},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, evt: evt},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, evt: evt},
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

			got, err := c.EntityEvents.Create(test.args.campID, test.args.entID, test.args.evt)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestEntityEventService_Update(t *testing.T) {
	evt := SimpleEntityEvent{
		Comment: "Fall of Malthen",
		Year:    699,
	}
	type args struct {
		campID int
		entID  int
		evtID  int
		evt    SimpleEntityEvent
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *EntityEvent
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testEntityEventUpdate,
			args:    args{campID: 5272, entID: 430214, evtID: 111, evt: evt},
			want:    &EntityEvent{SimpleEntityEvent: evt, ID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testEntityEventUpdate,
			args:    args{campID: -123, entID: 430214, evtID: 111, evt: evt},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid entID",
			status:  http.StatusOK,
			file:    testEntityEventUpdate,
			args:    args{campID: 5272, entID: -123, evtID: 111, evt: evt},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid evtID",
			status:  http.StatusOK,
			file:    testEntityEventUpdate,
			args:    args{campID: 5272, entID: 430214, evtID: -123, evt: evt},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, empty evt",
			status:  http.StatusOK,
			file:    testEntityEventUpdate,
			args:    args{campID: 5272, entID: 430214, evtID: 111, evt: SimpleEntityEvent{}},
			want:    &EntityEvent{SimpleEntityEvent: evt, ID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testEntityEventUpdate,
			args:    args{campID: -123, entID: -123, evtID: -123, evt: SimpleEntityEvent{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, evtID: 111, evt: evt},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, entID: -123, evtID: -123, evt: SimpleEntityEvent{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, evtID: 111, evt: evt},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, evtID: 111, evt: evt},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, evtID: 111, evt: evt},
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

			got, err := c.EntityEvents.Update(test.args.campID, test.args.entID, test.args.evtID, test.args.evt)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestEntityEventService_Delete(t *testing.T) {
	type args struct {
		campID int
		entID  int
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
			args:    args{campID: 5272, entID: 430214, evtID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, invalid campID",
			status:  http.StatusOK,
			args:    args{campID: -123, entID: 430214, evtID: 111},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid entID",
			status:  http.StatusOK,
			args:    args{campID: -123, entID: -123, evtID: 111},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid evtID",
			status:  http.StatusOK,
			args:    args{campID: 5272, entID: 430214, evtID: -123},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid args",
			status:  http.StatusOK,
			args:    args{campID: -123, entID: -123, evtID: -123},
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			args:    args{campID: 5272, entID: 430214, evtID: 111},
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			args:    args{campID: 5272, entID: 430214, evtID: 111},
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			args:    args{campID: 5272, entID: 430214, evtID: 111},
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

			err = c.EntityEvents.Delete(test.args.campID, test.args.entID, test.args.evtID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
		})
	}
}
