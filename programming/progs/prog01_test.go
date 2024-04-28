package progs

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

/*
 1. Find the Maximum and Minimum Element in an Array-
    Write a function that returns the maximum and minimum numbers in an array.

2. Reverse an Array - Implement a program to reverse an array or a string.

 3. Sort an Array of 0s, 1s, and 2s - Given an array consisting only of 0s, 1s, and 2s,
    sort the array without using any sorting algorithm.

 4. Find the "Kth" Max and Min Element of an Array -
    Given an array and a number k, find both the kth largest and the kth smallest elements in the array.

 5. Move all the Negative Elements to One Side of the Array -
    Rearrange the array elements so that all negative numbers appear on the left, and positive numbers appear on the right.

 6. Find the Duplicate in an Array of N+1 Integers -
    Assume there is exactly one duplicate number in the array, find the duplicate one.

 7. Merge Two Sorted Arrays Without Using Extra Space -
    Merge two sorted arrays into a single sorted array without using extra space for another array.

8. Kadane's Algorithm - Write a function that returns the maximum sum contiguous subarray within a one-dimensional numeric array.

 9. Find the Union and Intersection of Two Sorted Arrays - Given two sorted arrays,
    find their union and intersection.

10. Cyclically Rotate an Array by One - Given an array, cyclically rotate the array clockwise by one index.
*/
func TestXxx(t *testing.T) {
	fmt.Print()
}

func TestMinMax(t *testing.T) {
	tests := []struct {
		Name string
		arr  []int
		min  int
		max  int
	}{
		{"test1", []int{11, 2, 6, 8, 25, 13}, 2, 25},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			max, min := minmax(test.arr)
			if max != test.max || min != test.min {
				t.Errorf("expected %v,%v,got ", max, min)
			}
		})
	}
}
func minmax(arr []int) (int, int) {
	min := math.MaxInt32
	max := math.MinInt32

	for _, num := range arr {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	return max, min
}

func TestRevArray(t *testing.T) {
	tests := []struct {
		Name     string
		arr      []int
		expected []int
	}{
		{"test1", []int{2, 4, 6, 8, 10}, []int{10, 8, 6, 4, 2}},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := revArray(test.arr)
			if !reflect.DeepEqual(result, test.expected) {
				fmt.Printf("expected %v and got %v", test.expected, result)
			}
		})
	}
}
func revArray(arr []int) []int {

	j := len(arr) - 1

	for i := 0; i < len(arr)/2; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}

func TestSort012(t *testing.T) {
	tests := []struct {
		Name     string
		arr      []int
		expected []int
	}{
		{"test1", []int{2, 0, 1, 0, 1, 1, 0, 0, 2, 2}, []int{0, 0, 0, 0, 1, 1, 1, 2, 2, 2}},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := sort012(test.arr)
			if !reflect.DeepEqual(result, test.expected) {
				fmt.Printf("expected %v,found %v", result, test.expected)
			}
		})
	}

}

/*
Variables: The variables low, mid, and high are pointers used to separate the elements:
low is the position where the next 0 should be placed.
mid is the current element being examined.
high is the position where the next 2 should be placed.
Switch Case Logic:
When nums[mid] is 0, swap it with nums[low] (as 0 should be in the front), then increment both low and mid.
When nums[mid] is 1, just move to the next element by incrementing mid.
When nums[mid] is 2, swap it with nums[high] (as 2 should be at the end) and decrement high, but do not move mid because the swapped element needs to be checked in the next iteration.
Loop Condition: The loop continues until mid surpasses high.
*/
func sort012(arr []int) []int {
	low, mid, high := 0, 0, len(arr)-1

	for mid <= high {
		switch arr[mid] {
		case 0:
			{
				arr[low], arr[mid] = arr[mid], arr[low]
				low++
				mid++
			}
		case 1:
			{
				mid++
			}
		case 2:
			{
				arr[mid], arr[high] = arr[high], arr[mid]
				high--
			}
		}
	}
	return arr
}

func TestRevStr(t *testing.T) {
	tests := []struct {
		Name     string
		str      string
		expected string
	}{
		{"test1", "hallo", "ollah"},
	}
	sops := &StringOps{}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := sops.revstr(test.str)
			if reflect.DeepEqual(result, test.expected) {
				fmt.Println("not matching")
			}
		})
	}
}

type StringOps struct{}

func (sops *StringOps) revstr(str string) string {
	runes := []rune(str)

	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func checkPalin(str string) bool {
	runes := []rune(str)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

func isAnagram(str1 string, str2 string) bool {
	words := make(map[rune]int)

	for _, char := range str1 {
		words[char]++
	}

	for _, char := range str2 {
		if _, exists := words[char]; exists {
			words[char]--
			if words[char] == 0 {
				delete(words, char)
			}
		} else {
			return false
		}
	}

	return len(words) == 0
}

func firstNonRepeatingChar(str string)(rune,bool){
	words:=make(map[rune]int)
	runes:=[]rune(str)

	for _,val:=range runes{
		words[val]++
	}

	for _,val:=range runes{
		if words[val]==1{
			return val,true
		}
	}

	return 0,false
}

