package main

import (
	"testing"

	"github.com/pedroluis02/design-patterns-golang-s1/singleton"
)

func TestSingleton(t *testing.T) {
	connection := singleton.GetDatabaseConnection()
	status := connection.GetStatus()
	want := singleton.Success

	if status != want {
		t.Errorf("got %s, wanted %q", status, want)
	}
}
