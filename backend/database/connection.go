package database

import (
	"fmt"
	"log"

	"github.com/HiteshKumarMeghwar/L-M-S-2/models"
	_ "github.com/jackc/pgx/v4/stdlib" // Import the pgx driver
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SqlVariables struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (cfg SqlVariables) String() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
}

var DB *gorm.DB // Use gorm.DB for connection

func Connect() {
	cfg := SqlVariables{
		Host:     "localhost",
		Port:     "3306",
		User:     "root",
		Password: "",
		Database: "learning_management_system",
	}

	dsn := cfg.String()

	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		log.Println("Database Successfully Connected....!")
	}

	DB = connection

	// AutoMigrate the User model and any other models you need
	connection.AutoMigrate(
		&models.User{},
	)
}
