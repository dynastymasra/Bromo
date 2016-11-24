package settings

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Connector is function for connect to database
func Connector() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", os.Getenv("DATABASE_DATA"))

	if err != nil {
		log.Errorf("database_connector - Connector : %v", err)
		return nil, err
	}

	return db, nil
}
