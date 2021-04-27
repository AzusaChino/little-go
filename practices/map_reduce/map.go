package map_reduce

import (
	"fmt"
	"strings"
)

func MapStrToStr(arr []string, fn func(string) string) []string {
	var newArray []string
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}
	return newArray
}

func MapStrToInt(arr []string, fn func(s string) int) []int {
	var newArray []int
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}
	return newArray
}

func Reduce(arr []string, fn func(string) int) int {
	sum := 0
	for _, it := range arr {
		sum += fn(it)
	}
	return sum
}

func Filter(arr []int, fn func(int) bool) []int {
	var newArray []int
	for _, it := range arr {
		if fn(it) {
			newArray = append(newArray, it)
		}
	}
	return newArray
}

func main1() {
	var list = []string{"Hao", "Chen", "MegaEase"}
	x := MapStrToStr(list, func(s string) string {
		return strings.ToUpper(s)
	})
	fmt.Printf("%v\n", x)
	//["HAO", "CHEN", "MEGAEASE"]
	y := MapStrToInt(list, func(s string) int {
		return len(s)
	})
	fmt.Printf("%v\n", y)
	//[3, 4, 8]
	z := Reduce(list, func(s string) int {
		return len(s)
	})
	fmt.Printf("%v\n", z)

	var intset = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	out := Filter(intset, func(n int) bool {
		return n%2 == 1
	})
	fmt.Printf("%v\n", out)
	out = Filter(intset, func(n int) bool {
		return n > 5
	})
	fmt.Printf("%v\n", out)
}
