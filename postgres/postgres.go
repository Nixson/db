package postgres

import (
	"fmt"
	"github.com/Nixson/db"
	"github.com/Nixson/environment"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var gormInstance *gorm.DB

func Get() *gorm.DB {
	if gormInstance == nil {
		panic("not init DB")
	}
	return gormInstance
}

func InitDb() {
	env := environment.GetEnv()
	logLevel := logger.Silent
	if env.GetBool("db.showSql") {
		logLevel = logger.Info
	}
	newLogger := logger.New(
		log.New(&db.LoggerWriter{LogLevel: logLevel}, "", 0), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logLevel,    // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	db, err := gorm.Open(
		postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d search_path=%s sslmode=%s",
			env.GetString("db.host"),
			env.GetString("db.login"),
			env.GetString("db.password"),
			env.GetString("db.name"),
			env.GetInt("db.port"),
			env.GetString("db.schema"),
			env.GetString("db.ssl")),
		),
		&gorm.Config{
			Logger: newLogger,
		})
	if err != nil {
		panic("failed to connect database")
	}

	gormInstance = db
}
