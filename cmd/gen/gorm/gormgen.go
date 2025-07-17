package main

import (
	"golang-auth-app/app/datasources/sql/dbcontainer"

	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./app/datasources/sql/gorm/query",
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		FieldNullable: true,
	})

	gormdb, err := dbcontainer.GetGorm()
	if err != nil {
		panic(err)
	}

	g.UseDB(gormdb)

	for _, table := range g.GenerateAllTable() {
		g.ApplyBasic(table)
		g.ApplyInterface(func() {}, table)
	}

	g.Execute()
}
