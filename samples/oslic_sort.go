package samples

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

func Init_array() []int {
	var sort_numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 1, 61, 7, 18, 19, 20}
	var j = 0
	for len(sort_numbers) > j {
		sort_numbers[j] = rand.Intn(200) + 1
		j = j + 1
	}
	//	Print_array(sort_numbers)
	return sort_numbers
}

func Sort_array(income_array []int) {
	Print_array(income_array)

	now := time.Now()
	fmt.Println(now)
	sort.Ints(income_array)
	then := time.Now()
	res := then.Sub(now)
	fmt.Println(then)
	fmt.Println(res)
	fmt.Println("Время встроенной сортировки")
	Print_array(income_array)
}

func Custom_Sort_v1(income_array []int) {

	Bubble_Sort_array(income_array)

	Print_array(income_array)
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
	then1 := time.Now()
	res1 := then1.Sub(now1)
	fmt.Println(then1)
	fmt.Println(res1)
	fmt.Println("Время пузырьковой сортировки")

}

func Print_array(income_array []int) {
	var j = 0
	var j_s string
	for len(income_array) > j {
		j_s = j_s + strconv.Itoa(income_array[j]) + " "
		j = j + 1
	}
	println(j_s)
}

func ReInitiate_array(income_array []int) []int {
	var new_array = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 1, 61, 7, 18, 19, 20}
	var i = 0
	for len(income_array) > i {
		new_array[i] = income_array[i]
		i = i + 1
	}
	return new_array
}
