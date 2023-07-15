package models

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	NotesTable = "xnotes"
	UsersTable = "xusers"
)

type UserS struct {
	ID        uint64         `gorm:"primaryKey"`
	Name      string         `gorm:"size:255"`
	Email     string         `gorm:"size:255"`
	Password  string         `gorm:"size:255"`
	CreatedAt time.Time      `gorm:"index"`
	UpdatedAt time.Time      `gorm:"index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Note struct {
	ID        uint64         `gorm:"primaryKey"`
	Name      string         `gorm:"size:255"`
	Notes     string         `gorm:"size:255"`
	UserID    uint64         `gorm:"index"`
	CreatedAt time.Time      `gorm:"index"`
	UpdatedAt time.Time      `gorm:"index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Database Credentials
const (
	DB_HOST = "127.0.0.1"
	DB_PORT = "3306"
	DB_NAME = "accuknox"
	DB_USER = "root"
	DB_PASS = ""
)

var DB *gorm.DB

func ConnectDB() {
	sql := DB_USER + ":" + DB_PASS + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?parseTime=true"

	database, err := gorm.Open(mysql.Open(sql), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Mysql Database Connected!")
	}

	DB = database
}

func (UserS) TableName() string {
	return UsersTable
}

func (Note) TableName() string {
	return NotesTable
}

func DB_Table_Migrate() {
	DB.AutoMigrate(&UserS{})
	DB.AutoMigrate(&Note{})
}
