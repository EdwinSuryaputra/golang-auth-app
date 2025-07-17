package dbcontainer

import (
	"context"
	"fmt"
	"net"
	"os"
	"path"
	"runtime"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	gormPostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var dbName = "database"
var dbUser = "user"
var dbPass = "123456"

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	os.Chdir(dir)
	os.Chdir("../../../")
}

func StartContainer() testcontainers.Container {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	c, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		Started: true,
		Reuse:   true,
		ContainerRequest: testcontainers.ContainerRequest{
			Name:         "idm-api-postgres-testcontainer",
			Image:        "postgres:16",
			Hostname:     GetOutboundIP().String(),
			ExposedPorts: []string{"5432/tcp", "5432/tcp"},
			Env: map[string]string{
				"POSTGRES_USER":     dbUser,
				"POSTGRES_PASSWORD": dbPass,
				"POSTGRES_DB":       dbName,
			},
			WaitingFor: wait.ForListeningPort("5432/tcp"),
		},
	})
	if err != nil {
		panic(err)
	}

	return c
}

func StopContainer() {
	c := StartContainer()
	c.Terminate(context.Background())
}

func GetGorm() (*gorm.DB, error) {
	if db != nil {
		return db, nil
	}

	container := StartContainer()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	host, err := container.Host(ctx)
	if err != nil {
		return nil, err
	}

	port, err := container.MappedPort(ctx, "5432")
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, dbUser, dbPass, dbName, port.Port(),
	)

	gormdb, err := gorm.Open(gormPostgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	d, err := gormdb.DB()
	if err != nil {
		return nil, err
	}
	d.SetMaxIdleConns(0)

	driver, err := postgres.WithInstance(d, &postgres.Config{DatabaseName: dbName})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		return nil, err
	}

	m.Up()
	db = gormdb
	return gormdb, nil
}

// GetOutboundIP Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
