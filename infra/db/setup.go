package db

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/labib0x9/ProjectUnsafe/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	_ "github.com/lib/pq"
)

func Setup(cnf *config.DbConfig) error {
	superDb, err := NewSuperConnection(cnf)
	if err != nil {
		return err
	}
	defer superDb.Close()

	_, err = superDb.Exec(fmt.Sprintf(`
		DO $$
		BEGIN
			IF NOT EXISTS (SELECT FROM pg_roles WHERE rolname = '%s') THEN
				CREATE ROLE %s WITH LOGIN PASSWORD '%s';
			END IF;
		END
		$$;
	`, cnf.DBUser, cnf.DBUser, cnf.DBPass))
	if err != nil {
		return err
	}

	var exists bool
	err = superDb.QueryRow(
		`SELECT EXISTS(SELECT FROM pg_database WHERE datname = $1)`,
		cnf.DBName,
	).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		_, err = superDb.Exec(fmt.Sprintf(
			`CREATE DATABASE %s OWNER %s`,
			cnf.DBName, cnf.DBUser,
		))
		if err != nil {
			return err
		}
	}

	_, err = superDb.Exec(fmt.Sprintf(
		`GRANT ALL PRIVILEGES ON DATABASE %s TO %s`,
		cnf.DBName, cnf.DBUser,
	))
	if err != nil {
		return err
	}

	if err := runMigrations(cnf); err != nil {
		return fmt.Errorf("migrations: %w", err)
	}

	fmt.Println("Database setup complete")

	return nil
}

func runMigrations(cnf *config.DbConfig) error {
	dbSource := newConnectionString(cnf)

	appDB, err := sql.Open("postgres", dbSource)
	if err != nil {
		return err
	}
	defer appDB.Close()

	driver, err := postgres.WithInstance(appDB, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

func SetupAndConnection(cnf *config.DbConfig) *sqlx.DB {
	if err := Setup(cnf); err != nil {
		panic(err)
	}

	dbConn, err := NewConnection(cnf)
	if err != nil {
		panic(err)
	}

	return dbConn
}
