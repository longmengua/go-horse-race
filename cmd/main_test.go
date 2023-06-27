package main_test

import (
	structs "go-horse-race/structs"
	"testing"
)

/*
1. create a number of horse match race
2.
*/
func TestSlice(t *testing.T) {
	m := structs.NewMatch("Match", 10, 100)
	m.Start()
	if m.Winner == nil {
		t.Errorf("Failed")
	}
}
