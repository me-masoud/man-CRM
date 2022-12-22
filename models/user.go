package models

// User structure
type User struct {
	IdModel
	Name  string `gorm:"type:varchar(50)" json:"name" validate:"required"`
	Email string `gorm:"type:varchar(50)" json:"email" validate:"required,email"`
	TimeStamp
}

// TableName return name of database table
func (u *User) TableName() string {
	return "users"
}
