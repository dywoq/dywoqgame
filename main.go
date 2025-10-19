package main

import (
	"fmt"
	"reflect"

)

type user struct {
}

func main() {
	s := int32(10)
	fmt.Printf("%s", reflect.TypeOf(s).String())
}
