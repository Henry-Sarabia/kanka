package kanka

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const (
	testFileEmpty      string = "test_data/empty.json"
	testCharacterGet   string = "test_data/character_get.json"
	testCharacterIndex string = "test_data/character_index.json"
)

func testClient(status int, resp io.Reader) (*Client, *httptest.Server) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		io.Copy(w, resp)
	}))

	c := NewClient(testToken, ts.Client())
	c.rootURL = ts.URL + "/"

	return c, ts
}

func TestCharacterService_Index(t *testing.T) {
	chars := []*Character{
		&Character{
			SimpleCharacter: SimpleCharacter{
				Name:  "Jon Snow",
				Title: "Bastard of Winterfell",
			},
		},
		&Character{
			SimpleCharacter: SimpleCharacter{
				Name:  "Sansa Stark",
				Title: "Lady of Winterfell",
			},
		},
		&Character{
			SimpleCharacter: SimpleCharacter{
				Name:  "Daenerys Targaryen",
				Title: "Mother of Dragons",
			},
		},
	}
	type args struct {
		campID int
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    []*Character
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testCharacterIndex,
			args:    args{campID: 5272},
			want:    chars,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testCharacterIndex,
			args:    args{campID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272},
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

			got, err := c.Characters.Index(test.args.campID)
			if (err != nil) != test.wantErr {
				t.Errorf("got: <%v>, want error: <%v>", err, test.wantErr)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestCharacterService_Get(t *testing.T) {
	char := &Character{
		SimpleCharacter: SimpleCharacter{
			Name:       "Penny Galvenrise",
			Image:      "characters/pdt4F7zJjCyxDUu2flaZXBPqwHtkhCg8fmowXV05.jpeg",
			IsPrivate:  false,
			Tags:       []int{34696},
			LocationID: 26145,
			Title:      "Tinkerer",
			Age:        "24",
			Sex:        "Female",
			RaceID:     9477,
			Type:       "NPC",
			FamilyID:   16158,
			IsDead:     false,
		},
		ID:             116623,
		Entry:          "\n<p>She is the key to finding Mechanus</p>\n",
		ImageFull:      "https://kanka-user-assets.s3.eu-central-1.amazonaws.com/characters/pdt4F7zJjCyxDUu2flaZXBPqwHtkhCg8fmowXV05.jpeg",
		ImageThumb:     "https://kanka-user-assets.s3.eu-central-1.amazonaws.com/characters/pdt4F7zJjCyxDUu2flaZXBPqwHtkhCg8fmowXV05_thumb.jpeg",
		HasCustomImage: true,
		EntityID:       430214,
		CreatedAt:      "2020-01-03T01:18:30.000000Z",
		CreatedBy:      5600,
		UpdatedAt:      "2020-01-03T01:18:30.000000Z",
		UpdatedBy:      5600,
		Traits: Traits{
			Data: []*Trait{
				&Trait{
					ID:           85283,
					Name:         "Fears",
					Entry:        "Meteors",
					Section:      "personality",
					IsPrivate:    false,
					DefaultOrder: 0,
				},
				&Trait{
					ID:           85284,
					Name:         "Goals",
					Entry:        "Create a world-changing invention",
					Section:      "personality",
					IsPrivate:    false,
					DefaultOrder: 1,
				},
				&Trait{
					ID:           85285,
					Name:         "Mannerisms",
					Entry:        "Goes on a lot of tangents; Speaks rapidly",
					Section:      "personality",
					IsPrivate:    false,
					DefaultOrder: 2,
				},
				&Trait{
					ID:           85286,
					Name:         "Traits",
					Entry:        "Talkative\r\nBubbly",
					Section:      "personality",
					IsPrivate:    false,
					DefaultOrder: 3,
				},
				&Trait{
					ID:           85287,
					Name:         "Hair",
					Entry:        "Pinks",
					Section:      "appearance",
					IsPrivate:    false,
					DefaultOrder: 0,
				},
				&Trait{
					ID:           85288,
					Name:         "Eyes",
					Entry:        "Green",
					Section:      "appearance",
					IsPrivate:    false,
					DefaultOrder: 1,
				},
				&Trait{
					ID:           85289,
					Name:         "Skin",
					Entry:        "Fair",
					Section:      "appearance",
					IsPrivate:    false,
					DefaultOrder: 2,
				},
			},
		},
	}

	type args struct {
		campID int
		charID int
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *Character
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testCharacterGet,
			args:    args{campID: 5272, charID: 116623},
			want:    char,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testCharacterGet,
			args:    args{campID: -123, charID: 116623},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid charID",
			status:  http.StatusOK,
			file:    testCharacterGet,
			args:    args{campID: 5272, charID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testCharacterGet,
			args:    args{campID: -123, charID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, charID: 116623},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, charID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, charID: 116623},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, charID: 116623},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, charID: 116623},
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

			got, err := c.Characters.Get(test.args.campID, test.args.charID)
			if (err != nil) != test.wantErr {
				t.Errorf("got: <%v>, want error: <%v>", err, test.wantErr)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
