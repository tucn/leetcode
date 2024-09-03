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

3. Linked List In-Place Reversal
**When**: Rearrange the links between nodes in a linked list (Reverse a linked list, swap nodes)
**How**:
- Using three pointers (prev, current, next)

4. Monotonic Stack
**What**: Monotonic Stack maintaining elements in either increasing or decreasing order. 
**When**: Using stack to find the next `Greater` or `Smaller` element in an array
**How**:
- Initialize an empty stack.
- Iterate through the elements and for each element:
    while stack is not empty AND top of stack is more than the current element
        - Pop element from the stack
    Push the current element onto the stack.
- At the end of the iteration, the stack will contain the monotonic increasing order of elements.

5. Top `K` element
**When**: Find `K largest/smallest/most frequent` elements in the dataset. Normal shorting take O(nlogn)
**How**: Using `Min heap`
- Using a heap of K elements.
- For each element, compare it with the root of the heap.
- If the element is greater than the root, replace the root with the element.

6. Overlapping Intervals

7. Reverse integer
**When**: Reverse a 32-bit signed integer.
**How**:
- Using mod and divide
```
reversed := 0 
x := input
while x != 0 {
    reversed = reversed * 10 + x%10
    x = x/10
}
```