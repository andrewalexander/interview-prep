# bst-height

Binary Search Tree height, also known as depth. This is distinct from the number of nodes in a BST. If you have a simple 3-node BST, its depth is likely 1. 

Original Question:
```
Given a binary search tree, return its height—that is, the maximum depth reached by the tree.

Example: given a BST with a single node, your function would return 0.

Given a linear BST with only right side nodes 0 -> 1 -> 2 -> (null), where 2 is the tail, your function would return a max height of 2.

Hint: BSTs are a recursively defined data structure.

Hint #2: which tree traversal method covered in the traversal lecture might come in handy here?
```
I'm not sure what lectures it is referring to.

This one was fun as it required me to implement an insert() function for a BST to more easily generate a simple BST. It also forced me to learn the semantic definitions what for implementation decisions and such. Hooray for [Wikipedia](https://en.wikipedia.org/wiki/Binary_search_tree). 

I took the defintion that a node goes on the left if it is less than or equal to the root node and goes on the right only if it is greater (not equal) to the root node's value. Other than this, no re-ordering is possible. Insertions only append to leaf nodes.

This is where I could see interviews going many different directions, and this is why I draw attention to the fact that I only append to leaf nodes/nil nodes. Depending on where you look, there are lots of different methods of sorting Binary Search Trees (or re-ordering). I just made the most naive `insert()` function possible. It only looks for a leaf/nil node and puts values there. It doesn't re-order to ensure any sort of numerical order on the branches. This can lead to some really ineffecient traversals (or just ugly trees), making the order of insertion tied very closely to the structure of the tree.

## Solution Output
```
$ ./bst-height
max height: 4
```

The code inserts the list `8, 6, 25, 14, 3, 7, 15, 24, 9` into a BST and then calls our findHeight function. This tree becomes:

```
      8		---
     / \	 1
    6   25	--- 
   / \   /	 2
  3   7 14	---
       /  \	 3
      9    15	---
            \  	 4
            24  ---
```

So, the output of 4 is correct given this BST.
