package main

import (
	"context"
	"fmt"
	"time"

	"golang-auth-app/app/datasources/sql/gorm"
	config "golang-auth-app/config"

	"gorm.io/gorm/clause"
)

type upsertPayload[T any] struct {
	Data              []*T
	OnConflictColumns []clause.Column
	DoUpdateColumns   []string
}

func main() {
	config.Init()
	ctx := context.Background()

	now := time.Now()
	creator := "cli"

	if err := seedResource(ctx, now, creator); err != nil {
		panic(err)
	}

	fmt.Println("Seeding success!")
}

func upsertBatches[T any](ctx context.Context, payload upsertPayload[T]) error {
	db := gorm.InitDB()

	err := db.WithContext(ctx).
		Clauses(
			clause.OnConflict{
				Columns:   payload.OnConflictColumns,
				DoUpdates: clause.AssignmentColumns(payload.DoUpdateColumns),
			},
		).
		CreateInBatches(payload.Data, len(payload.Data)).Error
	if err != nil {
		return err
	}

	return nil
}
