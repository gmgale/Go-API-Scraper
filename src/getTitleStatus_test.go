package main

import "testing"

// TestGetTitleStatusCount is a unit test to simulate iterations of user inputted
// threads from 1 to len(urls) and comfirm correct number of status (succeed/fail)
// calls count is returned.
func TestGetTitleStatusCount(t *testing.T) {
	threads := 1

	// Arrange
	titleData := getTitle(threads)
	statusSuccess := titleData.status.success
	statusFail := titleData.status.fail

	// Act
	if (statusSuccess + statusFail) != len(urls) {

		// Assert
		t.Errorf("%d succeeded and %d failed responces returned totaling %d, expected a total of %d.",
			statusSuccess, statusFail, (statusSuccess + statusFail), len(urls))
	}
}
