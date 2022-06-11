package app_test

import (
	"server/app"
	"testing"
)

/*
Test app config
*/
func TestAppConfig(t *testing.T) {
	new_config := app.Config{
		Port:   "8080",
		DBHost: "localhost",
		DBUser: "root",
		DBPass: "root",
		DBName: "test",
	}

	if new_config.Port != "8080" {
		t.Error("Port is not correct")
	}

	if new_config.DBHost != "localhost" {
		t.Error("DB_Host is not correct")
	}

	if new_config.DBHost != "root" {
		t.Error("DB_User is not correct")
	}

	if new_config.DBPass != "root" {
		t.Error("DB_Pass is not correct")
	}

	if new_config.DBName != "test" {
		t.Error("DB_Name is not correct")
	}
}
