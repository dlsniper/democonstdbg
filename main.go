package main

import (
	"fmt"

	"play/internal/test"
)

const AnotherConst = "demo"

func main() {
	approved := "Approved"
	ot := AnotherConst

	check(approved, ot)
}

func check(approved string, ot string) {
	if approved == test.Approved {
		fmt.Println("You're in!", ot)
	}

	if ot == AnotherConst {
		fmt.Println("it works")
	}
}
