package main

func subtractProductAndSum(n int) int {
	product, sum := n%10, n%10
	for n /= 10; n > 0; n /= 10 {
		product *= n % 10
		sum += n % 10
	}
	return product - sum
}
