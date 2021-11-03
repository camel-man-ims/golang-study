package basic

import "fmt"

func DoPointer() {
	a := 3
	var b *int = &a
	fmt.Println(*b)
}

func Arr() {
	arr := []int{1, 2, 3, 4}

	arr = append(arr[:1], arr[2:]...)

	fmt.Println(arr)
}

func ThisIsMap() {
	ma := map[string]string{"a": "b"}
	fmt.Println(ma)
}
