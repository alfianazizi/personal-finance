package database

import (
	"fmt"
	"log"
	"sync"

	"github.com/alfianazizi/personal-finance/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDatabase struct {
	Db *gorm.DB
}

var (
	once       sync.Once
	dbInstance *postgresDatabase
)

func NewPostgresDatabase(conf *config.Config) Database {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s Timezone=%s",
			conf.Db.Host,
			conf.Db.Port,
			conf.Db.User,
			conf.Db.Password,
			conf.Db.DBName,
			conf.Db.SSLMode,
			conf.Db.Timezone,
		)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}
		dbInstance = &postgresDatabase{Db: db}
	})
	return dbInstance
}

func (p *postgresDatabase) GetDB() *gorm.DB {
	return dbInstance.Db
}
