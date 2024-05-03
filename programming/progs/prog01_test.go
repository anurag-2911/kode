package progs

import (
	"fmt"
	"log"
	"reflect"
	"sort"
	"testing"
)

func TestXxx(t *testing.T) {
	fmt.Print()
}

type ArrayOps struct {
}

var arrayOps ArrayOps

func init() {
	arrayOps = ArrayOps{}
}

/*
Longest Consecutive Sequence:Given an unsorted array of integers,
find the length of the longest consecutive elements sequence
*/

func TestLongestConsecutiveSequence(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"Basic case", []int{100, 4, 200, 1, 3, 2}, 4},
		{"No consecutive elements", []int{10, 20, 30}, 1},
		{"All elements are the same", []int{2, 2, 2}, 1},
		{"Array with negative numbers", []int{-1, -2, -3, -4}, 4},
		{"Mixed positive and negative numbers", []int{-1, 1, 2, -2, -3, 3, 4}, 4},
		{"Empty array", []int{}, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := arrayOps.longestConsecutiveSequence(test.nums)
			if result != test.expected {
				t.Errorf("expected %v, found %v", test.expected, result)
			}
		})
	}
}

func (arops ArrayOps) longestConsecutiveSequence(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	current, longest := 1, 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			current++
		} else {
			if current > longest {
				longest = current
			}
			current = 1
		}
	}

	if current > longest {
		longest = current
	}

	return longest
}

/*
Find the "Kth" Largest Element in an Array:
Given an integer array and a number k, return the kth largest element in the array.
*/
func TestKthLargestNum(t *testing.T) {
	tests := []struct {
		Name     string
		nums     []int
		k        int
		expected int
	}{
		{"test1", []int{101, 2, 56, 78, 9, 11}, 3, 56},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := arrayOps.kthlargest(test.nums, test.k)
			if result != test.expected {
				log.Printf("expected %v, found %v", test.expected, result)
			}
		})
	}
}
func (aops *ArrayOps) kthlargest(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	return nums[k-1]
}

/*
Move Zeroes:
Given an array nums, write a function to move all 0's to the end of it
while maintaining the relative order of the non-zero elements.
*/

func (aops *ArrayOps) movezeros(nums []int) []int {
	j := 0 // index to put non zero element

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != j {
				nums[j], nums[i] = nums[i], nums[j]
			}
			j++
		}
	}
	return nums
}
func TestMoveZeros(t *testing.T) {
	tests := []struct {
		Name     string
		nums     []int
		expected []int
	}{
		{"test1", []int{0, 12, 0, 34, 2, 3, 4, 0, 0, 67}, []int{12, 34, 2, 3, 4, 67, 0, 0, 0, 0}},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := arrayOps.movezeros(test.nums)
			if !reflect.DeepEqual(result, test.expected) {
				log.Printf("expected %v, got %v", test.nums, result)
			}
		})
	}
}
