package main

import "testing"

// TestGetTitle is a unit test to simulate itterations of user inputted
// threadsfrom 1 to len(urls).
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
