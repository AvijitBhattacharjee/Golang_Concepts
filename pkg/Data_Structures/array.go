// Copyright (c) avijit bhattacharjee 2024

package Data_Structures

import "fmt"

func PrintArray() { // Change main to a different function name
    arr := [4]string{"geek", "gfg", "Geeks1231", "GeeksforGeeks"}
    fmt.Println("Elements of the array:")
    for i := 0; i < 3; i++ {
        fmt.Println(arr[i])
    }
}