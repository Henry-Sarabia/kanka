package kanka

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

const (
	testQuestOrganizationIndex  string = "test_data/questorganization_index.json"
	testQuestOrganizationGet    string = "test_data/questorganization_get.json"
	testQuestOrganizationUpdate string = "test_data/questorganization_update.json"
)

func TestQuestOrganizationService_Index(t *testing.T) {
	orgs := []*QuestOrganization{
		&QuestOrganization{
			SimpleQuestOrganization: SimpleQuestOrganization{
				QuestID:        111,
				OrganizationID: 222,
				Role:           "King's Court",
			},
		},
		&QuestOrganization{
			SimpleQuestOrganization: SimpleQuestOrganization{
				QuestID:        333,
				OrganizationID: 444,
				Role:           "Brave Companions",
			},
		},
		&QuestOrganization{
			SimpleQuestOrganization: SimpleQuestOrganization{
				QuestID:        555,
				OrganizationID: 666,
				Role:           "Adventurer's Guild",
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
		want    []*QuestOrganization
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testQuestOrganizationIndex,
			args:    args{campID: 5272, qstID: 10394, sync: now},
			want:    orgs,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testQuestOrganizationIndex,
			args:    args{campID: -123, qstID: 10394, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid qstID",
			status:  http.StatusOK,
			file:    testQuestOrganizationIndex,
			args:    args{campID: 5272, qstID: -123, sync: now},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testQuestOrganizationIndex,
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

			got, err := c.QuestOrganizations.Index(test.args.campID, test.args.qstID, test.args.sync)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestQuestOrganizationService_Get(t *testing.T) {
	org := &QuestOrganization{
		SimpleQuestOrganization: SimpleQuestOrganization{
			OrganizationID: 123,
			Role:           "Villain's Army",
			Description:    "<!DOCTYPE html>\r\n<html>\r\n<head>\r\n</head>\r\n<body>\r\n<p>Villainous scofflaws!</p>\r\n</body>\r\n</html>",
			IsPrivate:      false,
		},
		ID:        715,
		CreatedBy: 0,
		UpdatedBy: 0,
	}

	type args struct {
		campID int
		qstID  int
		orgID  int
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *QuestOrganization
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testQuestOrganizationGet,
			args:    args{campID: 5272, qstID: 10394, orgID: 715},
			want:    org,
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testQuestOrganizationGet,
			args:    args{campID: -123, qstID: 10394, orgID: 715},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid qstID",
			status:  http.StatusOK,
			file:    testQuestOrganizationGet,
			args:    args{campID: 5272, qstID: -123, orgID: 715},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid orgID",
			status:  http.StatusOK,
			file:    testQuestOrganizationGet,
			args:    args{campID: 5272, qstID: 10394, orgID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testQuestOrganizationGet,
			args:    args{campID: -123, qstID: -123, orgID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, orgID: 715},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, qstID: 10394, orgID: -123},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, orgID: 715},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, orgID: 715},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, orgID: 715},
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

			got, err := c.QuestOrganizations.Get(test.args.campID, test.args.qstID, test.args.orgID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestQuestOrganizationService_Update(t *testing.T) {
	org := SimpleQuestOrganization{
		QuestID:        777,
		OrganizationID: 888,
		Role:           "Mage's Guild",
	}
	type args struct {
		campID int
		qstID  int
		orgID  int
		org    SimpleQuestOrganization
	}
	tests := []struct {
		name    string
		status  int
		file    string
		args    args
		want    *QuestOrganization
		wantErr bool
	}{
		{
			name:    "StatusOK, valid response, valid args",
			status:  http.StatusOK,
			file:    testQuestOrganizationUpdate,
			args:    args{campID: 5272, qstID: 10394, orgID: 111, org: org},
			want:    &QuestOrganization{SimpleQuestOrganization: org, ID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid campID",
			status:  http.StatusOK,
			file:    testQuestOrganizationUpdate,
			args:    args{campID: -123, qstID: 10394, orgID: 111, org: org},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid qstID",
			status:  http.StatusOK,
			file:    testQuestOrganizationUpdate,
			args:    args{campID: 5272, qstID: -123, orgID: 111, org: org},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, invalid orgID",
			status:  http.StatusOK,
			file:    testQuestOrganizationUpdate,
			args:    args{campID: 5272, qstID: 10394, orgID: -123, org: org},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, valid response, empty org",
			status:  http.StatusOK,
			file:    testQuestOrganizationUpdate,
			args:    args{campID: 5272, qstID: 10394, orgID: 111, org: SimpleQuestOrganization{}},
			want:    &QuestOrganization{SimpleQuestOrganization: org, ID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, valid response, invalid args",
			status:  http.StatusOK,
			file:    testQuestOrganizationUpdate,
			args:    args{campID: -123, qstID: -123, orgID: -123, org: SimpleQuestOrganization{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, valid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, orgID: 111, org: org},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Status OK, empty response, invalid args",
			status:  http.StatusOK,
			file:    testFileEmpty,
			args:    args{campID: -123, qstID: -123, orgID: -123, org: SimpleQuestOrganization{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, orgID: 111, org: org},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, orgID: 111, org: org},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			file:    testFileEmpty,
			args:    args{campID: 5272, qstID: 10394, orgID: 111, org: org},
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

			got, err := c.QuestOrganizations.Update(test.args.campID, test.args.qstID, test.args.orgID, test.args.org)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestQuestOrganizationService_Delete(t *testing.T) {
	type args struct {
		campID int
		qstID  int
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
			args:    args{campID: 5272, qstID: 10394, orgID: 111},
			wantErr: false,
		},
		{
			name:    "Status OK, invalid campID",
			status:  http.StatusOK,
			args:    args{campID: -123, qstID: 10394, orgID: 111},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid qstID",
			status:  http.StatusOK,
			args:    args{campID: 5272, qstID: -123, orgID: 111},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid orgID",
			status:  http.StatusOK,
			args:    args{campID: 5272, qstID: 10394, orgID: -123},
			wantErr: true,
		},
		{
			name:    "Status OK, invalid args",
			status:  http.StatusOK,
			args:    args{campID: -123, qstID: -123, orgID: -123},
			wantErr: true,
		},
		{
			name:    "StatusUnauthorized, valid args",
			status:  http.StatusUnauthorized,
			args:    args{campID: 5272, qstID: 10394, orgID: 111},
			wantErr: true,
		},
		{
			name:    "StatusForbidden, valid args",
			status:  http.StatusForbidden,
			args:    args{campID: 5272, qstID: 10394, orgID: 111},
			wantErr: true,
		},
		{
			name:    "StatusNotFound, valid args",
			status:  http.StatusNotFound,
			args:    args{campID: 5272, qstID: 10394, orgID: 111},
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

			err = c.QuestOrganizations.Delete(test.args.campID, test.args.qstID, test.args.orgID)
			if (err != nil) != test.wantErr {
				t.Fatalf("got err?: <%t>, want err?: <%t>\nerror: <%v>", (err != nil), test.wantErr, err)
			}
		})
	}
}
