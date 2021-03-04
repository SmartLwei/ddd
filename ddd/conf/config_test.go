package conf

import (
	"reflect"
	"testing"
)

func TestInit(t *testing.T) {
	var fileName = "config_example.yaml"
	var wantSetting = &Setting{
		Rest: &RestSetting{Port: ":8080", Mode: "debug"},
		DB: &DBSetting{
			URI:          "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local",
			LogLevel:     4,
			MaxIdleConns: 10,
			MaxOpenConns: 20,
		},
	}
	Init(fileName)
	setting = GetSetting()
	if eq := reflect.DeepEqual(wantSetting, setting); !eq {
		t.Errorf("GetSetting() got = \n%+v\n, want \n%+v\n", setting, wantSetting)
	}
}
