package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
	"time"
)

var DB *gorm.DB

//var Db *sql.DB

var _ error

func InitialMigrationForStaging() {

	dsn := "root:sifat@tcp(localhost:3306)/event_management?tls=false&parseTime=false&loc=Asia%2FDhaka&multiStatements=true"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Silent),
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		fmt.Println("failed to connect db, got error: %v", err)
		//	InitialMigrationForStaging()
	}

	if err := db.Use(dbresolver.Register(dbresolver.Config{
		Sources: []gorm.Dialector{mysql.Open(dsn)},
		//Replicas: []gorm.Dialector{mysql.Open(dsn1), mysql.Open(dsn2)},

		// sources/replicas load balancing policy
		Policy: dbresolver.RandomPolicy{},
	}).SetMaxIdleConns(10).
		SetMaxOpenConns(50).
		SetConnMaxIdleTime(time.Hour).
		SetConnMaxLifetime(15 * time.Minute)); err != nil {
		fmt.Println("failed to use plugin, got error: %v", err)
	}
	DB = db

}
