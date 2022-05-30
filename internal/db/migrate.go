package db

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func GooseExec(cfg DatabaseConfig, args []string) {
	dir := cfg.GetMigrationsPath()
	connString := cfg.GetConnString()
	command := args[0]
	fmt.Println(connString)
	fmt.Println(dir)
	db, err := goose.OpenDBWithDriver("postgres", connString)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	arguments := make([]string, 0)
	if len(args) > 1 {
		arguments = append(arguments, args[1:]...)
	}

	if err := goose.Run(command, db, dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}

func Migrate(cfg DatabaseConfig) {
	GooseExec(cfg, []string{"up"})
}
