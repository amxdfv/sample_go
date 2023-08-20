package main

import (
	"os"
	"oslic_v1/custom_sort"
	"oslic_v1/samples"
)

func main() {
	base()
}

func base() {
	var initial = samples.Init_array()
	var turn = ""
	if len(os.Args) > 1 {
		turn = os.Args[1]
	}
	switch {
	case turn == "one":
		samples.Sort_array(initial)
	case turn == "two":
		custom_sort.Custom_Sort_v1(initial)
	default:
		samples.Sort_array(initial)
	}

}
