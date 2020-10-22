package watch

import (
	"reflect"
	"testing"
)

func TestInit(t *testing.T) {
	Init()
}

func TestReadFile(t *testing.T) {
	filename := "example.ini"
	receive, _ := Read_file(filename)

	Print_config(receive)

	expect := map[string]string{
		"app_mode": "development",
		"data":     "/home/git/grafana",
	}

	if !reflect.DeepEqual(receive, expect) {
		t.Errorf("expect:")
	}
}
