package main

import (
	"fmt"
	"oslic_v1/samples"
)

func main() {

	var initial []int = samples.Init_array()
	initial_v2 := samples.Init_array()
	samples.Custom_Sort_v1(initial_v2)

	samples.Sort_array(initial)

	//var i = CountOslic(1, 2)
	//	fmt.Println(i)
	fmt.Println("hello world")
}
