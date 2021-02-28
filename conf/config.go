package conf

import "sync"

type Setting struct {
	Rest *RestSetting
}

var once = &sync.Once{}
var setting *Setting

type RestSetting struct {
	Port string `json:"port" yaml:"port"`
	Mode string `json:"mode" yaml:"mode"`
}

func Init(fileName string) {
	once.Do(func() {
		setting = &Setting{Rest: &RestSetting{Port: ":8080"}}
	})
}

func GetSetting() *Setting {
	if setting == nil {
		panic("get setting failed, please init setting first")
	}
	return setting
}
