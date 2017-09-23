# linked-list-reverse solution

There are many solutions to this problem. I naively tried to just jump in and code up a few solutions before really drawing this one out and planning the algorithm, and it reminded me of the value of being pragmatic sometimes ;).

After the third time of getting that "this is wayyyy too complicated" feeling, I decided to take a step back... And when I did, I quickly saw a key to greatly simplifying this problem. It may have been obvious/intuitive for some, but hey, learning is learning! 

I also did not have a formal education in computer science, so if this was the way it was taught to you formally, that sounds like a great way to learn it! The more I reflect on it, the more obvious "reversing" was, but I digress...


## Conceptual Solution
I started by writing this on a sheet of paper, as we usually do with interview problems:
```
in: A -> B -> C -> null
out nil -> C -> B -> A 
```

I think that that's how most would start it. After I ran through a naive attempt at doing this, I realized I was copying things all over and trying to maintain state of order and sorting and a whole mess of things that didn't matter. I needed to think a little smarter. That's when I realized that all I'm doing is reversing the direction of the arrows! Yes... it's reversing a string... of course I should have known that. [One of the 10,000](https://xkcd.com/1053/).

So I redrew the problem...
```
in:  A -> B -> C -> null
out: A <- B <- C <- null
```

And for some reason, this made the algorithm instantly pop out to me. Usually with interview questions, you can get started with an input and then just go to town manipulating it and hacking your way to the final answer. This was one that felt 1000x easier by coming up with a simple step by step algorithm and constantly pointing back to that when writing the code to keep yourself honest. There's no shame in admitting the problem is too complex to hold entirely in your head ;)

I knew the tricky part was going to be around "flipping the arrow" as easily as possible. I am pretty sure I still messed that up because I am doing pre- and post-loop operations on the output list when I'm 51% sure there's a better way...

Either way, the tl;dr is: 
For each input node we see, point it to the last node we just saw from the incoming list.

## Algorithm
Let input node be `in`

1. Create node/element in our returned list (`out`)
2. Set `out.data` to `in.data`, increment `in` pointer to next node
3. Create a placeholder node `tmp`
4. Set `tmp.next` to `out` (we just reversed the arrow!)
5. Set `out` to `tmp` 

Repeat for all elements in `in`. My interpretation of the problem was that the first node in our reversed list should be a non-nil node but have a "nil" value. With that being the case, we need to add one last node to `m` and then return that. In my code, I explicitly used the string "nil" instead of the default Golang [zero-value](https://tour.golang.org/basics/12) of empty string (`""`) for clarity. (Prints `A -> B -> C -> nil` instead of `A -> B -> C -> `.)


## Code breakdown
So let's look at how the code is implementing the above algorithm.

```
func reverseList(n Node) Node {
	// set up our root node
	m := &Node{
		data: n.data,
		next: nil,
	}
	n = *n.next
	for n.next != nil {
		// reverse the arrow!
		tmp := &Node{
			data: n.data,
			next: m,
		}
		m = tmp
		n = *n.next
	}
	tmp := &Node{
		next: m,
		data: n.data,
	}
	return *tmp
}
```
Let's run through this one loop at a time to really understand what's happening. To make it easier, after every "loop", I am going to spit out how much of `in` we have seen (read: copied; the pointer will always be one node ahead of what's printed), and what the current `out` looks like.


Example:
```
in:  A -> B
out: B -> A
```
This means that we have copied `A` and `B` over to our `out` variable, but the `in` pointer could already be at `C`.

#### Loop 1/(n=1)
We start by initializing our `out` list with its first node. We copy the "A" value from `in.data` to `out.data`.  Then we increment the `in` pointer to point to the `B` node.

For clarity, our loop is set up to run until our `in` list is at its last element (the one with its `data` value set to `"nil"`). Following along with our rewritten problem, this fails that evaluation when `n.data` is pointing to `null`. 

> This is imporant because it will fail this evaluation check before it copies that final string value of `nil` to our `out` list!

Now we enter the loop. We make our `tmp` node, allowing us to reverse the `next` pointers as we had planned.

We now copy `B` from `in.data` to `tmp.data`, and set `tmp`'s `next` pointer to point to `out`. Once we set `out` to be `tmp` (setting the arrow in the proper direction), `out` looks like:

```
in:  A -> B
out: B -> A
```

#### Loop 2/(n=(1,N))

At the end of the last loop, we incremented the `in` pointer to point to `C`. This means that the loop checks `in.next`; it is non-nil (`next` is the node with `.data == "nil"`). Just like above, we play a quick game of ping-pong to reverse the `next` pointer:

- `tmp.data` gets set to `C`
- `tmp.next` gets set to `out` 
- `out` gets set to `tmp`

So now we have:
```
in:  A -> B -> C
out: C -> B -> A
```

#### Loop 3/(n=N)

Getting repetitive now!

`in` got incremented to point to `.data == "nil"`, and so the `n.next != nil` check failed (`n.next == nil` now!). So as we warned before, we need to do one last ping pong operation after the loop to get all values from the `in` list.

- `tmp.data` gets set to `nil`
- `tmp.next` gets set to `out` 
- `out` gets set to `tmp`

After our final ping/pong operation:
```
in:  A -> B -> C -> nil
out: nil -> C -> B -> A
```

#### All the loops/Generalized algorithm
In my code, I did a few more than just a list of A-C to prove the algorithm was effective. 

If `N` equals the number of letters in the list and `n` is the count of loop iterations, then `n=1` would always behave like `Loop 1` above, `n=(1,N)` would behave like `Loop 2`, and `n=N` would behave like `Loop 3`.

### Efficiency

#### Runtime: O(n)
You only make one pass of the input list and build the output list concurrently while traversing the input list.
#### Space: O(2n) ~= O(n)
We will always be copying from one memory location to another, and thus will always at least need `2 * m` space in each iteration, where `m` is the total memory required to hold one Node object. All function-related memory overhead (stack pointers and the like) will be considered constant across iterations. 

In reality, garbage collection likely won't happen between iterations, so the likelihood that both the `in` and `out` lists exist in their entirety simultaneously is probably very far away from zero, and so that should be considered the "real case".
