package configs

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	DB *gorm.DB
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	gormConfig := &gorm.Config{
		// enhance performance config
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		Logger:                 dbLogger,
	}

	dsn := os.Getenv("DB_DSN")

	db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("success connect to database")

	app := &Config{
		DB: db,
	}

	return app
}
