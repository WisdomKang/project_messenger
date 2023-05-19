package database_test

import (
	"messenger_server/pkg/database"
	"testing"
)

func TestDatabaseInit(t *testing.T) {
	dbservice := database.GetInstance()

	if dbservice == nil {
		t.Log("fail")
	}
	t.Log("test")
}
