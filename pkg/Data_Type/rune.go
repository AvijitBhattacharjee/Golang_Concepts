// Copyright (c) avijit bhattacharjee 2024

package Data_Type

import (
	"fmt"
	"unsafe"
	"reflect"
)

func Rune() {
	fmt.Println("Implementing rune example")

	var r rune
	r = '@'
	fmt.Println("size %d", unsafe.Sizeof(r))
	fmt.Println("type %s", reflect.TypeOf(r))
	fmt.Println(r)

	var r1 = "ab%"
	var r2 = []rune(r1)
	fmt.Println(r2)

	var r3 = []rune{'x', 'y', 'z'}
	r4 := string(r3)
	fmt.Println(r4)


}
