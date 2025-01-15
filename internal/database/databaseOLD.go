package databaseOLD

//
//import (
//	"LoggingApp/internal/config"
//	"LoggingApp/pkg/database"
//	"database/sql"
//)
//
//type Database interface {
//	Open() (*sql.DB, error)
//}
//type dbBase struct {
//	cfg        config.Config
//	driverName string
//}
//type sqlDatabase struct {
//	*dbBase
//}
//
//func (d sqlDatabase) Open() (*sql.DB, error) {
//
//	db, err := database.Conn(d.driverName, d.cfg.ConnectionStrings)
//	if err != nil {
//		return nil, err
//	}
//	return db, nil
//}
//
//type PgDatabase struct {
//	sqlDatabase
//}
//type MySQLDatabase struct {
//	sqlDatabase
//}
//func New(driver string, cfg config.Config) Database {
//	db := &dbBase{
//		cfg:        cfg,
//		driverName: driver,
//	}
//	switch driver {
//	case "postgres":
//		return &PgDatabase{sqlDatabase{db}}
//	case "mysql":
//
//		return &MySQLDatabase{sqlDatabase{db}}
//	default:
//
//		return nil
//	}
//}
//func Conn(DriverName, DatabaseURL string) (*sql.DB, error) {
//	db, err := sql.Open(DriverName, DatabaseURL)
//	if err != nil {
//		return nil, err
//	}
//
//	if err := db.Ping(); err != nil {
//		return nil, err
//	}
//
//	db.SetConnMaxLifetime(time.Minute * 10)
//	db.SetMaxOpenConns(50)
//	db.SetMaxIdleConns(20)
//
//	return db, nil
//}
