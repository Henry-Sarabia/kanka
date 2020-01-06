package kanka

import (
	"net/http"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const (
	testCampaignIndex   string = "test_data/campaign_index.json"
	testCampaignGet     string = "test_data/campaign_get.json"
	testCampaignMembers string = "test_data/campaign_members.json"
)

func TestCampaignService_Index(t *testing.T) {
	camps := []*Campaign{
		&Campaign{
			Name: "The Adventures of Adventurers",
			ID:   111,
		},
		&Campaign{
			Name: "The Heroics of Heroes and Heroines",
			ID:   222,
		},
		&Campaign{
			Name: "The Traversings of Travelers",
			ID:   333,
		},
	}
	tests := []struct {
		name    string
		status  int
		file    string
		want    []*Campaign
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response",
			status:  http.StatusOK,
			file:    testCampaignIndex,
			want:    camps,
			wantErr: false,
		},
		{
			name:    "Status OK, empty response",
			status:  http.StatusOK,
			file:    testFileEmpty,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
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

			got, err := c.Campaigns.Index()
			if (err != nil) != test.wantErr {
				t.Fatalf("got: <%v>, want error: <%v>", err, test.wantErr)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestCampaignService_Get(t *testing.T) {
	camp := &Campaign{
		ID:         5272,
		Name:       "Children of Empires",
		Locale:     "en",
		Entry:      "\n\n",
		Image:      "",
		ImageFull:  "https://some-aws-server.com/cool_swords_and_stuff.png",
		ImageThumb: "https://some-aws-server.com/cool_thumbnail.png",
		Visibility: "private",
		Members: Members{
			Data: []*Member{
				&Member{
					ID: 111,
					User: User{
						ID:     222,
						Name:   "Henry",
						Avatar: "some_avatar.png",
					},
				},
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
		want    *Campaign
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testCampaignGet,
			args:    args{campID: 5272},
			want:    camp,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testCampaignGet,
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

			got, err := c.Campaigns.Get(test.args.campID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got: <%v>, want error: <%v>", err, test.wantErr)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestCampaignService_Members(t *testing.T) {
	mems := []*Member{
		&Member{
			ID: 111,
			User: User{
				ID:     222,
				Name:   "Jon",
				Avatar: "jon_brooding.png",
			},
		},
		&Member{
			ID: 333,
			User: User{
				ID:     444,
				Name:   "Daenerys",
				Avatar: "daeny_burning_something.png",
			},
		},
		&Member{
			ID: 555,
			User: User{
				ID:     666,
				Name:   "Stannis",
				Avatar: "stannis_also_brooding.png",
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
		want    []*Member
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testCampaignMembers,
			args:    args{campID: 5272},
			want:    mems,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testCampaignMembers,
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

			got, err := c.Campaigns.Members(test.args.campID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got: <%v>, want error: <%v>", err, test.wantErr)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
