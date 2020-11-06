package main

import "testing"

// TestGetTitle is a unit test to simulate user inputted
// threads from comfirm correct ammount of titles are
// returned.
func TestGetTitle(t *testing.T) {
	threads := 4

	// Arrange
	titleData := getTitle(threads)
	titles := titleData.Results.Title

	// Act
	if len(titles) < len(urls) {

		// Assert
		t.Errorf("%d titles returned, expected %d.", len(titles), len(urls))
	}
	if len(titles) > len(urls) {
		t.Errorf("%d titles returned, expected %d.", len(titles), len(urls))
	}

}
