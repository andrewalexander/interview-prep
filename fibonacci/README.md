# fibonacci

Taken from [this problem]() (the `lite.py`) and [this problem]() (`dynamic.py`). 

From the problem directly:
---
For this question, you will write a program that generates values from the
Fibonacci sequence. The Fibonnaci sequence is recursively defined by:

> *F<sub>n</sub> = F<sub>n - 1</sub> + F<sub>n - 2</sub>*

Using the following seed values:

> *F<sub>0</sub> = 0, F<sub>1</sub> = 1*

Given a number *n*, print the *n*th value of the Fibonacci sequence.

## Examples
Input:

```
12
```

Output:

```
144
```

Input:

```
30
```

Output:

```
832040
```
## Input Format and Restrictions
Each test case will consist of a single positive integer *n*.

The inputs will always satisfy the following restrictions:

* *F<sub>n</sub>* < 2^32 - 1,
* 0 <= *n* < 50

## Solution Output:

```
$ python lite.py 
input 12: 144
input 30: 832040

$ python dynamic.py 
input 30: 832040
input 12: 144
input 27: 196418
```

Uncomment the `print` statements to see all the calls to `fib` in `lite.py` and how few "cache misses" in `dynamic.py`.

