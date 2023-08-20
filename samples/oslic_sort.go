package samples

import (
	"fmt"
	"log"
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
	return sort_numbers
}

func Sort_array(income_array []int) {
	Print_array(income_array)
	defer duration(track("Sort_array"))

	now := time.Now()
	sort.Ints(income_array)
	then := time.Since(now)
	fmt.Println(then)
	fmt.Println("Время встроенной сортировки")
	Print_array(income_array)
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

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
