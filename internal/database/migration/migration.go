package migration

import (
	"fmt"

	"github.com/flastors/songius/pkg/client/postgresql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func NewMigration(sc postgresql.StorageConfig) (*migrate.Migrate, error) {
	m, err := migrate.New(
		"file://db/migrations",
		fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", sc.Username, sc.Password, sc.Host, sc.Port, sc.Database),
	)
	if err != nil {
		return nil, err
	}
	return m, nil

}
