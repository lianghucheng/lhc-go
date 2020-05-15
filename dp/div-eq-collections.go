package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	var sum = n*(n+1)/2
	if sum % 2 == 1 {
		return
	}

	var x = sum /2
	var a = make([]int, sum + 1)
	a[0] = 1
	for i := 0; i <= n; i++ {
		for j := i * (i + 1) / 2; j >= i; j-- {
			a[j] += a[j - i]
		}
		for j := 1; j <= i * (i + 1) / 2 && j <= x; j++ {
			fmt.Printf("%d  ",a[j])
		}
		fmt.Println()
	}

	fmt.Println(a[x])

	return
}