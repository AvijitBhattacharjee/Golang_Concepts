// Copyright (c) avijit bhattacharjee 2024

package Variadic

import (
	"fmt"
)

// The function that is called with the varying number of arguments is known as variadic function. 
// Or in other words, a user is allowed to pass zero or more arguments in the variadic function. 
// fmt.Printf is an example of the variadic function, it required one fixed argument at the starting after that it can accept any number of arguments. 

func Variadic() {

	fmt.Println("Example of variadic function")
	variadic_example("string1")
	variadic_example("string1", "string2")
	variadic_example("string1", "string2", "string3")

}

// Important Points: In the declaration of the variadic function, the type of the last parameter is preceded by an ellipsis, i.e, (…). 
// It indicates that the function can be called at any number of parameters of this type. 
func variadic_example(elements ... string) {
	fmt.Println(elements[len(elements)-1])
}
