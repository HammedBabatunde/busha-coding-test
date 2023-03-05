package migrations

import (
	pgMigrator "github.com/emekarr/coding-test-busha/migrations/postgres"
)

func RunMigrations() {
	pgMigrator.Migrate()
}
