package utils

import (
	"reflect"
	"testing"

	"github.com/henrylee2cn/aster/aster"
)

func TestLoadDirs(t *testing.T) {
	type args struct {
		dirs []string
	}
	tests := []struct {
		name    string
		args    args
		want    *aster.Program
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadDirs(tt.args.dirs...)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadDirs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadDirs() = %v, want %v", got, tt.want)
			}
		})
	}
}
