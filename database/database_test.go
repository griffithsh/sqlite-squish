package database

import "testing"

func TestFromString(t *testing.T) {
	d, err := FromString("")
	if len(d.Tables) > 0 {
		t.Errorf("Unexpected length %d", len(d.Tables))
	}
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}
}
