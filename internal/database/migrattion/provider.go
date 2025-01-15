package migrattion

import "LoggingApp/pkg/base_migrate"

func getMigrations(m base_migrate.Migrate) []base_migrate.Migration {
	return []base_migrate.Migration{
		UserMigration{Migrate: m},
		PostMigration{Migrate: m},
	}
}
