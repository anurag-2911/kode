package progs

import (
	"fmt"
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
	if len(nums)==0{
		return 0
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i]<nums[j]
	})
	current,longest:=1,1

	for i:=1;i<len(nums);i++{
		if nums[i]==nums[i-1]+1{
			current++
		}else{
			if current>longest{
				longest=current
			}
			current=1
		}
	}

	if current>longest{
		longest=current
	}

	return longest
}

/*
Two Sum Problem:
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
*/

