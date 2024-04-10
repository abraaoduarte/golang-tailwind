package database

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

//func RunMigration() {
//	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=docker dbname=postgres sslmode=disable")
//	if err != nil {
//		fmt.Println("error dbection")
//		panic(err)
//	}
//	defer db.Close()
//
//	driver, err := postgres.WithInstance(db, &postgres.Config{})
//	if err != nil {
//		fmt.Println("error with driver")
//		panic(err)
//	}
//
//	migration, err := migrate.NewWithDatabaseInstance("file://internal/database/migrations", "postgres", driver)
//	if err != nil {
//		panic("problem to run migrations")
//	}
//	migration.Up()
//	fmt.Println("migrations ran with success")
//}

func RunMigration() {
	m, err := migrate.New(
		"file://internal/database/migrations",
		"postgres://postgres:docker@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}

	fmt.Println("migrations run")
	m.Up()
}
