package MinSumAfterReplaceTwoWithSum

/*
Url: https://www.geeksforgeeks.org/minimize-the-sum-calculated-by-repeatedly-removing-any-two-elements-and-inserting-their-sum-to-the-array/

Minimize the sum calculated by repeatedly removing any two elements and inserting their sum to the Array
Given N elements, you can remove any two elements from the list, note their sum and add the sum to the list. Repeat these steps while there are more than a single element in the list. The task is to minimize the sum of these chosen sum in the end.

Examples:

Input: arr[] = {1, 4, 7, 10}
Output: 39
Choose 1 and 4, Sum = 5, arr[] = {5, 7, 10}
Choose 5 and 7, Sum = 17, arr[] = {12, 10}
Choose 12 and 10, Sum = 39, arr[] = {22}

Input: arr[] = {1, 3, 7, 5, 6}
Output: 48

Approaches:

1. Using recursion

Initialize sum with 0
Keep following below operation till only one element left in the slice

	sort input slice in ascending order.

	calculate sum of first two elements(addResult) from sorted slice
	add addResult to input slice
	add addResult to sum.

*/
