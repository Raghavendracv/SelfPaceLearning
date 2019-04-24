package MinOperationsToMoveUpperThanLower

/*
Url: https://www.geeksforgeeks.org/minimum-number-of-operations-to-move-all-uppercase-characters-before-all-lower-case-characters/

Given a string str, containing upper case and lower case characters. In a single operations, any lowercase character can be converted to an uppercase character and vice versa. The task is to print the minimum number of such operations required so that the resultant string consists of zero or more upper case characters followed by zero or more lower case characters.

Examples:

Input: str = “geEks”
Output: 1
Either the first 2 characters can be converted to uppercase characters i.e. “GEEks” with 2 operations.
Or the third character can be converted to a lowercase character i.e. “geeks” with a single operation.



Input: str = “geek”
Output: 0
The string is already in the specified format.

Approaches:

1. Traverse input and calculate last occurrence of Upper case character (lastUpper)
& first occurrence of lower case character (firstLower) and return minimum operation count based on below logic:

	if lastUpper < firstLower
		no operation required
	else minmum of below
		i. lastUpper
		ii. length of input - firstLower
*/
