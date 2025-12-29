package database

import (
	"fmt"
	"log"

	"github.com/anthoz69/go-temfy/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// ConnectMysql connects to MySQL database
func ConnectMysql(cfg *config.Config) error {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return fmt.Errorf("failed to connect to MySQL database: %w", err)
	}

	if err := configureConnectionPool(); err != nil {
		return err
	}

	log.Println("MySQL database connection established successfully")
	return nil
}

// ConnectPostgres connects to PostgreSQL database
func ConnectPostgres(cfg *config.Config) error {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.Port,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return fmt.Errorf("failed to connect to PostgreSQL database: %w", err)
	}

	if err := configureConnectionPool(); err != nil {
		return err
	}

	log.Println("PostgreSQL database connection established successfully")
	return nil
}

// configureConnectionPool sets up common connection pool settings
func configureConnectionPool() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	// Connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return nil
}

func GetDB() *gorm.DB {
	return DB
}
