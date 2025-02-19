package config 

import (
	"testing"
)

func TestReadConfig(t *testing.T) {

	expected := struct {
		db_url string
	} {
		db_url: "postgres://example",
	}

	actual, err := Read()

	if err != nil {
		t.Errorf(
`
--- TEST FAILED ---
Error in Read(): %s
`, err)
	}

	if actual.DBUrl != expected.db_url {
t.Errorf(
`
--- TEST FAILED ---
Unexpected db_url in actual
expected db_url: %s
actual db_url: %s
`, expected.db_url, actual.DBUrl)		
	}
}