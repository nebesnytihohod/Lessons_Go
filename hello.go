package main

import (
	"fmt"
	"time"
)

var name, surname string

func init() {
	name = "John"
}
func init() {
	if surname == "" {
		surname = "Doe"
	}
}

var isGlobalVar string = "Hello, Max!"

func myFunctionOne(x, y int, z string) (result1, result2 int, result3 string) {
	//func body
	var returnAsResult1 int
	returnAsResult1 = x + y

	var returnAsResult2 = x - y

	returnAsResult3 := z + "_"

	if x > 5 {
		result2 = x - y
	} else if y > 5 {
		result2 = y - x
	} else {
		result2 = x + 10
	}

	return returnAsResult1, returnAsResult2, returnAsResult3
}

func myFunctionSideEffectsNoReturn(str1, str2 string) {
	fmt.Println("string 1: ", str1, " String 2: ", str2)
}

func myFunctionTwo(str1, str2 string) string {
	startTime := time.Now()
	defer func(st time.Time) { println("Время выполнения функции: ", time.Now().Sub(st)) }(startTime)
	println("--------start func--------------------------")
	defer println("defered output 1")
	for i := 1; i <= 3; i++ {
		defer println("defered output 2: i = ", i)
	}
	println(str1)
	println(str2)
	println("--------end func--------------------------")

	return str1 + " " + str2
}

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

	var myMap1 = make(map[string]string)
	myMap1["a"] = "A"
	var myMap10 map[string]string
	myMap10 = make(map[string]string)
	var myMap11 = map[int]string{11: "one", 12: "two"}
	myMap2 := make(map[int]string)
	myMap2[2] = "AA"
	myMap21 := map[int]string{21: "one", 22: "two", 33: ""}

	type MyMap map[string]int64
	var map3 MyMap
	map3 = make(MyMap, 7)

	var m1 = myMap1["a"]
	m2 := myMap2[2]
	m11, isFilledMyMap11 := myMap11[11]
	m21, isFilledMyMap21 := myMap21[22]
	m211, isFilledMyMap211 := myMap21[33]
	m2111, isFilledMyMap2111 := myMap21[44]

	fmt.Println(myMap1, myMap2, map3)
	fmt.Println(myMap11, myMap21)
	fmt.Println(myMap1["a"], m1, m2, m11, isFilledMyMap11, m21, isFilledMyMap21, m211, isFilledMyMap211, m2111, isFilledMyMap2111)

	if myMap10 == nil {
		println("Map is not initialized!")
	} else {
		m10, isFilledMyMap10 := myMap10[""]
		println("myMap10: ", myMap10, " ", "m10: ", m10, " ", "Filled?: ", isFilledMyMap10)
	}

	for k, v := range myMap11 {
		fmt.Printf("Ключ %v, имеет значение %v \n", k, v)
	}

	f1, f2, _ := myFunctionOne(1, 2, "fuck")
	fmt.Println(f1, f2)
	_, _, f3 := myFunctionOne(6, 0, "fuck")
	fmt.Println(f1, f2, f3)
	f1, f2, f3 = myFunctionOne(1, 1, "fuck")
	fmt.Println(f1, f2, f3)

	myFunctionSideEffectsNoReturn("Hello", "Max")

	//undeclared and unnamed function = literal form = anonymous function
	lambdaFunc := func(s string) string { return s }
	myFunctionSideEffectsNoReturn("Hello", lambdaFunc("Masha"))

	fullName := myFunctionTwo(name, surname)
	fmt.Println("Fullname: ", fullName)
}
