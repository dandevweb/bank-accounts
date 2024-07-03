package main

import "fmt"

type CurrentAccount struct {
	holder  string
	agency  int
	account int
	balance float64
}

func main() {
	firstAccount := CurrentAccount{"Danilo", 12, 123232, 10.90}
	secondAccount := CurrentAccount{"Maria", 12, 123232, 10.90}

	fmt.Println(firstAccount)
	fmt.Println(secondAccount)
}
