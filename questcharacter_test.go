package kanka

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

const (
	testQuestCharacterIndex  string = "test_data/questcharacter_index.json"
	testQuestCharacterGet    string = "test_data/questcharacter_get.json"
	testQuestCharacterCreate string = "test_data/questcharacter_create.json"
	testQuestCharacterUpdate string = "test_data/questcharacter_update.json"
)

func TestQuestCharacterService_Index(t *testing.T) {
	qchs := []*QuestCharacter{
		{
			SimpleQuestCharacter: SimpleQuestCharacter{
				QuestID:     111,
				CharacterID: 222,
				Role:        "Hero",
			},
		},
		{
			SimpleQuestCharacter: SimpleQuestCharacter{
				QuestID:     333,
				CharacterID: 444,
				Role:        "Goddess",
			},
		},
		{
			SimpleQuestCharacter: SimpleQuestCharacter{
				QuestID:     555,
				CharacterID: 666,
				Role:        "Father",
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
		want    []*QuestCharacter
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testQuestCharacterIndex,
			args:    args{campID: 5272, qstID: 10394, sync: now},
			want:    qchs,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testQuestCharacterIndex,
			args:    args{campID: -123, qstID: 10394, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid qstID",
			status:  http.StatusOK,
			file:    testQuestCharacterIndex,
			args:    args{campID: 5272, qstID: -123, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testQuestCharacterIndex,
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

			got, err := c.QuestCharacters.Index(test.args.campID, test.args.qstID, test.args.sync)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestQuestCharacterService_Get(t *testing.T) {
	qch := &QuestCharacter{
		SimpleQuestCharacter: SimpleQuestCharacter{
			CharacterID: 24326,
			Description: "\n<p>The princess trapped in the tower</p>\n",
			Role:        "Princess",
		},
		ID:        6849,
		CreatedBy: 0,
		UpdatedBy: 0,
	}

	type args struct {
		campID int
		qstID  int
		qchID  int
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *QuestCharacter
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testQuestCharacterGet,
			args:    args{campID: 5272, qstID: 10394, qchID: 6849},
			want:    qch,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testQuestCharacterGet,
			args:    args{campID: -123, qstID: 10394, qchID: 6849},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid qstID",
			status:  http.StatusOK,
			file:    testQuestCharacterGet,
			args:    args{campID: 5272, qstID: -123, qchID: 6849},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid qchID",
			status:  http.StatusOK,
			file:    testQuestCharacterGet,
			args:    args{campID: 5272, qstID: 10394, qchID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testQuestCharacterGet,
			args:    args{campID: -123, qstID: -123, qchID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, qchID: 6849},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, qstID: 10394, qchID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, qchID: 6849},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, qchID: 6849},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, qchID: 6849},
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

			got, err := c.QuestCharacters.Get(test.args.campID, test.args.qstID, test.args.qchID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestQuestCharacterService_Create(t *testing.T) {
	qch := SimpleQuestCharacter{
		QuestID:     777,
		CharacterID: 888,
		Role:        "Threshold Guardian",
	}
	type args struct {
		campID int
		qstID  int
		qch    SimpleQuestCharacter
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *QuestCharacter
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testQuestCharacterCreate,
			args:    args{campID: 5272, qstID: 10394, qch: qch},
			want:    &QuestCharacter{SimpleQuestCharacter: qch},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testQuestCharacterCreate,
			args:    args{campID: -123, qstID: 10394, qch: qch},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid qstID",
			status:  http.StatusOK,
			file:    testQuestCharacterCreate,
			args:    args{campID: 5272, qstID: -123, qch: qch},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, empty qch",
			status:  http.StatusOK,
			file:    testQuestCharacterCreate,
			args:    args{campID: 5272, qstID: 10394, qch: SimpleQuestCharacter{}},
			want:    &QuestCharacter{SimpleQuestCharacter: qch},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testQuestCharacterCreate,
			args:    args{campID: -123, qstID: -123, qch: SimpleQuestCharacter{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, qch: qch},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, qstID: -123, qch: SimpleQuestCharacter{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, qch: qch},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, qch: qch},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, qch: qch},
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

			got, err := c.QuestCharacters.Create(test.args.campID, test.args.qstID, test.args.qch)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestQuestCharacterService_Update(t *testing.T) {
	qch := SimpleQuestCharacter{
		QuestID:     999,
		CharacterID: 101010,
		Role:        "Temptress",
	}
	type args struct {
		campID int
		qstID  int
		qchID  int
		qch    SimpleQuestCharacter
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *QuestCharacter
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testQuestCharacterUpdate,
			args:    args{campID: 5272, qstID: 10394, qchID: 111, qch: qch},
			want:    &QuestCharacter{SimpleQuestCharacter: qch, ID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testQuestCharacterUpdate,
			args:    args{campID: -123, qstID: 10394, qchID: 111, qch: qch},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid qstID",
			status:  http.StatusOK,
			file:    testQuestCharacterUpdate,
			args:    args{campID: 5272, qstID: -123, qchID: 111, qch: qch},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid qchID",
			status:  http.StatusOK,
			file:    testQuestCharacterUpdate,
			args:    args{campID: 5272, qstID: 10394, qchID: -123, qch: qch},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, empty qch",
			status:  http.StatusOK,
			file:    testQuestCharacterUpdate,
			args:    args{campID: 5272, qstID: 10394, qchID: 111, qch: SimpleQuestCharacter{}},
			want:    &QuestCharacter{SimpleQuestCharacter: qch, ID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testQuestCharacterUpdate,
			args:    args{campID: -123, qstID: -123, qchID: -123, qch: SimpleQuestCharacter{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, qchID: 111, qch: qch},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, qstID: -123, qchID: -123, qch: SimpleQuestCharacter{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, qchID: 111, qch: qch},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, qchID: 111, qch: qch},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, qchID: 111, qch: qch},
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

			got, err := c.QuestCharacters.Update(test.args.campID, test.args.qstID, test.args.qchID, test.args.qch)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestQuestCharacterService_Delete(t *testing.T) {
	type args struct {
		campID int
		qstID  int
		qchID  int
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
			args:    args{campID: 5272, qstID: 10394, qchID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, invalid campID",
			status:  http.StatusOK,
			args:    args{campID: -123, qstID: 10394, qchID: 111},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid qstID",
			status:  http.StatusOK,
			args:    args{campID: 5272, qstID: -123, qchID: 111},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid qchID",
			status:  http.StatusOK,
			args:    args{campID: 5272, qstID: 10394, qchID: -123},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid args",
			status:  http.StatusOK,
			args:    args{campID: -123, qstID: -123, qchID: -123},
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			args:    args{campID: 5272, qstID: 10394, qchID: 111},
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			args:    args{campID: 5272, qstID: 10394, qchID: 111},
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			args:    args{campID: 5272, qstID: 10394, qchID: 111},
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

			err = c.QuestCharacters.Delete(test.args.campID, test.args.qstID, test.args.qchID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
		})
	}
}
