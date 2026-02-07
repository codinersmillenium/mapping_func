package main

import (
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
	"strings"

	"mapping_func/config"
)

func migrateAll(db *config.Database, folder string) error {
	files := []string{}
	err := filepath.Walk(folder, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".sql") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return err
	}

	sort.Strings(files)

	for _, f := range files {
		log.Println("Running migration:", f)
		sqlBytes, err := ioutil.ReadFile(f)
		if err != nil {
			return err
		}
		_, err = db.Exec(string(sqlBytes))
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	cfg := config.Load()

	db, err := config.NewDatabase(cfg.DBDriver, cfg.DSN)
	if err != nil {
		log.Fatal(err)
	}

	if err := migrateAll(db, cfg.MigrationsPath); err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("All migrations completed!")
}
