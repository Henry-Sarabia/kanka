package kanka

import (
	"net/http"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const (
	testCampaignGet string = "test_data/campaign_get.json"
)

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
						Avatar: "users/aFstunGZFZ53yr3Zy19IqKBoceuwFGaAJVFfOKUf.png",
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
