package ciorba

import (
	"testing"
)

func TestB64ToHexUID(t *testing.T) {
	given1 := "rB4An1Z3WRwAwU3G_E4fAg"
	expected1 := "9F001EAC1C597756C64DC100021F4EFC"
	actual1, err := B64ToHexUID(given1)
	if err != nil {
		t.Error(err)
		t.Fail()
	} else if expected1 != actual1 {
		t.Logf("%s != %s", expected1, actual1)
		t.Fail()
	}

	// Non URL-safe
	given2 := "rB4An1Z3WRwAwU3G/E4fAg=="
	expected2 := "9F001EAC1C597756C64DC100021F4EFC"
	actual2, err := B64ToHexUID(given2)
	if err != nil {
		t.Error(err)
	} else if expected2 != actual2 {
		t.Logf("\n%s\n%s\n", expected2, actual2)
		t.Fail()
	}
}
