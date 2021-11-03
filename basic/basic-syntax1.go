package basic

import (
	"fmt"
)

// date : 21.10.15

func getVariables() string {
	name := "안녕"
	return name
}

func getNakedReturns() (a1 string, a2 string) {
	return "a1", "a2"
}

func getIntArray() []int {
	array := []int{1, 2, 3, 4, 5, 6}
	return array
}

func forLoopV1() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func forLoopV2(numbers []int) {
	for value, index := range numbers {
		fmt.Println(value, index)
	}
}
