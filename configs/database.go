package configs

import (
	log "github.com/sirupsen/logrus"
	"goers/pkg/event/model/entity"
	entity2 "goers/pkg/order/model/entity"
	entity3 "goers/pkg/user/model/entity"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(config Configuration) *gorm.DB {
	DBName := config.Get("DB_NAME")
	DBUser := config.Get("DB_USER")
	DBPassword := config.Get("DB_PASS")
	DBHost := config.Get("DB_HOST")

	dsn := "host=" + DBHost + " port=5432" + " user=" + DBUser + " password=" + DBPassword + " dbname=" + DBName + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Error("Failed to connect to database. \n", err)
	}
	log.Println("Connected to database.")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migration...")
	err = db.AutoMigrate(
		&entity.Event{}, &entity.Ticket{},
		&entity2.Order{},
		&entity3.User{},
	)
	if err != nil {
		log.Fatal("Failed to run migration. \n", err)
	} else {
		log.Println("Migration completed.")
	}
	return db
}

func NewDatabaseTest() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Error("Failed to connect to database. \n", err)
	}
	log.Println("Connected to database.")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migration...")
	err = db.AutoMigrate(
		&entity.Event{}, &entity.Ticket{},
		&entity2.Order{},
		&entity3.User{},
	)
	if err != nil {
		log.Fatal("Failed to run migration. \n", err)
	} else {
		log.Println("Migration completed.")
	}
	return db
}
