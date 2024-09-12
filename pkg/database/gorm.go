package database

import (
	//"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"

	//"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"go_tutorial/internal/domain"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	Debug    bool
}

// New ...
func New(conf Config) (db *gorm.DB, err error) {
	// DSN (Data Source Name) without the database name for initial connection
	//dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
	//	conf.Host,
	//	conf.User,
	//	conf.Password,
	//	conf.DBName,
	//	conf.Port,
	//)

	var newLogger logger.Interface
	if conf.Debug {
		newLogger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second, // Slow SQL threshold
				LogLevel:      logger.Info, // Log level
				Colorful:      true,        // Disable color
			},
		)
	}

	// Open the initial connection
	db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return
	}

	return
}

func Migrate(db *gorm.DB) (err error) {
	return db.AutoMigrate(
		&domain.User{},
	)
}
