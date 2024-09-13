# Introduce 
Here is the flow chart for leet code problem 

# Flow chart 

1. Is it a graph problem?
    - Is it a tree? -> DFS
    - Is it a DAG? -> Topological sort
    - Is it the problem of finding the shortest path in a DAG?
        - If yes, is it weighted? -> Dijkstra
        - If no, then it is a simple DAG problem -> BFS
    - Is the problem involving connectivity ? -> Disjoin Set/Union
        - Is the problem has small constrains ? DFS/Backtracking
        -> BFS
2. Need to solve for kth smallest/largest? -> Heap/Shorting
3. Involves linked list? -> Two pointer
4. Small constrains?
    - Is brute force is fast enough? -> Brute force/Backtracking
      - if no, is it a DP problem? -> DP
5. Is it a subarray/substring problem?
    - Is it dealing with sum or additive ? -> Prefix Sum (Yes) /Sliding Window (No)
6. Calculating of max/min value? 
    - Monotonic condition? -> Binary Search
    - Can be split into subproblem? -> DP
    - Greedily calculate? -> Greedy
7. Asking for number of ways?
    - Is brute force is fast enough? -> Brute force/Backtracking
    - If no, is it a DP problem? -> DP
8. Muliple sequence/array? 
    - Monotonic condition? -> Two pointer
    - Can be split into subproblem? -> DP
9. Find or enumerate indices ? 
    - Monotonic condition? -> Two pointer
10. O(1) space? 
    - Monotonic condition? -> Two pointer