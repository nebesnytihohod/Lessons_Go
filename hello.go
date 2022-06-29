package main

import (
	"fmt"
)

var isGlobalVar string = "Hello, Max!"

func main() {
	fmt.Println("Hello, world!")
	fmt.Println(isGlobalVar)

	var AExportedVar string
	AExportedVar = "variable AExportedVar"

	var b string = "variable b"
	var c = `variable 
	
	---->> c`
	dIsOnlyLocal := "variable d"

	println(AExportedVar)
	println(b)
	println(c)
	println(dIsOnlyLocal)

	var (
		firstName, lastName string
		ago                 int
		life                bool = true
	)

	println("Full name: ", firstName, " ", lastName)
	println(ago)
	println(life)

	ago = 36

	if !life {
		fmt.Println("Life is life")
	} else if deltaAgo := 5; (ago + deltaAgo) > 40 {
		fmt.Println("Show must go on!")
	} else {
		fmt.Println("WTF?")
	}

	firstName = "Masha"

	switch firstName {
	case "Max":
		fmt.Println("Father")
	case "Masha":
		fmt.Println("Mother")
	default:
		fmt.Println("Who is it?")
	}

	switch {
	case ago == 6:
		fmt.Println("My name is iVan")
	case ago > 35:
		fmt.Println("We are parents")
	case ago < 0:
		fmt.Println("Who are you?")
	}

	switch fullName := firstName + lastName; {
	case fullName == "":
		fmt.Println("No such person")
	case fullName == "Masha":
		fmt.Println("Maaashaaaa!")
		fallthrough
	case fullName == "Max":
		fmt.Println("!!!")
	default:
		fmt.Println("Who is there?")
	}

	println(ago)
	for i := 1; i < 12; i++ {
		if i%2 != 0 {
			println(i)
			continue
		}

		if ago > 50 {
			println(i, "break", ago)
			break
		}

		ago += i
		println(i, " ", ago)
	}
	println(ago)

	var pointerAgo *int
	pointerAgo = &ago
	println("pointer to address of variable ago: ", pointerAgo)
	println("value of variable ago read across pointer: ", *pointerAgo)

	type DUL int64
	pointerDUL := new(DUL)
	println("pointer to address of my type: ", pointerDUL)
}
