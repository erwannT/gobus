package utils

import "testing"

func TestAdd(t *testing.T) {
	result, _ := Add("14:00:00", 5)

	if result != "14:05:00" {
		t.Fail()
	}
}

func TestAdd_failure(t *testing.T) {
	result, _ := Add("14:00:00", 6)

	if result != "14:05:00" {
		t.Fail()
	}
}
