# Implementation of the Collatz Conjecture

The [Collatz conjecture](https://en.wikipedia.org/wiki/Collatz_conjecture) is a very interesting conjecture in mathematics that concerns a sequence defined as follows: start with any positive integer n. Then each term is obtained from the previous term as follows:
    - If the previous term is even, the next term is one half of the previous term.
    - If the previous term is odd, the next term is 3 times the previous term plus 1.

The conjecture is that no matter what value of n, the sequence will always reach 1.


## Implementation Tasks

This implementation for the Collatz Conjecture is broken down into 4 parts as part of the assignment. These are:

- **Part(a)**: Implement a program to calculate the number of steps to reach 1 the first time for any given m ∈ N.
- **Part(b)**: Adapt the implementation to calculate the number of steps for each m ∈ [1, 10000].
- **Part(c)**: Try to improve the implementation to calculate the steps as fast as possbile.
- **Part(d)**: Print the number of steps for each m ∈ [1, 10000] as diagram.

### Solution: Part(a)

*Implement a program to calculate the number of steps to reach 1 the first time for any given m ∈ N:*

The solution for part(a) is available in `part_a.go`. I've implemented the sequence formula to compute the number of steps for m to reach 1 the first time using a loop.

How to run part(a) solution locally:

The program asks for an integer input m from stdin and it returns the number of steps to reach 1 in the stdout.

```bash

>> go run part_a.go

Enter value for m:

10

It takes 6 steps for the number 10 to reach the number 1.

```

### Solution: Part(b)

*Adapt the implementation to calculate the number of steps for each m ∈ [1, 10000].*

The solution uses an algorithm to compute the steps for m to reach 1 for all m in the range of `startValue` and `endValue`. The number m and its corresponding steps are stored in a dict, where key is number m and steps to compute the number is the value.

How to run part(b) solution locally:

The program asks for an integer input ...

```bash

>> go run part_b.go

Enter start value:
1

Enter end value:
10

All steps for each number within a range of numbers from 1 to 10

map[1:0 2:1 3:7 4:2 5:5 6:8 7:16 8:3 9:19 10:6]

```

### Solution: Part(c)

*Try to improve the implementation to calculate the steps as fast as possbile.*

The solution uses an optimized algorithm to compute the steps for m to reach 1 for all m in the range of `startValue` and `endValue`. The number m and its corresponding steps are stored in a dict, where key is number m and steps to compute the number is the value.

The optimization techniques are:

1. Stores the steps for each m in the cache, so that the pre-computed value can be reused next time it comes across the number. This will speed up performance.

2. Any number which is a power of 2 will always be even when divided by 2 and will always reach 1. By taking the log 2 of these numbers, the performance will speed up. The log gives the number of divisions required to reach 1, which is same as number of steps required to reach 1.

3. Since an odd number when multiplied by 3 and added by 1 results to a even number, it can further be divided by 2 in the same step hence skipping one loop iteration but still maintaining the steps.

How to run part(c) solution locally:

The program asks for an integer input ...

```bash

>> go run part_c.go

Enter start value:
1

Enter end value:
10

All steps for each number within a range of numbers from 1 to 10

map[1:0 2:1 3:7 4:2 5:5 6:8 7:16 8:3 9:19 10:6]

```

**With optimization (Part c):**

    - Took 0.003035394 seconds for range [1, 10000]
    - Took 0.003045096 seconds for range [1, 10000]
    - Took 0.003011613 seconds for range [1, 10000]

**Without optimization (Part b):**

    - Took 0.003786034 seconds for range [1, 10000]
    - Took 0.002952388 seconds for range [1, 10000]
    - Took 0.002959046 seconds for range [1, 10000]

### Solution: Part(d)

*Print the number of steps for each m ∈ [1, 10000] as diagram.*

In this solution [glot](github.com/Arafatk/glot) library is used to generate a scattered graph diagram which is a function of (m, steps) for m b/w `startValue` and `endValue` and m are the positive integers in this range and steps are the corresponding steps for m to reach 1. The resulting graph is saved in the "Collatz.png" file in the same directory.

How to run part(d) solution locally:

The program asks for an integer input ...

```bash

>> go run part_d.go

Enter start value:
1

Enter end value:
10

Open Collatz.png to view plotted diagram.

```

Range 1 to 10

![Range 1 to 10](https://github.com/StephenDsouza90/go-collatz/blob/main/images/Collatz_1_to_10.png)

Range 1 to 10000

![Range 1 to 10000](https://github.com/StephenDsouza90/go-collatz/blob/main/images/Collatz_1_to_10000.png)


## Install dependencies

The dependencies for this application are:

```bash

>> go get github.com/Arafatk/glot

```

Additional dependencies can be found on [glot](github.com/Arafatk/glot).


## Assumptions

The following assumptions are made when implementing the above tasks:

1. The number of steps does not refer to how many loops are required for m to reach 1, but the numbers in the sequence from m to 1.
2. The optimizations in part(c) are of two fold, the first is an optimization for any single value of m to reach 1 and the second optimization is for all m in a range to reach 1 by using pre-computed steps of already known m.
3. The diagram in part(d) is a plotted graph of x=m and y=steps.
4. Inputs are inputted from stdin and outputs are outputted to stdout.
5. Each task is completely isolated from other tasks, ie. there is no python imports of the functions to be re-used in other parts.