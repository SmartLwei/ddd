package domain

type User struct {
	Model
	Name string `json:"name" gorm:"type:varchar(255);index: idx_name"`
}

type UserFilter struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
}
