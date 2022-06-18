package geoutil_test

import (
	"testing"

	"github.com/airdb/sailor/geoutil"
)

func TestGetByNumeric(t *testing.T) {
	code, _ := geoutil.GetByNumeric(840)

	if code.Name != "United States" {
		t.Fatalf("GetByNumeric failed")
	}
}
