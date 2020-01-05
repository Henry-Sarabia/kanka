package kanka

import (
	"net/http"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const testProfileGet string = "test_data/profile_get.json"

func TestProfileService_Get(t *testing.T) {
	prof := &Profile{
		ID:                111,
		Name:              "Henry",
		Avatar:            "https://some-aws-server.com/some_avatar.png",
		AvatarThumb:       "https://some-aws-server.com/some_thumbnail.png",
		Locale:            "en-US",
		Timezone:          "UTC",
		DateFormat:        "m/d/Y",
		DefaultPagination: 45,
		LastCampaignID:    5272,
		IsPatreon:         false,
	}
	tests := []struct {
		name    string
		status  int
		file    string
		want    *Profile
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response",
			status:  http.StatusOK,
			file:    testProfileGet,
			want:    prof,
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

			got, err := c.Profiles.Get()
			if (err != nil) != test.wantErr {
				t.Fatalf("got: <%v>, want error: <%v>", err, test.wantErr)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
