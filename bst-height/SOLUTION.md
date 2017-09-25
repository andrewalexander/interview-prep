# bst-height Solution

While the problem only deals with finding the height/depth of the node, I found it more educational to learn how to build the insert() function to build the BST up to really get the hang of BSTs.

The problem statement helps clear up some ambiguity for us:

> given a BST with a single node, your functino would return 0

So we know to return 0 when we don't have a left or right node. I liked calling these nodes the leaf nodes, but technically speaking (according to my interpretation of Wikipedia), it looks like the leaf nodes are the nil nodes that are technically on the left and right side of what I had been referring to as the leaf. 

## Conceptual Solution
As stated above, we know that a node with no left and right nodes has a depth of zero, and we return that value accordingly. For any other values, we basically just ask each of our left and right nodes how deep they are, and they keep asking the nodes below them until eventually we get to the edges of each and every node and finally return that 0 value. For each node, we only care about the maximum depth, so when asking our left and right nodes how deep they are (insert joke about how this is a rather personal question to be asking...), we return the max of the two.

This is my terrible way of personifying recursion, but it helps me to think about it as one node saying "I don't know, let me ask these guys first".

The implication is that you traverse to any/all nodes at least once as every node asks all children nodes what their depth is. Like any node-based problem, it helps to break things down one node at a time.

Our goal is to find the edges/leafs of the tree (remember that what I call a leaf technically isn't...). When we do, we return a 0. If we have a left node, we add 1 to it plus the result of a recursive call to `findHeight(n.left)`. If we have a right node, we do the same. 

## Algorithm
Let `n` be the root node input to our recursive algorithm.

1. Check current node's left and right children; if both are nil, we are a leaf. Depth is 0
2. If left node is not nil, add 1 to our left count and make a recursive call to this algorithm to determine that node's size
3. If right node is not nil, add 1 to our right count and make a recursive call to this algorithm to determine that node's size
4. Take the max of left count and right count. Return this value

## Code Breakdown
This is one example where the expressiveness of Golang really comes through, at least in my mind... Seems very readable to me, but I have been writing a lot of Go lately, so I may be biased.
```
func findHeight(n *BSTNode) int {
	if n.left == nil && n.right == nil 
		return 0
	}
	lh := 0
	rh := 0
	if n.left != nil {
		lh = 1 + findHeight(n.left)
	}
	if n.right != nil {
		rh = 1 + findHeight(n.right)
	}
	max := lh
	if rh > lh {
		max = rh
	}
	return max
}
```

With every call, our goal is clear - find the leaf node and return 0 as fast as possible. For each node that has a left or a right, we add 1 to our count and recurse a level for each. We make sure to go to both left and right if they both exist. Then we return the max of our two counts. 

## Efficiency
### Runtime: O(n)
We traverse to every single node exactly once, so that's O(n).

### Space: O(n)
I think? We only ever have `n` nodes to store in memory, and we always traverse it one node at a time. I think that means it's O(n), but I'm not the best at efficiency analysis.
