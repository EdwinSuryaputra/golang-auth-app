package main

import (
	"database/sql"
	"fmt"

	config "golang-auth-app/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	config.Init()
	connProfile := config.Datasource.Postgres

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s port=%d dbname=%s sslmode=disable",
		connProfile.Host, connProfile.Username, connProfile.Password, connProfile.Port, connProfile.DBName,
	)

	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		panic(err)
	}
	defer driver.Close()

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		panic(err)
	}
	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			fmt.Println(err)
			return
		} else {
			fmt.Printf("%v", err)
			return
		}
	}

	fmt.Println("Migrations applied successfully!")
}
