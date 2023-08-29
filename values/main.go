package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
}

func main() {

	oneNumber := 1
	oneFloat32 := 1.4321
	oneFloat64 := 1.1234
	oneString := "string"
	oneBoolean := true
	oneBoolean2 := false

	myUser := User{
		Name: "Fernando",
	}

	var emptyValueForString string
	var emptyValueForNumber int
	var emptyValueForBoolean bool

	fmt.Println("-------------------")

	fmt.Println(" typeof emptyValueForString ", reflect.TypeOf(emptyValueForString))
	fmt.Println(" typeof emptyValueForNumber ", reflect.TypeOf(emptyValueForNumber))
	fmt.Println(" typeof emptyValueForBoolean ", reflect.TypeOf(emptyValueForBoolean))
	fmt.Println(" typeof oneNumber ", reflect.TypeOf(oneNumber))
	fmt.Println(" typeof oneFloat32 ", reflect.TypeOf(oneFloat32))
	fmt.Println(" typeof oneFloat64 ", reflect.TypeOf(oneFloat64))
	fmt.Println(" typeof oneString ", reflect.TypeOf(oneString))
	fmt.Println(" typeof oneBoolean ", reflect.TypeOf(oneBoolean))
	fmt.Println(" typeof oneBoolean2 ", reflect.TypeOf(oneBoolean2))

	fmt.Println("-------------------")

	fmt.Println("emptyValueForString", emptyValueForString)
	fmt.Println("emptyValueForNumber", emptyValueForNumber)
	fmt.Println("emptyValueForBoolean", emptyValueForBoolean)
	fmt.Println("oneNumber", oneNumber)
	fmt.Println("oneFloat32", oneFloat32)
	fmt.Println("oneFloat64", oneFloat64)
	fmt.Println("oneString", oneString)
	fmt.Println("oneBoolean", oneBoolean)
	fmt.Println("oneBoolean2", oneBoolean2)

	fmt.Println("myUser", myUser)
	fmt.Println("myUser.Name", myUser.Name)
}
