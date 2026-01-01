package migrate

import (
	"context"
	"database/sql"
	"os"
)

func Run(ctx context.Context, db *sql.DB, schemaDir string) error {
	files, err := LoadSchemaFiles(schemaDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		sqlBytes, err := os.ReadFile(file)
		if err != nil {
			return err
		}

		if _, err := db.ExecContext(ctx, string(sqlBytes)); err != nil {
			return err
		}
	}
	return nil
}