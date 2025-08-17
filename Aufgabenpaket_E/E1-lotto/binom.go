package main

// gibt den Binomialkoeffizienten n über k zurück.
func binom(n, k int) int {
	if k < 0 || k > n || n < 0 {
		return 0
	}
	if k == 0 || k == n {
		return 1
	}
	return binom(n-1, k-1) + binom(n-1, k)
}
