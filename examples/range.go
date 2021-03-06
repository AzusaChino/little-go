package examples

import "fmt"

func init() {

	nums := []int{2, 3, 4}
	sum := 0

	for _, n := range nums {
		sum += n
	}
	fmt.Println(sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
