package models

type UserStuff struct {
	IdModel
	UserId  uint
	StuffId uint
	StatusRelation
	TimeStamp
}

// TableName return name of database table
func (s *UserStuff) TableName() string {
	return "user_stuff"
}
