package control_structures

import "testing"

func TestSeason(t *testing.T) {
	_, err := Season(2)
	if err != nil {
		t.Fatal("test failed")
	}
}
