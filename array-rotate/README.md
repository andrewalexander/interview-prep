# array-rotate

[Original](https://github.com/codingforinterviews/practice-problems/tree/master/array_rotate), as always:
---
This problem is to rotate a given array to the right by `n` steps.

For example:

Given `[1, 2, 3]` and `n = 1`, you should return `[3, 1, 2]`

Each step, the last element in the array is moved to the front of the array, and the rest are shifted right.

Another example:

Given `[1, 2, 3, 4, 5]` and `n = 3`, you should return `[3, 4, 5, 1, 2]`

## Questions for further understanding

1. What is the time complexity of your solution? How about space?
2. Can you do this in-place?

*Challenge:* There is an O(n) time / O(1) extra space solution.
---

Code Output:
```
naive:
hit counter: 4
[1, 2, 3] -> [3, 1, 2]
hit counter: 8
[1, 2, 3, 4, 5] -> [3, 4, 5, 1, 2]
hit counter: 15
[1, 2, 3, 4, 5, 6, 7, 8, 9, 10] -> [6, 7, 8, 9, 10, 1, 2, 3, 4, 5]
-------
better:
hit counter: 2
[1, 2, 3] -> [3, 1, 2]
hit counter: 4
[1, 2, 3, 4, 5] -> [3, 4, 5, 1, 2]
hit counter: 9
[1, 2, 3, 4, 5, 6, 7, 8, 9, 10] -> [6, 7, 8, 9, 10, 2, 3, 4, 5, 1]
```
