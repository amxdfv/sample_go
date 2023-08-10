package main

import "fmt"

type oslic struct {
	number_one, number_two int
}

func CountOslic(one int, two int) int {
	fmt.Println("Ослиное число один")
	fmt.Scanf("%d\n", &one)
	fmt.Println("Ослиное число два")
	fmt.Scanf("%d\n", &two)

	var sample = oslic{one, two}
	var result = sample.number_two + sample.number_one
	return result

}
