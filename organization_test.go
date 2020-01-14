package kanka

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

const (
	testOrganizationIndex  string = "test_data/organization_index.json"
	testOrganizationGet    string = "test_data/organization_get.json"
	testOrganizationCreate string = "test_data/organization_create.json"
	testOrganizationUpdate string = "test_data/organization_update.json"
)

func TestOrganizationService_Index(t *testing.T) {
	orgs := []*Organization{
		&Organization{
			SimpleOrganization: SimpleOrganization{
				Name: "Golden Company",
				Type: "Company",
			},
		},
		&Organization{
			SimpleOrganization: SimpleOrganization{
				Name: "Iron Bank",
				Type: "Bank",
			},
		},
		&Organization{
			SimpleOrganization: SimpleOrganization{
				Name: "Brotherhood of Banners",
				Type: "Company",
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
		want    []*Organization
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testOrganizationIndex,
			args:    args{campID: 5272, sync: now},
			want:    orgs,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testOrganizationIndex,
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

			got, err := c.Organizations.Index(test.args.campID, test.args.sync)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestOrganizationService_Get(t *testing.T) {
	org := &Organization{
		SimpleOrganization: SimpleOrganization{
			Name:           "Brave Companions",
			Entry:          "\n<p>Also known as the \"Bloody Mummers\"</p>\n",
			Type:           "Company",
			Image:          "organisations/PZQVp6lFwpcXSbw0kERQUfklZ7nSc4rAlieHaieh.png",
			IsPrivate:      false,
			Tags:           []int{35131},
			LocationID:     115366,
			OrganizationID: 23578,
		},
		ID:             23579,
		ImageFull:      "https://kanka-user-assets.s3.eu-central-1.amazonaws.com/organisations/PZQVp6lFwpcXSbw0kERQUfklZ7nSc4rAlieHaieh.png",
		ImageThumb:     "https://kanka-user-assets.s3.eu-central-1.amazonaws.com/organisations/PZQVp6lFwpcXSbw0kERQUfklZ7nSc4rAlieHaieh_thumb.png",
		HasCustomImage: true,
		EntityID:       437024,
		CreatedBy:      5600,
		UpdatedBy:      5600,
		Members:        1,
	}

	type args struct {
		campID int
		orgID  int
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *Organization
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testOrganizationGet,
			args:    args{campID: 5272, orgID: 23579},
			want:    org,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testOrganizationGet,
			args:    args{campID: -123, orgID: 23579},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid orgID",
			status:  http.StatusOK,
			file:    testOrganizationGet,
			args:    args{campID: 5272, orgID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testOrganizationGet,
			args:    args{campID: -123, orgID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, orgID: 23579},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, orgID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, orgID: 23579},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, orgID: 23579},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, orgID: 23579},
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

			got, err := c.Organizations.Get(test.args.campID, test.args.orgID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestOrganizationService_Create(t *testing.T) {
	org := SimpleOrganization{
		Name: "Knights of the Vale",
		Type: "Army",
	}
	type args struct {
		campID int
		org    SimpleOrganization
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *Organization
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testOrganizationCreate,
			args:    args{campID: 5272, org: org},
			want:    &Organization{SimpleOrganization: org},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testOrganizationCreate,
			args:    args{campID: -123, org: org},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid organization",
			status:  http.StatusOK,
			file:    testOrganizationCreate,
			args:    args{campID: 5272, org: SimpleOrganization{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testOrganizationCreate,
			args:    args{campID: -123, org: SimpleOrganization{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, org: org},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, org: SimpleOrganization{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, org: org},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, org: org},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, org: org},
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

			got, err := c.Organizations.Create(test.args.campID, test.args.org)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestOrganizationService_Update(t *testing.T) {
	org := SimpleOrganization{
		Name: "Iron Fleet",
		Type: "Armada",
	}
	type args struct {
		campID int
		orgID  int
		org    SimpleOrganization
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *Organization
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testOrganizationUpdate,
			args:    args{campID: 5272, orgID: 111, org: org},
			want:    &Organization{SimpleOrganization: org, ID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testOrganizationUpdate,
			args:    args{campID: -123, orgID: 111, org: org},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid orgID",
			status:  http.StatusOK,
			file:    testOrganizationUpdate,
			args:    args{campID: 5272, orgID: -123, org: org},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid org",
			status:  http.StatusOK,
			file:    testOrganizationUpdate,
			args:    args{campID: 5272, orgID: 111, org: SimpleOrganization{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testOrganizationUpdate,
			args:    args{campID: -123, orgID: -123, org: SimpleOrganization{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, orgID: 111, org: org},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, orgID: -123, org: SimpleOrganization{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, orgID: 111, org: org},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, orgID: 111, org: org},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, orgID: 111, org: org},
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

			got, err := c.Organizations.Update(test.args.campID, test.args.orgID, test.args.org)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestOrganizationService_Delete(t *testing.T) {
	type args struct {
		campID int
		orgID  int
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
			args:    args{campID: 5272, orgID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, invalid campID",
			status:  http.StatusOK,
			args:    args{campID: -123, orgID: 111},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid orgID",
			status:  http.StatusOK,
			args:    args{campID: 5272, orgID: -123},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid args",
			status:  http.StatusOK,
			args:    args{campID: -123, orgID: -123},
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			args:    args{campID: 5272, orgID: 111},
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			args:    args{campID: 5272, orgID: 111},
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			args:    args{campID: 5272, orgID: 111},
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

			err = c.Organizations.Delete(test.args.campID, test.args.orgID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
		})
	}
}
