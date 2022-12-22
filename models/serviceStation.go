package models

type ServiceStation struct {
	IdModel
	Name        string `gorm:"type:varchar(50)" json:"name" validate:"required"`
	Description string `gorm:"type:varchar(50)" json:"description" validate:"required"`
	StatusRelation
	TimeStamp
}

// TableName return name of database table
func (s *ServiceStation) TableName() string {
	return "service_station"
}
