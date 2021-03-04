package domain

type Demo struct {
	Model
	Name string `json:"name" gorm:"type:varchar(255);index: idx_name"`
}

type DemoFilter struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
}
