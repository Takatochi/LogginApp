package base_migrate

import (
	"LoggingApp/pkg/database"
)

type Migrate interface {
	Migrate(dst ...interface{}) error
}
type migrate struct {
	db database.Database
}

func NewMigrate(store database.Database) Migrate {
	return migrate{
		db: store,
	}
}

func (m migrate) Migrate(dst ...interface{}) error {
	err := m.db.GetDB().Migrator().AutoMigrate(dst...)
	if err != nil {
		return err
	}
	return nil
}

type Migration interface {
	Up() error
}
