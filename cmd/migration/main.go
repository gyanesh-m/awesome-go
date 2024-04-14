package main

import (
	"awesome-go/internal/boot"
	"awesome-go/internal/configs"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"os"
	"path"
)

func main() {
	boot.Setup()
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.Config.Db.User, configs.Config.Db.Password, configs.Config.Db.Host, configs.Config.Db.Port, configs.Config.Db.DbName)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("failed to create driver: %v", err)
	}
	dir, _ := os.Getwd()
	migrationPath := "file://" + path.Join(dir, "internal/db/migrations")
	m, err := migrate.NewWithDatabaseInstance(
		migrationPath,
		configs.Config.Db.DbName, driver)
	if err != nil {
		boot.Lgr.Fatalf("failed to create migration instance: %v", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		boot.Lgr.Fatalf("failed to apply migrations: %v", err)
	}

	boot.Lgr.Info("all migrations applied successfully")
}
