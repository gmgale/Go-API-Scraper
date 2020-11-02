package main

import "testing"

// TestGetTitle is a unit test to simulate iterations of user inputted
// threadsfrom 1 to len(urls) and comfirm correct ammount of titles are
// returned.
func TestGetTitle(t *testing.T) {
	for i := 1; i <= len(urls); i++ {
		threads := i
		titles, _ := getTitle(threads)
		if len(titles) < len(urls) {
			t.Errorf("%d titles returned, expected %d.", len(titles), len(urls))
		}
		if len(titles) > len(urls) {
			t.Errorf("%d titles returned, expected %d.", len(titles), len(urls))
		}
	}
}

// TestGetTitleStatusCount is a unit test to simulate iterations of user inputted
// threads from 1 to len(urls) and comfirm correct number of status (succeed/fail)
// calls count is returned.
func TestGetTitleStatusCount(t *testing.T) {
	for i := 1; i <= len(urls); i++ {
		threads := i
		_, status := getTitle(threads)
		if (status[0] + status[1]) != len(urls) {
			t.Errorf("%d succeeded and %d failed responces returned totaling %d, expected a total of %d.",
				status[0], status[1], (status[0] + status[1]), len(urls))
		}
	}
}
