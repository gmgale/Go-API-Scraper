package main

import "testing"

// TestGetTitle is a unit test to simulate iterations of user inputted
// threadsfrom 1 to len(urls) and comfirm correct ammount of titles are
// returned.
func TestGetTitle(t *testing.T) {
	for i := 1; i <= len(urls); i++ {
		threads := i

		// Arrange
		titles, _ := getTitle(threads)

		// Act
		if len(titles) < len(urls) {

			// Assert
			t.Errorf("%d titles returned, expected %d.", len(titles), len(urls))
		}
		if len(titles) > len(urls) {
			t.Errorf("%d titles returned, expected %d.", len(titles), len(urls))
		}
	}
}
