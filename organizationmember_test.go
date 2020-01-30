package kanka

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

const (
	testOrganizationMemberIndex  string = "test_data/organizationmember_index.json"
	testOrganizationMemberGet    string = "test_data/organizationmember_get.json"
	testOrganizationMemberUpdate string = "test_data/organizationmember_update.json"
)

func TestOrganizationMemberService_Index(t *testing.T) {
	mems := []*OrganizationMember{
		&OrganizationMember{
			SimpleOrganizationMember: SimpleOrganizationMember{
				CharacterID:    118490,
				IsPrivate:      false,
				OrganizationID: 23579,
				Role:           "Leader",
			},
			ID: 42070,
		},
		&OrganizationMember{
			SimpleOrganizationMember: SimpleOrganizationMember{
				CharacterID:    126247,
				IsPrivate:      false,
				OrganizationID: 23579,
				Role:           "Maester",
			},
			ID: 48013,
		},
		&OrganizationMember{
			SimpleOrganizationMember: SimpleOrganizationMember{
				CharacterID:    126248,
				IsPrivate:      false,
				OrganizationID: 23579,
				Role:           "Soldier",
			},
			ID: 48017,
		},
	}
	n := time.Now()
	now := &n

	type args struct {
		campID int
		orgID  int
		sync   *time.Time
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    []*OrganizationMember
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testOrganizationMemberIndex,
			args:    args{campID: 5272, orgID: 23579, sync: now},
			want:    mems,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testOrganizationMemberIndex,
			args:    args{campID: -123, orgID: 23579, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid orgID",
			status:  http.StatusOK,
			file:    testOrganizationMemberIndex,
			args:    args{campID: 5272, orgID: -123, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testOrganizationMemberIndex,
			args:    args{campID: -123, orgID: -123, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, orgID: 23579, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, orgID: -123, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, orgID: 23579, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, orgID: 23579, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, orgID: 23579, sync: now},
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

			got, err := c.OrganizationMembers.Index(test.args.campID, test.args.orgID, test.args.sync)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestOrganizationMemberService_Get(t *testing.T) {
	mem := &OrganizationMember{
		SimpleOrganizationMember: SimpleOrganizationMember{
			CharacterID:    118490,
			IsPrivate:      false,
			OrganizationID: 23579,
			Role:           "Leader",
		},
		ID:        42070,
		CreatedBy: 0,
		UpdatedBy: 0,
	}

	type args struct {
		campID int
		orgID  int
		memID  int
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *OrganizationMember
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testOrganizationMemberGet,
			args:    args{campID: 5272, orgID: 23579, memID: 42070},
			want:    mem,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testOrganizationMemberGet,
			args:    args{campID: -123, orgID: 23579, memID: 42070},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid orgID",
			status:  http.StatusOK,
			file:    testOrganizationMemberGet,
			args:    args{campID: 5272, orgID: -123, memID: 42070},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid memID",
			status:  http.StatusOK,
			file:    testOrganizationMemberGet,
			args:    args{campID: 5272, orgID: 23579, memID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testOrganizationMemberGet,
			args:    args{campID: -123, orgID: -123, memID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, orgID: 23579, memID: 42070},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, orgID: 23579, memID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, orgID: 23579, memID: 42070},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, orgID: 23579, memID: 42070},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, orgID: 23579, memID: 42070},
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

			got, err := c.OrganizationMembers.Get(test.args.campID, test.args.orgID, test.args.memID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestOrganizationMemberService_Update(t *testing.T) {
	mem := SimpleOrganizationMember{
		Role:           "Treasurer",
		CharacterID:    111,
		OrganizationID: 222,
		IsPrivate:      false,
	}
	type args struct {
		campID int
		orgID  int
		memID  int
		mem    SimpleOrganizationMember
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *OrganizationMember
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testOrganizationMemberUpdate,
			args:    args{campID: 5272, orgID: 23579, memID: 111, mem: mem},
			want:    &OrganizationMember{SimpleOrganizationMember: mem, ID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testOrganizationMemberUpdate,
			args:    args{campID: -123, orgID: 23579, memID: 111, mem: mem},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid orgID",
			status:  http.StatusOK,
			file:    testOrganizationMemberUpdate,
			args:    args{campID: 5272, orgID: -123, memID: 111, mem: mem},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid memID",
			status:  http.StatusOK,
			file:    testOrganizationMemberUpdate,
			args:    args{campID: 5272, orgID: 23579, memID: -123, mem: mem},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, empty mem",
			status:  http.StatusOK,
			file:    testOrganizationMemberUpdate,
			args:    args{campID: 5272, orgID: 23579, memID: 111, mem: SimpleOrganizationMember{}},
			want:    &OrganizationMember{SimpleOrganizationMember: mem, ID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testOrganizationMemberUpdate,
			args:    args{campID: -123, orgID: -123, memID: -123, mem: SimpleOrganizationMember{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, orgID: 23579, memID: 111, mem: mem},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, orgID: -123, memID: -123, mem: SimpleOrganizationMember{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, orgID: 23579, memID: 111, mem: mem},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, orgID: 23579, memID: 111, mem: mem},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, orgID: 23579, memID: 111, mem: mem},
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

			got, err := c.OrganizationMembers.Update(test.args.campID, test.args.orgID, test.args.memID, test.args.mem)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestOrganizationMemberService_Delete(t *testing.T) {
	type args struct {
		campID int
		orgID  int
		memID  int
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
			args:    args{campID: 5272, orgID: 23579, memID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, invalid campID",
			status:  http.StatusOK,
			args:    args{campID: -123, orgID: 23579, memID: 111},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid orgID",
			status:  http.StatusOK,
			args:    args{campID: -123, orgID: -123, memID: 111},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid memID",
			status:  http.StatusOK,
			args:    args{campID: 5272, orgID: 23579, memID: -123},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid args",
			status:  http.StatusOK,
			args:    args{campID: -123, orgID: -123, memID: -123},
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			args:    args{campID: 5272, orgID: 23579, memID: 111},
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			args:    args{campID: 5272, orgID: 23579, memID: 111},
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			args:    args{campID: 5272, orgID: 23579, memID: 111},
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

			err = c.OrganizationMembers.Delete(test.args.campID, test.args.orgID, test.args.memID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
		})
	}
}
