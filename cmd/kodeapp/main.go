package main

import "fmt"

// Function to determine if it's possible to read all chapters within 'days' days when reading 'x' pages per day.
func canFinish(pages []int32, days int32, x int32) bool {
	dayCount := int32(1) // We start on the first day
	pagesReadToday := int32(0)

	for _, pagesInChapter := range pages {
		if pagesInChapter > x { // We can never read more than 'x' pages in a day
			return false
		}
		if pagesReadToday+pagesInChapter > x { // If we cannot read the next chapter today, we'll read it tomorrow
			dayCount++
			pagesReadToday = pagesInChapter
		} else { // If we can read the next chapter today, we'll do it
			pagesReadToday += pagesInChapter
		}
	}

	return dayCount <= days
}

// Function to find the minimum number of pages 'x' to read per day in order to finish all chapters within 'days' days.
func minimumNumberOfPages(pages []int32, days int32) int32 {
	// Setting up the search space for 'x'
	low := int32(1) // The minimum possible value for 'x' is 1
	high := int32(0) // The maximum is the sum of all pages
	for _, p := range pages {
		high += p
	}

	// Binary search for the minimum 'x'
	var mid int32
	for low < high {
		mid = low + (high-low)/2
		if canFinish(pages, days, mid) {
			high = mid // Try a smaller 'x'
		} else {
			low = mid + 1 // Try a larger 'x'
		}
	}

	// The final check to see if 'low' allows us to read all chapters in 'days' days
	if canFinish(pages, days, low) {
		return low
	} else {
		return -1 // It's not possible to finish reading all chapters within the given days
	}
}

func main() {
	pages := []int32{2, 3, 4, 5}
	days := int32(5)
	minimumPages := minimumNumberOfPages(pages, days)
	fmt.Printf("The minimum number of pages to read each day is: %d\n", minimumPages)
}
