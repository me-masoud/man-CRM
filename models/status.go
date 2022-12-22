package models

type Status struct {
	IdModel
	Name        string `gorm:"type:varchar(50)" json:"name" validate:"required"`
	Key         string `gorm:"type:varchar(50)" json:"key" validate:"required"`
	Description string `gorm:"type:varchar(50)" json:"description" validate:"required"`
	Color       string `gorm:"type:varchar(50)" json:"color" validate:"required"`
	TimeStamp
}

// TableName return name of database table
func (s *Status) TableName() string {
	return "statuses"
}
