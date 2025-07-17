package gorm

import (
	"fmt"
	"sync"

	config "golang-auth-app/config"

	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Module = fx.Module("datasources/sql/gorm", fx.Provide(
	InitDB,
))

var (
	db   *gorm.DB
	once sync.Once
)

func InitDB() *gorm.DB {
	cfg := config.Datasource.Postgres

	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		cfg.Username, cfg.Password, cfg.Host,
		cfg.Port, cfg.DBName, cfg.SSLMode,
	)

	once.Do(func() {
		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(fmt.Errorf("error occurred during connect to database: %w", err))
		}
	})

	return db
}
