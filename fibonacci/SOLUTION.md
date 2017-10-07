# fibonacci Solution

Almost any introductory programming course/video series/educational material out there will probably have some form of the fibonacci sequence. The definition was in the `README`, but can be summed up as a list of numbers such that any given number is the result of the previous two values in the list. Ideally, we want to build/store this in a lookup table so we can avoid computing things multiple times. This is the only difference between the `lite` and `dynamic` problems. 

## Conceptual Solution
The equation summed it up pretty well:
> *F<sub>n</sub> = F<sub>n - 1</sub> + F<sub>n - 2</sub>*

Using the following seed values:

> *F<sub>0</sub> = 0, F<sub>1</sub> = 1*

They mention the seed values because we need to handle the cases for 0 and 1. There is technically a different implementation for negative values, so by specifying these seeds and providing the equation, we have no ambiguity on which algorithm to implement. 

As far as our code is concerned, we now know that we want to explicitly handle these two cases. Since we know that we are dealing with the set of all positive integers, `n-2` is undefined for both `n=0` and `n=1`; and `n-1` is also undefined for `n=0`.

## Algorithm

See equation above.

## Code Breakdown
lite.py:
```
def fib(n):
    # print 'n: '+str(n)
    if n == 0:
        return 0
    if n == 1:
        return 1
    return fib(n-1) + fib(n-2)
```

Not really sure how to get any simpler than that... We have our two base cases, and an optional print statement showing just how many times we call `fib()` in this naive implementation. 

dynamic.py:
```
seen = {0:0, 1:1}

def fib(n):
    if n == 0:
        return 0
    if n == 1:
        return 1
    if not seen.get(n):
        # print str(n)+': cache miss...'
        seen[n] = fib(n-1) + fib(n-2)

    return seen.get(n)
```
This one is almost identical, except we add a hashmap (dict in python) to keep track of values we've seen. Now every time `fib()` is called, we look to see if we've computed that value before. If we have, we get to skip calling `fib()` and simply return. This ensures that we only ever call `fib()` for a value of `n` that we have never seen. 

The `main()` function demonstrates this clearly when you un-comment the `print` statement in the `fib` function as it prints `cache miss...` for all values during the call to `fib(n=30)`, but immediately returns the values for `fib(n=12)` and `fib(n=27)`.



## Efficiency

### Runtime: It's complicated?

I needed to cheat a bit on this one... There's a lot of stuff out there on it, but I basically started by thinking about how many times `n` is called per iteration. The answer changes slightly based on implementation, though so I will walk through my thought process. Admittedly, this is something I struggle with a lot, so if this is totally off, please edit/Pull Request/File an issue/email-me/whatever. Just let me know!

#### lite:
- For 0 and 1, fib() was called once: `fib(n=0)` and `fib(n=1)`
- For 2, fib() was called 3 times: `fib(n=2) -> fib(n=1) + fib(n=0)`
- For 3, fib() is called 5 times: `fib(n=3) -> fib(n=2) + fib(n=1)`; `fib(n=2)` counts as 3 as we determined from the last line.
- For 4, fib() is called 9 times: `fib(n=4) -> fib(n=3) + fib(n=2)`; `fib(n=3)` == 5 + `fib(n=2)` == 3

We quickly see that this resembles a 2n-1 relationship... So I think we just call that O(n) or O(2n)?

```
n=0: O(1)
n>=1: O(2n-1)
```

#### dynamic:
So this implementation depends on the order you call things in. If you call `fib(5)` and then `fib(15)` and then `fib(30)`, you don't get much benefit of the memoization that goes on. In the `lite` version, it didn't matter how you called it because it was going to call fib `2n-1` times for all values of `n>=1`


So let's pretend we call it in the same order as `lite` above:

- For 0 and 1, fib() was still called once: `fib(n=0)` and `fib(n=1)`
- For 2, fib() was still called 3 times: `fib(n=2) -> fib(n=1) + fib(n=0)`
- For 3, fib() is now called 3 times: `fib(n=3) -> fib(n=2) + fib(n=1)`; `fib(n=2)` has its value stored in our memo/map
- For 4, fib() is called 3 times: `fib(n=4) -> fib(n=3) + fib(n=2)`; `fib(n=3)` and `fib(n=2)` both are stored in memo/map

So it appears that we always get 3 calls to fib(), right? Well not really... In our case, we were lucky that fib(n-1) and fib(n-2) were already calculated for each increasing value of n. But what if we skipped to n=45 after n=4? We would have to calculate n=5 to n=45. I don't know how to model that in terms of an equation.

If instead we called it in a different order (like the test parameters), it calls `fib(30)`, which would calculate `fib(29)`, `fib(28)`, and so on just like the lite version, but storing things in the memo as they go. So when it gets all the way down to n=0 and n=1, it starts the cascade on back up. Now as we step back up the call stack, we don't need to recalculate from 1->n every time. This still results in `2n-1`

If `fib(12)` is called right after `fib(30)`, it returns in O(1) time. As does `fib(27)`


So best case (we've seen the value before), we get O(1). Worst case is still 2n-1. So the average would most likely be 1.5n? Joking... but again, I'm super bad at this stuff!

### Space: O(n)

Because my implementation requires each recursive call to return before returning itself, each stack frame that is created as the program recurses must be saved. Thus, we have O(n) space complexity in the worst case. We could optimize the recursive function to be O(1) in terms of space complexity, but with the memoized version to save runtime complexity, we will still be O(n) as we need to store those elements somewhere. 

