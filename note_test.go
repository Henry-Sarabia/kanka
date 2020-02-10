package kanka

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

const (
	testNoteIndex  string = "test_data/note_index.json"
	testNoteGet    string = "test_data/note_get.json"
	testNoteCreate string = "test_data/note_create.json"
	testNoteUpdate string = "test_data/note_update.json"
)

func TestNoteService_Index(t *testing.T) {
	notes := []*Note{
		{
			SimpleNote: SimpleNote{
				Name: "To Emilia",
				Type: "Letter",
			},
		},
		{
			SimpleNote: SimpleNote{
				Name: "A Private Diary",
				Type: "Diary",
			},
		},
		{
			SimpleNote: SimpleNote{
				Name: "The Surrounding Verdant Forest: A Guide",
				Type: "Book",
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
		want    []*Note
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testNoteIndex,
			args:    args{campID: 5272, sync: now},
			want:    notes,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testNoteIndex,
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

			got, err := c.Notes.Index(test.args.campID, test.args.sync)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestNoteService_Get(t *testing.T) {
	note := &Note{
		SimpleNote: SimpleNote{
			Name:      "Emberweed",
			Entry:     "\n<p>A very tasty plant whose seeds taste of pepper, Emberweed possesses a hidden danger. Its famous seeds, used often in Dwarven and Gnomish cooking, are incredibly flammable and will explode if left too long near a fire.</p>\n",
			Image:     "notes/2XAeetUHrO0TDB0lU6tAsa46wvdIEmZ1jmsSUC20.jpeg",
			IsPrivate: false,
			Tags:      []int{3742},
			Type:      "Worldbuilding",
		},
		ID:             2142,
		ImageFull:      "https://kanka-user-assets.s3.eu-central-1.amazonaws.com/notes/2XAeetUHrO0TDB0lU6tAsa46wvdIEmZ1jmsSUC20.jpeg",
		ImageThumb:     "https://kanka-user-assets.s3.eu-central-1.amazonaws.com/notes/2XAeetUHrO0TDB0lU6tAsa46wvdIEmZ1jmsSUC20_thumb.jpeg",
		HasCustomImage: true,
		EntityID:       86314,
		CreatedBy:      5600,
		UpdatedBy:      5600,
	}

	type args struct {
		campID int
		noteID int
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *Note
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testNoteGet,
			args:    args{campID: 5272, noteID: 2142},
			want:    note,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testNoteGet,
			args:    args{campID: -123, noteID: 2142},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid noteID",
			status:  http.StatusOK,
			file:    testNoteGet,
			args:    args{campID: 5272, noteID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testNoteGet,
			args:    args{campID: -123, noteID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, noteID: 2142},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, noteID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, noteID: 2142},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, noteID: 2142},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, noteID: 2142},
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

			got, err := c.Notes.Get(test.args.campID, test.args.noteID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestNoteService_Create(t *testing.T) {
	note := SimpleNote{
		Name: "A Burned Letter",
		Type: "Letter",
	}
	type args struct {
		campID int
		note   SimpleNote
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *Note
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testNoteCreate,
			args:    args{campID: 5272, note: note},
			want:    &Note{SimpleNote: note},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testNoteCreate,
			args:    args{campID: -123, note: note},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid note",
			status:  http.StatusOK,
			file:    testNoteCreate,
			args:    args{campID: 5272, note: SimpleNote{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testNoteCreate,
			args:    args{campID: -123, note: SimpleNote{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, note: note},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, note: SimpleNote{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, note: note},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, note: note},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, note: note},
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

			got, err := c.Notes.Create(test.args.campID, test.args.note)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestNoteService_Update(t *testing.T) {
	note := SimpleNote{
		Name: "Xanathar's Spellbook",
		Type: "Spellbook",
	}
	type args struct {
		campID int
		noteID int
		note   SimpleNote
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *Note
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testNoteUpdate,
			args:    args{campID: 5272, noteID: 111, note: note},
			want:    &Note{SimpleNote: note, ID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testNoteUpdate,
			args:    args{campID: -123, noteID: 111, note: note},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid noteID",
			status:  http.StatusOK,
			file:    testNoteUpdate,
			args:    args{campID: 5272, noteID: -123, note: note},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid note",
			status:  http.StatusOK,
			file:    testNoteUpdate,
			args:    args{campID: 5272, noteID: 111, note: SimpleNote{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testNoteUpdate,
			args:    args{campID: -123, noteID: -123, note: SimpleNote{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, noteID: 111, note: note},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, noteID: -123, note: SimpleNote{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, noteID: 111, note: note},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, noteID: 111, note: note},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, noteID: 111, note: note},
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

			got, err := c.Notes.Update(test.args.campID, test.args.noteID, test.args.note)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestNoteService_Delete(t *testing.T) {
	type args struct {
		campID int
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
			args:    args{campID: 5272, noteID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, invalid campID",
			status:  http.StatusOK,
			args:    args{campID: -123, noteID: 111},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid noteID",
			status:  http.StatusOK,
			args:    args{campID: 5272, noteID: -123},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid args",
			status:  http.StatusOK,
			args:    args{campID: -123, noteID: -123},
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			args:    args{campID: 5272, noteID: 111},
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			args:    args{campID: 5272, noteID: 111},
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			args:    args{campID: 5272, noteID: 111},
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

			err = c.Notes.Delete(test.args.campID, test.args.noteID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
		})
	}
}
