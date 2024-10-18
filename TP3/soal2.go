package main

import "fmt"

func segitiga(a, b, c float32) bool {
	return a > 0 && b > 0 && c > 0 && (a+b > c) && (a+c > b) && (b+c > a)
}

func main() {
	var a, b, c float32
	fmt.Scan(&a, &b, &c)

	if segitiga(a, b, c) {
		fmt.Printf("%.1f, %.1f, dan %.1f segitiga? true\n", a, b, c)
	} else {
		fmt.Printf("%.1f, %.1f, dan %.1f segitiga? false\n", a, b, c)
	}
}