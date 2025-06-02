package main

import "testing"

func TestMedia(t *testing.T) {
	expectedResut := 5
	result := media(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	if result != expectedResut {
		t.Errorf("Expected media %d but got %d", expectedResut, result)
	}

}
