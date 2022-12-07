package goronext

import (
	"testing"
)

func TestFindStock(t *testing.T) {
	s, err := findAllStock()
	if err != nil {
		t.Error(err)
	}
	if len(s) == 0 {
		t.Error("No stock found")
	}
	if s[0].ID != "FR0010285965" {
		t.Error("Wrong ID")
	}
	if s[0].Symbol != "ALMIL" {
		t.Error("Wrong Symbol")
	}
	if s[54].ID != "FR0013384369" {
		t.Error("Wrong ID")
	}
	if s[54].Symbol != "ALBKK" {
		t.Error("Wrong Symbol")
	}
}
