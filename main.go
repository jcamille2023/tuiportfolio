package main

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

func main() {
	// printing hello world
	fmt.Println("Hello, world!")
	for j := 0; j < 10; j++ { // for loops
		fmt.Println(j)
	}
	if 1 == 1 { // if conditions
		fmt.Println("1 does in fact equal 1. Excellent.")
	}

	fmt.Println("The time is", time.Now())

	switch time.Now().Weekday() { // switch case
	case time.Saturday:
		fmt.Println("It's Saturday!")
	default:
		fmt.Println("It's not Saturday.")
	}
	// arrays
	arr := [3]int{1, 2, 3}
	fmt.Println("New array: ", arr)
	//slices
	slice := []int{1, 2, 3}
	fmt.Println("New slice: ", slice)
	//multi-dimensional slice
	twoD := [][]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	// iterating through the slice
	for i := 0; i < len(twoD); i++ {
		fmt.Println(twoD[i])
	}
	// initalizing a map (dictionary)
	dict := make(map[string]string)
	dict["me and boat"] = "getting to the bucks all day"
	dict["me and gmoney pouring up"] = "hi-tec always"
	dict["they was hating on me back in the hs days"] = "everything's changed everythings so fake"
	fmt.Println(dict)
	fmt.Println("Length of dict is ", len(dict))
	fmt.Printf("Length of dict is %d\n", len(dict))
	// range
	for _, num := range slice {
		fmt.Println(num)
	}
	// self defined print fucntion, everything must be of string type
	print("ur mom got bent over ", dict, " too many times.")
}

func print(a ...interface{}) {
	for _, item := range a {
		if reflect.TypeOf(item) != reflect.TypeOf("string") {
			item = str(item)
			fmt.Printf("%s", item)
		} else {
			fmt.Printf("%s", item)
		}
	}
	fmt.Printf("\n")
}

func str(a interface{}) string {
	var new_string string
	if reflect.TypeOf(a) == reflect.TypeOf(2) {
		new_string = fmt.Sprintf("%d", a)
	} else if reflect.TypeOf(a) == reflect.TypeOf(2.2) {
		new_string = fmt.Sprintf("%f", a)
	} else if reflect.TypeOf(a).Kind() == reflect.Map {
		new_string += "map["
		v := reflect.ValueOf(a)
		for i, key := range v.MapKeys() {
			val := v.MapIndex(key)
			if i > 0 {
				if i < v.Len()-1 {
					new_string += ","
				}
				new_string += " "

			}
			new_string += fmt.Sprintf("%v:%v", key, val)
		}
		new_string += "]"

	} else if reflect.TypeOf(a).Kind() == reflect.Array {
		new_string += "["
		v := reflect.ValueOf(a)
		for i := 0; i < v.Len(); i++ {
			new_string += fmt.Sprintf("%v", v.Index(i))
			if i < v.Len()-1 {
				new_string += ", "
			}
		}
		new_string += "]"
	} else {
		panic(errors.New("str function given unsupported type"))
	}
	return new_string
}
