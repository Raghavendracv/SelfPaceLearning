package MinElementSum

/*
Url: https://www.geeksforgeeks.org/delete-odd-and-even-numbers-at-alternate-step-such-that-sum-of-remaining-elements-is-minimized/

Delete odd and even numbers at alternate step such that sum of remaining elements is minimized
Given an array arr[] of N elements. At any step, we can delete a number of a different parity from the just previous step, i.e., if, at the previous step, an odd number was deleted then delete an even number in the current step or vice versa.

It is allowed to start by deleting any number. Deletion is possible till we can delete numbers of different parity at every step. The task is to find the minimum possible sum of the elements left at the end.

Examples:

Input: arr[] = {1, 5, 7, 8, 2}
Output: 0
Delete elements in the order 1, 2, 5, 8 and finally 7.
There are multiple ways of deletion,
resulting in the same minimized sum.

Input: arr[] = {2, 2, 2, 2}
Output: 6
Delete 2 in first step.
Cannot delete any number, since there are no odd numbers left.
Hence, the leftover elements sum is 8.

Approaches:

1. Traverse input slice and create two slices for odd and even numbers.
	If count of odd and even number slices are same return 0
	Else If odd slice contains more elements
		sort odd slice in ascending order and return sum of first (Odd Slice Length - Even Slice Length) elements from odd slice.
	Else
		sort even slice in ascending order and return sum of first (even Slice Length - odd Slice Length) elements from even slice.
*/
