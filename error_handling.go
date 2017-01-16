package main

import "fmt"
import "runtime/debug"
//import "log"

func getElementAtIndex(a []int32, index int32) int32{
    return a[index]
}

func divide(a, b float32) float32{
    return a/b
}

func safelyGet(a []int32, index int32) int32{
    defer func(){
        if err := recover(); err!=nil{
            fmt.Println("Error occurred", err )
            fmt.Println(fmt.Sprintf("%s", debug.Stack()))
        }
    }()
    return getElementAtIndex(a, index)
}

func main() {
    //fmt.Println(divide(2, 0))
    a := []int32{1,2,3,4}
    //v:= getElementAtIndex(a, 4)
    v:= safelyGet(a, 4)
    fmt.Println("value of v=", v)
}
