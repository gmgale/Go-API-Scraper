package main

import "testing"

// TestGetTitleStatusCount is a unit test to simulate iterations of user inputted
// threads from 1 to len(urls) and comfirm correct number of status (succeed/fail)
// calls count is returned.
func TestGetTitleStatusCount(t *testing.T) {
	for i := 1; i <= len(urls); i++ {
		threads := i

		// Arrange
		_, status := getTitle(threads)

		// Act
		if (status[0] + status[1]) != len(urls) {

			// Assert
			t.Errorf("%d succeeded and %d failed responces returned totaling %d, expected a total of %d.",
				status[0], status[1], (status[0] + status[1]), len(urls))
		}
	}
}
