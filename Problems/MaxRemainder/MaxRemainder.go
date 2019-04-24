package MaxRemainder

/*
Url: https://www.geeksforgeeks.org/maximum-modulo-pairs-array-arri-arrj/

Maximum modulo of all the pairs of array where arr[i] >= arr[j]
Given an array of n integers. Find the maximum value of arr[i] mod arr[j] where arr[i] >= arr[j] and 1 <= i, j <= n
Examples:

Input: arr[] = {3, 4, 7}
Output: 3
Explanation:
There are 3 pairs which satisfies arr[i] >= arr[j] are:-
4, 3 => 4 % 3 = 1
7, 3 => 7 % 3 = 1
7, 4 => 7 % 4 = 3
Hence Maximum value among all is 3.

Input: arr[] = {3, 7, 4, 11}
Output: 4

Input: arr[] = {4, 4, 4}
Output: 0

Approaches:

1. Brute fore way of finding reminder of all permutations & combinations.

2. Find largest two elements and return second largest number % largest number.

*/

// CalculateMaxRemainder returns maximum reminder of two elements from input slice
func CalculateMaxRemainder(input []int) int {

	n := len(input)
	largest := 0
	secondLargest := 0

	limit := (n / 2)
	if n%2 != 0 {
		limit++
	}

	for i := 0; i < limit; i++ {

		if input[i] > largest {
			secondLargest = largest
			largest = input[i]
		}

		if input[i] > secondLargest && input[i] < largest {
			secondLargest = input[i]
		}

		if input[n-i-1] > largest {
			secondLargest = largest
			largest = input[n-i-1]
		}

		if input[n-i-1] > secondLargest && input[n-i-1] < largest {
			secondLargest = input[n-i-1]
		}
	}

	return secondLargest % largest
}
