# array-rotate Solution

This one was another fun one for me... I spent more time than I care to admit trying to get the "challenge" part out the gate. Ended up needing to salvage my pride and just implement the "naive" approach first with two arrays. I think I had a slightly fun take on it, but it basically ended up being  O(n+) runtime and was  O(n+) for space. Once I wrote the naive version, though, things just "clicked" and I saw what needed to be done to get the challenging version. _Really_ got stuck on this one for a bit and was starting to get worried... 

Either way, this one was one of the more rewarding ones I've done recently as it appears simple on the surface, but has many components that will trip people up. I love this problem as an interview problem for so many reasons. 

## Conceptual Solution
This was a deceptively simple problem, in my opinion. 

> "All problems in computer science are data transformation problems". 

Not sure who first said that, but Bill Kennedy of ArdanLabs shares that in his Golang courses and it is the most true statement I have ever heard about software engineering/computer science. This statement felt especially relevant with this problem because I felt like I really had to play with the data in order for the algorithm to jump out at me. I started mapping out character by character what I wanted to happen, and then I saw what I was getting stuck at with my fixed offset. Initially, I tried to jump right in and play with the length and do offsets based on n and the whole nine yards. I knew I needed to do that eventually, but I needed to start somewhere simpler first. 

This is a wonderful interview problem because it will tease that "shoot first, ask questions later" out of the candidate right away - myself included. I got blinded by the fact that I knew I was going to have to use _a tool_ and rather than being pragmatic about the use of that tool (offsets and clever indexing/swapping), I just started abusing it. I needed to take a step back and really look at what I was doing. Then it all made sense. 

So enough of my yapping. What is it we are doing here? For the naive version of my function, I used a second array to copy values to. I realized that the hard part of this problem is handling the wrap-around values (why I was so stuck...), so I began by putting those in a third array in order. Then I loop over the original input array, starting with the first `n` values that wrapped around (`n` being the number of spaces to rotate; also non-coincidentally the number of elements that wrap around), and then finishing up with the "front" of the original array.

```
input array = [1, 2, 3, 4, 5] (len == k)
n = 3
wrap array = [3, 4, 5] (len == n)
remaining from input = [1, 2, X, X, X]
new array = [wrap array] + remaining from input 
new array = [3, 4, 5, 1, 2] (len == k)
```

This seemed fun/cool to me, as I knew the key insight was around handling those wraparound characters. I knew it wasn't the worst we could do, but still - it is O(n+k) runtime and O(2n+k) space; the challenge said there's a true O(n)/O(1) way to do this, and I am stubborn!

So I looked at what the beginning/final versions looked like and tried to just blindly apply an "algorithm" to it visually. 

```
n = 1:
[1, 2, 3] -> [3, 2, 1]
n = 3:
[1, 2, 3, 4, 5] -> [3, 4, 5, 1, 2]
n = 5:
[1, 2, 3, 4, 5, 6, 7, 8, 9, 10] -> [6, 7, 8, 9, 10, 1, 2, 3, 4, 5]
```

This is where I realized the magic. Looking at the `n=1` case doesn't really show it, but the `n=3` and `n=5` cases definitely do. 

When `n=3`, the value originally in the 3rd index ends up in the first index. (i = 2; i = 0, respectively for zero-indexed arrays). The value originally in the 4th index ends up at the second index, then 5th to the 3rd. Now we are passed the "wrapped" values. From here, the value originally in the 1st index ends up in the 4th and the value originally in the second index ends up in the fifth. I got tripped up because these aren't direct swaps the entire way through the loop. What I mean by that is that if we just swapped the first index with the third index, we would have `1` where `5` should be in the final array. But if we expect a constant space solution, we must do the swaps in place (can't have two temp variables - only one). So when we swap the 3rd index with the first, we must put the original value from the first index into the third, at least for now. 

The key insight is to realize that the value now in the third element is not in its final resting place. As we move our sliding window to the right, we will sort/clear things up. So we start at the left, swapping the element i with i + offset (which will be more clearly defined down in the Algorithm section below) until the right side of our window hits the end of the array. Now we know that we have swapped all the offset bits. Now we simply swap adjacent values to clean up our mess from swapping in place...

## Alogorithm
Our goal is to loop through the array once, so we know we want to reference things relative to i, where i starts at 0 and increments through the length of the array (`k` below).

```
k=5 (len(arr))
n=3
offset = 2 (k - n)
---
i=0: [1, 2, 3, 4, 5]
      ^-----^ swap
i=1: [3, 2, 1, 4, 5]
         ^-----^ swap
i=2: [3, 4, 1, 2, 5]
            ^-----^ swap
i=3: [3, 4, 5, 2, 1]
               ^--^ swap
```

So when i+offset was less than the length of the array, we swapped element i with the element i+offset. When i+offset was greater than the array, we simply swap element i with element i+1. Since we will be addressing i+1, we know our loop bounds need to stop at k-1. I will call the i+offset value the "swapping offset window". When i=3, that right side of that window would reference an element in our array that doesn't exist; at this point we know we just need to swap element 4 and 5 (i=3,i=4, respectively). Formally this means that if i+offset > k, we swap arr[i] with arr[i+1].

Generalized pseudocode of above:
```
k = len(arr)
offset = k-n
for i=0; i<k-1; i++:
	if i+offset < k:
		tmp = arr[i+offset]
		arr[i+offset] = arr[i]
		arr[i] = tmp
	else:
		tmp = arr[i+1]
		arr[i+1] = arr[i]
		arr[i] = tmp
```

## Code Breakdown

Let's start with the more naive approach:

```
def rotate_naive(arr, n):
    new = []
    wrap = []
    counter = 0
    # n digits wrap around every time, so lets grab those from the end
    for i in range(len(arr)-n, len(arr)):
        counter = counter + 1
        wrap.append(arr[i])
    for i in range(len(arr)):
        counter = counter + 1
        if i < n:
            new.append(wrap[i])
        else:
            new.append(arr[i-n])
    print 'hit counter: '+ str(counter)
    return new
```
We have a counter to show the runtime efficiency in terms of loop counts. What this code is doing is first grabbing the elements that we know will wrap around (same number of elements as our `offset` above - the last `n` elements where `n` is the number of elements to rotate). Then we iterate through our input array and assign values to our `new` output array based on whether the value should have come from wrapping around or sliding down. Functionally, this obviously still accomplishes what we want, but does so inefficiently.

We are creating an entirely new array of size `k` plus our `wrap` array of size `n`. 

The more elegant approach:
```

def rotate(arr, n):
    # swap one pair at a time, separated by our wrap offset. When our right
    # bound reaches the end of the array, we just swap adjacent pairs since we
    # know things have been sorted up until there
    offset = len(arr)-n
    counter = 0
    for i in range(len(arr)-1):
        counter = counter + 1
        if i + offset < len(arr):
            tmp = arr[i+offset]
            arr[i+offset] = arr[i]
            arr[i] = tmp
        else:
            tmp = arr[i+1]
            arr[i+1] = arr[i]
            arr[i] = tmp
    print 'hit counter: '+ str(counter)
    return arr
```

This does exactly what the pseudocode is doing without additional overhead. This is what I was trying to get out of the gate, but I needed to see _an_ implementation to see how to optimize its inefficiencies away. In both examples, we are recognizing that the "key" is whether a value had to "wrap" around the array or not and placing values according to that offset. 

The loop goes for 1 less than the size of the input array. The first condition checks to make sure the right side of the "swapping offset window" doesn't extend beyond our array. Once we reach the end, we know all "wrapped" values are where they belong (everything to the left of our sliding window), and the remaining values from the left of our window to the end of the array were there as a result of handling those swaps, and just need to be exchanged one at a time with each other until the original order is restored.

## Efficiency
This was an educational moment for me to properly define the space/runtime considerations of algorithms. The "elegant" solution seemed to be O(n+1) for space since I was counting the input array in the calculations. Once some homework was done, it made logical sense that the algorithms are measured in terms of their added costs. So if you are receiving an input, you can assume that that cost was calculated/amortized elsewhere and should not factor into your algorithms' efficiency analysis. The reality of how the function is invoked (stack frames generated, memory allocated on heap if necessary, etc.) is important, but again, we can assume that all that is fixed for any kind of algorithm you'd implement, so should not factor in to your calculations. My way of looking at is along the lines of realizing that you have your data to manipulate no matter what. When shopping for algorithms, you don't need them to repeat the cost of your data back to you, you just need to know how much _more_ the algorithm is going to cost you.

Once again though - _this is all my interpretation_. I am an academic at heart and I do do my research on these things, but I tend to make assumptions/generalizations and have been known to jump to conclusions prematurely. I sometimes think I know or understand something when in fact I missed the ball entirely. So again (and I'll probably say this in every one of these), please keep me honest!

### Runtime: O(n+k)/O(n)

In the "naive" version, I loop through the array one full time plus a few extra bits because I fetch the digits that will be wrapping around. I call the "extra bits" `k`, while `n` is the number of elements in the input array.

In the more elegant solution, we touch each element exactly once - `n`.

### Space: O(n+k)/O(1)

Naive version allocated an entirely new array equal in size to the input (`n`) plus the array for the "wrap" digits, so `n` + `k`.

Elegant version only ever needed to store one value in memory at a time per loop, making it O(1).


