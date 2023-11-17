package database_test

import (
	"messenger_server/pkg/database"
	"testing"
)

func TestDatabaseInit(t *testing.T) {
	// Do not use cache on test.
	t.Parallel()
	database.InitDB()

	db := database.GetDB()

	if db != nil {
		t.Fail()
	}
}
