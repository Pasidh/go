package main

import "fmt"

func fibonacci() func() int {
	sum := 0
	pre := 0
	return func() int {
		if sum == 0 {
			fmt.Println(sum)
			sum = sum + 1
		} else {
			sum = sum + pre
			pre = sum - pre
		}
		return sum

	}
}
func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
