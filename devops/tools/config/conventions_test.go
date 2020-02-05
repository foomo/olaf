package config

import (
	"path"
	"runtime"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func GetFileFromCurrentDir(file string) string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(filename), file)
}

func TestCollect(t *testing.T) {
	groups, errCollect := Collect(GetFileFromCurrentDir("../../../"), nil)
	spew.Dump(groups, errCollect)
}

func Test_collectFrontends(t *testing.T) {
	frontends, errCollect := collectFrontends(GetFileFromCurrentDir("../../.."), nil)
	spew.Dump(frontends, errCollect)
}
func Test_collectServices(t *testing.T) {
	services, errCollect := collectServices(GetFileFromCurrentDir("../../.."), nil)
	spew.Dump(services, errCollect)
}

func _Test_extractNameAndGroup(t *testing.T) {

}

func Test_extractNameAndGroup(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name      string
		args      args
		wantName  string
		wantGroup GroupName
		wantErr   bool
		err       error
	}{
		{
			name:      "valid name and group",
			args:      args{n: "helloworld_admin"},
			wantName:  "admin",
			wantGroup: GroupName("helloworld"),
			wantErr:   false,
		},
		{
			name:      "name only",
			args:      args{n: "helloworld"},
			wantName:  "helloworld",
			wantGroup: GroupName("helloworld"),
			wantErr:   false,
		},
		{
			name:      "too many parts",
			args:      args{n: "helloworld_foo_bar"},
			wantName:  "",
			wantGroup: GroupName(""),
			wantErr:   true,
			err:       errExtractNameAndGroupTooManyParts,
		},
		{
			name:      "empty group err",
			args:      args{n: "_foo"},
			wantName:  "",
			wantGroup: GroupName(""),
			wantErr:   true,
			err:       errExtractNameAndGroupGroupIsEmpty,
		},
		{
			name:      "empty name err",
			args:      args{n: "foo_"},
			wantName:  "",
			wantGroup: GroupName(""),
			wantErr:   true,
			err:       errExtractNameAndGroupNameIsEmpty,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotName, gotGroup, err := extractNameAndGroup(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractNameAndGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotName != tt.wantName {
				t.Errorf("extractNameAndGroup() gotName = %v, want %v", gotName, tt.wantName)
			}
			if gotGroup != tt.wantGroup {
				t.Errorf("extractNameAndGroup() gotGroup = %v, want %v", gotGroup, tt.wantGroup)
			}
			if tt.err != nil {
				if tt.err != err {
					t.Errorf("unexpected error = %v expected = %v", err, tt.err)
				}
			}
		})
	}
}
