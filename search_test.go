package kanka

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

const testSearch string = "test_data/search.json"

func TestClient_Search(t *testing.T) {
	rslt := []*Result{
		&Result{
			ID:                  1194,
			EntityID:            80025,
			Name:                "Shop",
			Image:               "https://kanka.io/images/defaults/attribute_templates.jpg",
			ImageThumb:          "https://kanka.io/images/defaults/attribute_templates_thumb.jpg",
			HasCustomImage:      false,
			Type:                "attribute_template",
			Tooltip:             "Shop",
			URL:                 "https://kanka.io/campaign/5272/attribute_templates/1194",
			IsAttributesPrivate: 0,
			IsPrivate:           false,
			CreatedBy:           5600,
			UpdatedBy:           5600,
		},
		&Result{
			ID:                  26141,
			EntityID:            80918,
			Name:                "The Rope Shop",
			Image:               "https://kanka.io/images/defaults/locations.jpg",
			ImageThumb:          "https://kanka.io/images/defaults/locations_thumb.jpg",
			HasCustomImage:      false,
			Type:                "location",
			Tooltip:             "",
			URL:                 "https://kanka.io/campaign/5272/locations/26141",
			IsAttributesPrivate: 0,
			IsPrivate:           false,
			CreatedBy:           5600,
			UpdatedBy:           5600,
		},
		&Result{
			ID:                  912,
			EntityID:            443499,
			Name:                "At The Magic Shop",
			Image:               "https://kanka.io/images/defaults/conversations.jpg",
			ImageThumb:          "https://kanka.io/images/defaults/conversations_thumb.jpg",
			HasCustomImage:      false,
			Type:                "conversation",
			Tooltip:             "At The Magic Shop",
			URL:                 "https://kanka.io/campaign/5272/conversations/912",
			IsAttributesPrivate: 0,
			IsPrivate:           false,
			CreatedBy:           5600,
			UpdatedBy:           5600,
		},
	}
	n := time.Now()
	now := &n

	type args struct {
		campID int
		qry    string
		sync   *time.Time
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    []*Result
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testSearch,
			args:    args{campID: 5272, qry: "shop", sync: now},
			want:    rslt,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testSearch,
			args:    args{campID: -123, qry: "shop", sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid qry",
			status:  http.StatusOK,
			file:    testSearch,
			args:    args{campID: 5272, qry: "", sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testSearch,
			args:    args{campID: -123, qry: "", sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, qry: "shop", sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, qry: "", sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, qry: "shop", sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, qry: "shop", sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, qry: "shop", sync: now},
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

			got, err := c.Search(test.args.campID, test.args.qry, test.args.sync)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
