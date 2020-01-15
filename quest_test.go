package kanka

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

const (
	testQuestIndex  string = "test_data/quest_index.json"
	testQuestGet    string = "test_data/quest_get.json"
	testQuestCreate string = "test_data/quest_create.json"
	testQuestUpdate string = "test_data/quest_update.json"
)

func TestQuestService_Index(t *testing.T) {
	qsts := []*Quest{
		&Quest{
			SimpleQuest: SimpleQuest{
				Name: "The Search for Excalibur",
				Type: "Cliche",
			},
		},
		&Quest{
			SimpleQuest: SimpleQuest{
				Name: "Slay the Dragon",
				Type: "Cliche",
			},
		},
		&Quest{
			SimpleQuest: SimpleQuest{
				Name: "Slay the Princess",
				Type: "Original",
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
		want    []*Quest
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testQuestIndex,
			args:    args{campID: 5272, sync: now},
			want:    qsts,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testQuestIndex,
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

			got, err := c.Quests.Index(test.args.campID, test.args.sync)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestQuestService_Get(t *testing.T) {
	qst := &Quest{
		SimpleQuest: SimpleQuest{
			Name:        "Save the Princess",
			Entry:       "\n<p>Just do it</p>\n",
			Image:       "quests/MWmHjSk0NvwkuZMXJyfIc1CIMBqBauUxpFjZIUK7.jpeg",
			IsPrivate:   false,
			Tags:        []int{35131},
			CharacterID: 116623,
			QuestID:     10393,
			Type:        "Cliche",
		},
		ID:             10394,
		ImageFull:      "https://kanka-user-assets.s3.eu-central-1.amazonaws.com/quests/MWmHjSk0NvwkuZMXJyfIc1CIMBqBauUxpFjZIUK7.jpeg",
		ImageThumb:     "https://kanka-user-assets.s3.eu-central-1.amazonaws.com/quests/MWmHjSk0NvwkuZMXJyfIc1CIMBqBauUxpFjZIUK7_thumb.jpeg",
		HasCustomImage: true,
		EntityID:       442249,
		CreatedBy:      5600,
		UpdatedBy:      5600,
		Characters:     0,
		Locations:      0,
	}

	type args struct {
		campID int
		qstID  int
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *Quest
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testQuestGet,
			args:    args{campID: 5272, qstID: 10394},
			want:    qst,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testQuestGet,
			args:    args{campID: -123, qstID: 10394},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid qstID",
			status:  http.StatusOK,
			file:    testQuestGet,
			args:    args{campID: 5272, qstID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testQuestGet,
			args:    args{campID: -123, qstID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, qstID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394},
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

			got, err := c.Quests.Get(test.args.campID, test.args.qstID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestQuestService_Create(t *testing.T) {
	qst := SimpleQuest{
		Name: "Uncover the Mummy",
		Type: "Cliche",
	}
	type args struct {
		campID int
		qst    SimpleQuest
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *Quest
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testQuestCreate,
			args:    args{campID: 5272, qst: qst},
			want:    &Quest{SimpleQuest: qst},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testQuestCreate,
			args:    args{campID: -123, qst: qst},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid quest",
			status:  http.StatusOK,
			file:    testQuestCreate,
			args:    args{campID: 5272, qst: SimpleQuest{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testQuestCreate,
			args:    args{campID: -123, qst: SimpleQuest{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, qst: qst},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, qst: SimpleQuest{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, qst: qst},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, qst: qst},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, qst: qst},
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

			got, err := c.Quests.Create(test.args.campID, test.args.qst)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestQuestService_Update(t *testing.T) {
	qst := SimpleQuest{
		Name: "Children of Empires",
		Type: "Original",
	}
	type args struct {
		campID int
		qstID  int
		qst    SimpleQuest
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *Quest
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testQuestUpdate,
			args:    args{campID: 5272, qstID: 111, qst: qst},
			want:    &Quest{SimpleQuest: qst, ID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testQuestUpdate,
			args:    args{campID: -123, qstID: 111, qst: qst},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid qstID",
			status:  http.StatusOK,
			file:    testQuestUpdate,
			args:    args{campID: 5272, qstID: -123, qst: qst},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid qst",
			status:  http.StatusOK,
			file:    testQuestUpdate,
			args:    args{campID: 5272, qstID: 111, qst: SimpleQuest{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testQuestUpdate,
			args:    args{campID: -123, qstID: -123, qst: SimpleQuest{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 111, qst: qst},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, qstID: -123, qst: SimpleQuest{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 111, qst: qst},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 111, qst: qst},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 111, qst: qst},
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

			got, err := c.Quests.Update(test.args.campID, test.args.qstID, test.args.qst)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestQuestService_Delete(t *testing.T) {
	type args struct {
		campID int
		qstID  int
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
			args:    args{campID: 5272, qstID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, invalid campID",
			status:  http.StatusOK,
			args:    args{campID: -123, qstID: 111},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid qstID",
			status:  http.StatusOK,
			args:    args{campID: 5272, qstID: -123},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid args",
			status:  http.StatusOK,
			args:    args{campID: -123, qstID: -123},
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			args:    args{campID: 5272, qstID: 111},
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			args:    args{campID: 5272, qstID: 111},
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			args:    args{campID: 5272, qstID: 111},
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

			err = c.Quests.Delete(test.args.campID, test.args.qstID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
		})
	}
}
