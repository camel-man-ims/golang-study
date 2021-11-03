package main

import (
	"basic"
	"fmt"
)

func main() {
	books := []string{"java", "python"}
	person := basic.Person{Name: "ersu", Age: 10, Books: books}
	fmt.Println(person)
}
