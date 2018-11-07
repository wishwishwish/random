package main

import "fmt"

func main(){
    var age int
    fmt.Println("my age is", age)
    age = 29
    fmt.Println("my age is", age)
    age = 54
    fmt.Println("my age is", age)

    var new_age int = 29
    fmt.Println("my new age is", new_age)

    var no_type_age = 29
    fmt.Println("my no type age is", no_type_age)

    var width, height int = 100, 50
    fmt.Println("width is", width, "height is", height)

    var (
        vname = "naveen"
        vage = 29
        vheight int
    )
    fmt.Println("my name is", vname, "age is", vage, "height is", vheight)

    shorthand_name, shorthand_age := "naveen", 29
    fmt.Println("my name is", shorthand_name, "age is", shorthand_age)
}
