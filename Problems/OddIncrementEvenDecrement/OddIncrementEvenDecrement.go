package OddIncrementEvenDecrement

/*
Url: https://www.geeksforgeeks.org/increment-odd-positioned-elements-by-1-and-decrement-even-positioned-elements-by-1-in-an-array/

Increment odd positioned elements by 1 and decrement even positioned elements by 1 in an Array
Given an array arr[], the task is increment all the odd positioned elements by 1 and decrement all the even positioned elements by 1.

Examples:

Input: arr[] = {3, 6, 8}
Output: 4 5 9

Input: arr[] = {9, 7, 3}
Output: 10 6 4

Approaches:

1.
Initialize x as 1.
Traverse through input slice do following:
	add x to element
	change polarity of x

2. Reduced input traversal by half.

*/

// GenerateSlice return new slice by incrementing odd elements by one and decrementing event elements by one from input slice.
func GenerateSlice(input []int) []int {

	N := len(input)

	if N < 1 {
		return []int{}
	}

	if N == 1 {
		return []int{(input[0] + 1)}
	}

	if (N % 2) != 0 {
		return generateSliceForOddSize(input, N)
	}

	return generateSliceForEvenSize(input, N)
}

func generateSliceForOddSize(input []int, N int) []int {

	x := 1
	xd := 1
	limit := N / 2

	newlist := input[0:N]

	for i := 0; i < limit; i++ {
		newlist[i] = input[i] + x
		x = x * -1

		newlist[N-i-1] = input[N-i-1] + xd
		xd = xd * -1
	}

	newlist[limit] = input[limit] + x

	return newlist
}

func generateSliceForEvenSize(input []int, N int) []int {

	x := 1
	xd := -1
	limit := N / 2

	newlist := input[0:N]

	for i := 0; i < limit; i++ {
		newlist[i] = input[i] + x
		x = x * -1

		newlist[N-i-1] = input[N-i-1] + xd
		xd = xd * -1
	}

	return newlist
}
