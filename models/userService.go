package models

type UserService struct {
	IdModel
	UserId           uint
	ServiceStationId uint
	ServiceId        uint
	StatusRelation
	TimeStamp
}

// TableName return name of database table
func (s *UserService) TableName() string {
	return "user_service"
}
