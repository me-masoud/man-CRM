package models

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

var db *gorm.DB

// IdModel is sample of common table structure
type IdModel struct {
	ID uint `gorm:"primary_key" json:"id,omitempty"`
}

type StatusRelation struct {
	StatusId uint `gorm:"not null" json:"status_id" sql:"DEFAULT:1"`
}

type TimeStamp struct {
	CreatedAt time.Time  `gorm:"not null" json:"created_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
}

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	msql := mysql.Config{}
	log.Println(msql)

	conn, err := gorm.Open("mysql", username+":"+password+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8&parseTime=True&loc=Asia%2FKolkata")

	if err != nil {
		fmt.Print(err)
	}
	db = conn

	//Printing query
	db.LogMode(true)

	//Automatically create migration as per model
	db.Debug().AutoMigrate(
		&User{},
		&Status{},
		&Service{},
		&ServiceStation{},
		&Stuff{},
		&UserStuff{},
		&UserService{},
	)
}

// GetDB function return the instance of db
func GetDB() *gorm.DB {
	return db
}
