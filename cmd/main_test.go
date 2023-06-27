package main_test

import (
	"testing"
)

func assertion(t *testing.T, assert any, result any) {
	if assert != result {
		t.Errorf("Error, %s, %s", result, assert)
	}
}

func TestSlice(t *testing.T) {
	assertion(t, true, true)
}
