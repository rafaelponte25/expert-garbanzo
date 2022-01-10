package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
)

func main() {
	data, err := ioutil.ReadFile("./data.json")
	fmt.Println(string(data))
	fmt.Println(err)
	fmt.Println(reflect.TypeOf(data))
}
