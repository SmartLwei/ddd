package db

import "gorm.io/gorm"

type Factory struct {
	demoRepo DemoItf
}

func NewFactory(orm *gorm.DB) *Factory {
	var f Factory
	f.demoRepo = NewDemo(orm)
	return &f
}

func (f *Factory) GetDemoRepo() DemoItf {
	return f.demoRepo
}
