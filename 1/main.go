package main

import "fmt"

func main() {

	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)
	fmt.Println(permutation(a, c), combination(a, c))
	fmt.Println(permutation(b, d), combination(b, d))
}

func factorial(n int) int {

	var factorial int = 1
	for i := n; i >= 1; i-- {
		factorial *= i
	}
	return factorial
}

func permutation(n, r int) int {

	var permutate int
	nfactorial := factorial(n)
	nrfactorial := factorial(n - r)

	permutate = nfactorial / nrfactorial

	return permutate
}

func combination(n, r int) int {

	var combinate int
	nfactorial := factorial(n)
	rnrfactorial := factorial(r) * factorial(n-r)

	combinate = nfactorial / rnrfactorial

	return combinate
}
