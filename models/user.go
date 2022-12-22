package models

// User structure
type User struct {
	IdModel
	Name     string `gorm:"type:varchar(50)" json:"name" validate:"required"`
	Email    string `gorm:"type:varchar(50)" json:"email" validate:"required,email"`
	Phone    string `gorm:"type:varchar(50)" json:"phone" validate:"required"`
	Password string `gorm:"type:varchar(100)" json:"password" validate:"required"`
	StatusRelation
	TimeStamp
}

// TableName return name of database table
func (u *User) TableName() string {
	return "users"
}
