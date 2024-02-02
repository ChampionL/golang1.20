package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 10
	typeOfA := reflect.TypeOf(a).String()
	fmt.Println(typeOfA)
}
