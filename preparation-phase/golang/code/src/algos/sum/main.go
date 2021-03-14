// Say I wanted to calculate the sum of the first n numbers, and I’m wondering how long this will take. Firstly, can you think of a simple algorithm to do the calculation? It should be a function that has n as a parameter, and returns the sum of the first n numbers. You don’t need to do anything fancy, but please do take the time to write out an algorithm and think about how long it will take to run on small or large inputs.

package sum

func sum(n int) int {
	tot := 0
	count := 0
	for ; count <= n; count++ {
		tot += count
	}
	return tot
}
