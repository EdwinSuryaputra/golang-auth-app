TIMESTAMP = `date +%Y%m%d%H%M%S`

.PHONY: build
build: 
	@BUILD_START=$$(date +%s) \
	;CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/app \
	&& echo "Build took $$(($$(date +%s)-BUILD_START)) seconds"

.PHONY: run
run:
	@go run main.go

.PHONY: generate-sql-migration
generate-sql-migration:
	@touch ./migrations/${TIMESTAMP}_$(name).up.sql && touch ./migrations/${TIMESTAMP}_$(name).down.sql

.PHONY: migrate-sql
migrate-sql:
	@go run ./cmd/migrate/migrate.go

.PHONY: generate-gorm-model
generate-gorm-model:
	@go run ./cmd/gen/gorm/gormgen.go -tags=integration

.PHONY: run-data-seeding
run-data-seeding:
	@go run ./cmd/seed/. -tags=integration