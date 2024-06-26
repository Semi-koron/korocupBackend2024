package database

import (
	"fmt"

	"github.com/semikoron/korocupbackend/utils/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConnectDB() {
	dsn := "host=db user=" + config.PostgreSQLUser + " password=" + config.PostgreSQLPassword + " dbname=" + config.PostgreSQLDBName + " port=" + config.PostgreSQLPort + " " + config.PostgreSQLConfig
	fmt.Println(dsn)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	DB.AutoMigrate(&User{}, &Post{}, &Like{})
}
