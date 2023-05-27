package main

import "fmt"

func main() {
    var num1, num2, num3 float64
    fmt.Println("Digite três números reais:")
    fmt.Scanln(&num1)
    fmt.Scanln(&num2)
    fmt.Scanln(&num3)

    if num1 <= num2 && num1 <= num3 {
        fmt.Print(num1, " ")
        if num2 <= num3 {
            fmt.Println(num2, num3)
        } else {
            fmt.Println(num3, num2)
        }
    } else if num2 <= num1 && num2 <= num3 {
        fmt.Print(num2, " ")
        if num1 <= num3 {
            fmt.Println(num1, num3)
        } else {
            fmt.Println(num3, num1)
        }
    } else {
        fmt.Print(num3, " ")
        if num1 <= num2 {
            fmt.Println(num1, num2)
        } else {
            fmt.Println(num2, num1)
        }
    }
}
