package kanka

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

const (
	testEntityNoteIndex  string = "test_data/entitynote_index.json"
	testEntityNoteGet    string = "test_data/entitynote_get.json"
	testEntityNoteCreate string = "test_data/entitynote_create.json"
	testEntityNoteUpdate string = "test_data/entitynote_update.json"
)

func TestEntityNoteService_Index(t *testing.T) {
	notes := []*EntityNote{
		&EntityNote{
			SimpleEntityNote: SimpleEntityNote{
				Name:     "Memories",
				EntityID: 111,
			},
		},
		&EntityNote{
			SimpleEntityNote: SimpleEntityNote{
				Name:     "Secrets",
				EntityID: 222,
			},
		},
		&EntityNote{
			SimpleEntityNote: SimpleEntityNote{
				Name:     "Stashes",
				EntityID: 333,
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
		want    []*EntityNote
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testEntityNoteIndex,
			args:    args{campID: 5272, entID: 430214, sync: now},
			want:    notes,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testEntityNoteIndex,
			args:    args{campID: -123, entID: 430214, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid entID",
			status:  http.StatusOK,
			file:    testEntityNoteIndex,
			args:    args{campID: 5272, entID: -123, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testEntityNoteIndex,
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

			got, err := c.EntityNotes.Index(test.args.campID, test.args.entID, test.args.sync)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestEntityNoteService_Get(t *testing.T) {
	note := &EntityNote{
		SimpleEntityNote: SimpleEntityNote{
			Name:       "Memories",
			EntityID:   430214,
			Entry:      "\n<ul>\n<li>Finding the Wind-up Chronicle</li>\n<li>Losing her older sister</li>\n<li>Creating her first gadget</li>\n</ul>\n",
			IsPrivate:  false,
			Visibility: "all",
		},
		ID:        17762,
		CreatedBy: 5600,
		UpdatedBy: 0,
	}

	type args struct {
		campID int
		entID  int
		noteID int
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *EntityNote
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testEntityNoteGet,
			args:    args{campID: 5272, entID: 430214, noteID: 17762},
			want:    note,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testEntityNoteGet,
			args:    args{campID: -123, entID: 430214, noteID: 17762},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid entID",
			status:  http.StatusOK,
			file:    testEntityNoteGet,
			args:    args{campID: 5272, entID: -123, noteID: 17762},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid noteID",
			status:  http.StatusOK,
			file:    testEntityNoteGet,
			args:    args{campID: 5272, entID: 430214, noteID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testEntityNoteGet,
			args:    args{campID: -123, entID: -123, noteID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, noteID: 17762},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, entID: 430214, noteID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, noteID: 17762},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, noteID: 17762},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, noteID: 17762},
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

			got, err := c.EntityNotes.Get(test.args.campID, test.args.entID, test.args.noteID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestEntityNoteService_Create(t *testing.T) {
	note := SimpleEntityNote{
		Name:     "Hideouts",
		EntityID: 444,
	}
	type args struct {
		campID int
		entID  int
		note   SimpleEntityNote
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *EntityNote
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testEntityNoteCreate,
			args:    args{campID: 5272, entID: 430214, note: note},
			want:    &EntityNote{SimpleEntityNote: note},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testEntityNoteCreate,
			args:    args{campID: -123, entID: 430214, note: note},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid entID",
			status:  http.StatusOK,
			file:    testEntityNoteCreate,
			args:    args{campID: 5272, entID: -123, note: note},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid note",
			status:  http.StatusOK,
			file:    testEntityNoteCreate,
			args:    args{campID: 5272, entID: 430214, note: SimpleEntityNote{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testEntityNoteCreate,
			args:    args{campID: -123, entID: -123, note: SimpleEntityNote{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, note: note},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, entID: -123, note: SimpleEntityNote{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, note: note},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, note: note},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, note: note},
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

			got, err := c.EntityNotes.Create(test.args.campID, test.args.entID, test.args.note)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestEntityNoteService_Update(t *testing.T) {
	note := SimpleEntityNote{
		Name:     "Letters",
		EntityID: 555,
	}
	type args struct {
		campID int
		entID  int
		noteID int
		note   SimpleEntityNote
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *EntityNote
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testEntityNoteUpdate,
			args:    args{campID: 5272, entID: 430214, noteID: 111, note: note},
			want:    &EntityNote{SimpleEntityNote: note, ID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testEntityNoteUpdate,
			args:    args{campID: -123, entID: 430214, noteID: 111, note: note},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid entID",
			status:  http.StatusOK,
			file:    testEntityNoteUpdate,
			args:    args{campID: 5272, entID: -123, noteID: 111, note: note},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid noteID",
			status:  http.StatusOK,
			file:    testEntityNoteUpdate,
			args:    args{campID: 5272, entID: 430214, noteID: -123, note: note},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid note",
			status:  http.StatusOK,
			file:    testEntityNoteUpdate,
			args:    args{campID: 5272, entID: 430214, noteID: 111, note: SimpleEntityNote{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testEntityNoteUpdate,
			args:    args{campID: -123, entID: -123, noteID: -123, note: SimpleEntityNote{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, noteID: 111, note: note},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, entID: -123, noteID: -123, note: SimpleEntityNote{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, noteID: 111, note: note},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, noteID: 111, note: note},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, entID: 430214, noteID: 111, note: note},
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

			got, err := c.EntityNotes.Update(test.args.campID, test.args.entID, test.args.noteID, test.args.note)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestEntityNoteService_Delete(t *testing.T) {
	type args struct {
		campID int
		entID  int
		noteID int
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
			args:    args{campID: 5272, entID: 430214, noteID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, invalid campID",
			status:  http.StatusOK,
			args:    args{campID: -123, entID: 430214, noteID: 111},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid entID",
			status:  http.StatusOK,
			args:    args{campID: -123, entID: -123, noteID: 111},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid noteID",
			status:  http.StatusOK,
			args:    args{campID: 5272, entID: 430214, noteID: -123},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid args",
			status:  http.StatusOK,
			args:    args{campID: -123, entID: -123, noteID: -123},
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			args:    args{campID: 5272, entID: 430214, noteID: 111},
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			args:    args{campID: 5272, entID: 430214, noteID: 111},
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			args:    args{campID: 5272, entID: 430214, noteID: 111},
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

			err = c.EntityNotes.Delete(test.args.campID, test.args.entID, test.args.noteID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
		})
	}
}
