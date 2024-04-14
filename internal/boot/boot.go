package boot

import (
	"awesome-go/internal/configs"
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
)

var Lgr *zap.SugaredLogger
var Db *gorm.DB

func Setup() {
	// Load configs
	configs.Load()

	// Load Logger
	InitLogger()

	// Load Db
	InitDB()
}

func InitLogger() {
	// Load logger
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	Lgr = logger.Sugar()
}

func InitDB() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.Config.Db.User, configs.Config.Db.Password, configs.Config.Db.Host, configs.Config.Db.Port, configs.Config.Db.DbName)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	gormDialect := mysql.New(mysql.Config{
		Conn: db,
	})
	gormDb, err := gorm.Open(gormDialect, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	Db = gormDb.Debug()
}
