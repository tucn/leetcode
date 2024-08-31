# leetcode
Leet code problem in Go

# Noted
- Single number in Array problem: Using XOR through all arrays


# CheatSheet Pattern
1. Prefix sum
**When**: Query sum of element in the subarrays
Simple solution will take O(n*m) (n is the length of subarrays and m is the length of the arrays)
**How**:
- Turn the original m arrays into M with each M[i] = M[0] + M[1] + .. M[i]
- sum[i+j] = M[i] - M[j]
Hint: if we don't need the input arrays, just use original array as M

2. Two pointer
**When**: 
- Can be used as sliding windows to find the longest subarray or the arrays with maximum sum
- Palindrome

2.1. Fast and slow pointer (rabbit and turtle problems)
**When**:
- Problem related to linkedlist or arrays involving finding cycle (loop). Example: [1,`2,3,4,5,6,7`,2,3,4,5,6,7,2,...]
- Finding middle node of a given linked list
**How**
- Moving two pointer at different speed