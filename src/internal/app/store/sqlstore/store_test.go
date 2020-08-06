package sqlstore_test

import (
	"os"
	"testing"
)

var databaseURL string

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		databaseURL = "host=full_db_postgres port=5432 user=admin password=secret dbname=restapi_test sslmode=disable"
	}

	os.Exit(m.Run())
}
