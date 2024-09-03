package loaders

import (
	"backend/internal/domain/model"
	"backend/loaders/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {

	dbURL := config.Conf.DBUrl
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		logrus.Fatal("Error connecting to database")
	}

	DB = db
	if err := Migrate(); err != nil {
		logrus.Fatal("Error migrating database")
	}

	logrus.Debugln("INIT DATABASE CONNECTION")

}

func Migrate() error {
	if err := DB.AutoMigrate(
		new(model.User),
		new(model.ChatRoom),
		new(model.Notification),
		new(model.Tracking),
		new(model.Token),
		new(model.Tracking),
		new(model.Pet)); err != nil {
		return err
	}
	logrus.Debugln("Migrated database")
	return nil
}
