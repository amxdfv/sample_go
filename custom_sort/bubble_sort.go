package custom_sort

import (
	"fmt"
	"oslic_v1/samples"
	"time"
)

func Custom_Sort_v1(income_array []int) {
	samples.Print_array(income_array)
	Bubble_Sort_array(income_array)
	samples.Print_array(income_array)
}

func Bubble_Sort_array(income_array []int) {
	now1 := time.Now()
	fmt.Println(now1)
	var j = 0
	var first int
	var second int
	var N = 1
	var M = 2
	for M > N {
		M = 1
		for len(income_array)-1 > j {
			first = income_array[j]
			second = income_array[j+1]
			if first > second {
				income_array[j+1] = first
				income_array[j] = second
				M = 2
			}
			j = j + 1
			//	Print_array(income_array)
		}
		j = 0
	}
	then1 := time.Since(now1)
	fmt.Println(then1)
	fmt.Println("Время пузырьковой сортировки")

}
