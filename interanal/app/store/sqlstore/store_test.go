package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("database_url")
	if databaseURL == "" {
		databaseURL = "postgres://postgres:211217ns@localhost/restapi_test?sslmode=disable"
	}

	os.Exit(m.Run())
}
