package database

import (
	"log"

	"github.com/gurtinho/video_encoder/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/lib/pq"
)

type Database struct {
	Db            *gorm.DB
	Dsn           string
	DsnTest       string
	DbType        string
	DbTypeTest    string
	Debug         bool
	AutoMigrateDB bool
	Env           string
}

func NewDB() *Database {
	return &Database{}
}

func NewDBTest() *gorm.DB {
	dbinstance := NewDB()
	dbinstance.Env = "test"
	dbinstance.DbTypeTest = "sqlite3"
	dbinstance.DsnTest = ":memory:"
	dbinstance.AutoMigrateDB = true
	dbinstance.Debug = true

	connection, err := dbinstance.Connect()
	if err != nil {
		log.Fatalf("Test db error %v", err)
	}
	return connection
}

func (db *Database) Connect() (*gorm.DB, error) {
	var err error

	if db.Env == "test" {
		db.Db, err = gorm.Open(db.DbTypeTest, db.DsnTest)
	} else {
		db.Db, err = gorm.Open(db.DbType, db.Dsn)
	}

	if err != nil {
		return nil, err
	}
	if db.Debug {
		db.Db.LogMode(true)
	}
	if db.AutoMigrateDB {
		db.Db.AutoMigrate(&domain.Video{}, &domain.Job{})

		db.Db.Model(&domain.Job{}).AddForeignKey("video_id", "videos(id)", "cascade", "cascade")

	}
	return db.Db, nil
}
