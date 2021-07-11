Big O - how the runtime scales with respect to some input variables
constant time - O(1): the runtime doesn't change as the input grows. 
linear time - O(n): the runtime increases linearly as the input grows
logarithmic time - O(log(n)): what power do I need to power my base by to get N? The key is "halving". Think binary search. 
Weird ones: O(n * log(n)) - halving but for each of our n items. Merge sort and quick sort. 
quadratic time - (n^2): the runtime increases by n^2
exponential time - o(2^n) backtracing problems or recursive problems. 
factorial time - O(n!) - permutations of things, 

Rules for big O

1. if you have distinct steps in your algorithm, you add those steps together. Eg step a then steb b - O(a + b)
2. Drop constants because we are just looking at how things scale (linear, quadratically, etc). 
3. If you have different inputs, use different variables to represent them. 
4. Drop non-dominant terms. 

https://www.bigocheatsheet.com/

Often you can tradeoff between time and space complexity, (more time for less space, or less time for more space). Generally want to increase space and decrease time because you can always scale up. 

Space complexity: how does space utilization use as the input scales up. NB that runtime stack often is counted so ask the interviewer. 

Logarithms - inverse of taking the exponent of something. Eg 2^3 = 8 ⇔ log base 2 (8) = 3. 

