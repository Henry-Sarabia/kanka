package kanka

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

const (
	testJournalIndex  string = "test_data/journal_index.json"
	testJournalGet    string = "test_data/journal_get.json"
	testJournalCreate string = "test_data/journal_create.json"
	testJournalUpdate string = "test_data/journal_update.json"
)

func TestJournalService_Index(t *testing.T) {
	jrns := []*Journal{
		{
			SimpleJournal: SimpleJournal{
				Name: "How I Learned to Stop Worrying and Love The Eruption",
				Type: "Treatise",
			},
		},
		{
			SimpleJournal: SimpleJournal{
				Name: "What I Think About When I Think About Spiders",
				Type: "Treatise",
			},
		},
		{
			SimpleJournal: SimpleJournal{
				Name: "Easier Done Than Said",
				Type: "Treatise",
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
		want    []*Journal
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testJournalIndex,
			args:    args{campID: 5272, sync: now},
			want:    jrns,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testJournalIndex,
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

			got, err := c.Journals.Index(test.args.campID, test.args.sync)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestJournalService_Get(t *testing.T) {
	jrn := &Journal{
		SimpleJournal: SimpleJournal{
			Name:        "A Dissertation on the Shortcomings of Dead Reckoning",
			Entry:       "\n<div>\n<p>A small book with an unremarkable beryl-colored cover. The book is titled \"A Dissertation on the Shortcomings of Dead Reckoning\". The author is Alec Grisbane.</p>\n<p>The book is a commentary on the current state of nautical navigation and exploration. The author proposes that dead reckoning is inaccurate, inefficient, and ultimately dangerous.</p>\n</div>\n<p>The introduction begins with the supposition that dead reckoning should be done away with completely in favor of a new method of navigation involving accurate timekeeping, nautical velocity measurements, and a knowledge of astronomy.</p>\n<p>The author spends the next few pages explaining why dead reckoning is the primary form of nautical navigation in the modern era and how it is generally accomplished.</p>\n<p>The author then spends some time going into detail about the myriad methods of timekeeping at sea. Of the numerous approaches, none are accurate enough for meaningful use, he argues. This section of the essay is lengthier than one would expect.</p>\n<p>The author continues with details concerning the main method of velocity measurement and, more specifically, the inaccuracies it can produce. He posits that the lack of alternative methods is due to the corresponding lack of ingenuity and education of maritime explorers.</p>\n<p>The author digresses briefly to discuss the imprecision of maritime cartographers and their share of the blame for dead reckoning's proliferation.</p>\n<p>Subsequently, the author returns to his main point and expounds on the culmination of the aforementioned inadequacies in regards to the mechanisms for timekeeping, velocity measurement, and nautical cartography. The sum of these parts, he argues, establishes a flawed system putting maritime explorers and traders at grave risk when setting out to sea. What's worse, he continues, is the fundamental lack of options they have available to them if they are fortuitous enough, or unfortuitous enough depending on your worldview, to recognize their miscalculation while in the open ocean.</p>\n<p>The author closes with a cursory summation of his argument.</p>\n<p>Notably, the author does not propose a material alternative to dead reckoning.</p>\n",
			Image:       "journals/ZpHh9IXq79rbVojliMS2nJu4Rm7yfpxFRnVQ6tG8.jpeg",
			IsPrivate:   false,
			Tags:        []int{},
			LocationID:  25989,
			CharacterID: 24657,
			Date:        "2020-01-15",
			Type:        "Treatise",
		},
		ID:             1142,
		ImageFull:      "https://kanka-user-assets.s3.eu-central-1.amazonaws.com/journals/ZpHh9IXq79rbVojliMS2nJu4Rm7yfpxFRnVQ6tG8.jpeg",
		ImageThumb:     "https://kanka-user-assets.s3.eu-central-1.amazonaws.com/journals/ZpHh9IXq79rbVojliMS2nJu4Rm7yfpxFRnVQ6tG8_thumb.jpeg",
		HasCustomImage: true,
		EntityID:       80946,
		CreatedBy:      5600,
		UpdatedBy:      5600,
	}

	type args struct {
		campID int
		jrnID  int
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *Journal
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testJournalGet,
			args:    args{campID: 5272, jrnID: 1142},
			want:    jrn,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testJournalGet,
			args:    args{campID: -123, jrnID: 1142},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid jrnID",
			status:  http.StatusOK,
			file:    testJournalGet,
			args:    args{campID: 5272, jrnID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testJournalGet,
			args:    args{campID: -123, jrnID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, jrnID: 1142},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, jrnID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, jrnID: 1142},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, jrnID: 1142},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, jrnID: 1142},
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

			got, err := c.Journals.Get(test.args.campID, test.args.jrnID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestJournalService_Create(t *testing.T) {
	jrn := SimpleJournal{
		Name: "Everyone I Know Is Either Dead Or Dying",
		Type: "Novel",
	}
	type args struct {
		campID int
		jrn    SimpleJournal
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *Journal
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testJournalCreate,
			args:    args{campID: 5272, jrn: jrn},
			want:    &Journal{SimpleJournal: jrn},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testJournalCreate,
			args:    args{campID: -123, jrn: jrn},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid journal",
			status:  http.StatusOK,
			file:    testJournalCreate,
			args:    args{campID: 5272, jrn: SimpleJournal{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testJournalCreate,
			args:    args{campID: -123, jrn: SimpleJournal{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, jrn: jrn},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, jrn: SimpleJournal{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, jrn: jrn},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, jrn: jrn},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, jrn: jrn},
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

			got, err := c.Journals.Create(test.args.campID, test.args.jrn)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestJournalService_Update(t *testing.T) {
	jrn := SimpleJournal{
		Name: "Stop Making Me Happy, You're Making It Worse",
		Type: "Novel",
	}
	type args struct {
		campID int
		jrnID  int
		jrn    SimpleJournal
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *Journal
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testJournalUpdate,
			args:    args{campID: 5272, jrnID: 111, jrn: jrn},
			want:    &Journal{SimpleJournal: jrn, ID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testJournalUpdate,
			args:    args{campID: -123, jrnID: 111, jrn: jrn},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid jrnID",
			status:  http.StatusOK,
			file:    testJournalUpdate,
			args:    args{campID: 5272, jrnID: -123, jrn: jrn},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid jrn",
			status:  http.StatusOK,
			file:    testJournalUpdate,
			args:    args{campID: 5272, jrnID: 111, jrn: SimpleJournal{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testJournalUpdate,
			args:    args{campID: -123, jrnID: -123, jrn: SimpleJournal{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, jrnID: 111, jrn: jrn},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, jrnID: -123, jrn: SimpleJournal{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, jrnID: 111, jrn: jrn},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, jrnID: 111, jrn: jrn},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, jrnID: 111, jrn: jrn},
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

			got, err := c.Journals.Update(test.args.campID, test.args.jrnID, test.args.jrn)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestJournalService_Delete(t *testing.T) {
	type args struct {
		campID int
		jrnID  int
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
			args:    args{campID: 5272, jrnID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, invalid campID",
			status:  http.StatusOK,
			args:    args{campID: -123, jrnID: 111},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid jrnID",
			status:  http.StatusOK,
			args:    args{campID: 5272, jrnID: -123},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid args",
			status:  http.StatusOK,
			args:    args{campID: -123, jrnID: -123},
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			args:    args{campID: 5272, jrnID: 111},
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			args:    args{campID: 5272, jrnID: 111},
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			args:    args{campID: 5272, jrnID: 111},
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

			err = c.Journals.Delete(test.args.campID, test.args.jrnID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
		})
	}
}
