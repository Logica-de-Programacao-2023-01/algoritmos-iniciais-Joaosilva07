package main

import "fmt"

func main() {
    var num1, num2, num3 int
    fmt.Println("Digite três números inteiros:")
    fmt.Scanln(&num1)
    fmt.Scanln(&num2)
    fmt.Scanln(&num3)
    soma := num1 + num2 + num3
    fmt.Println("A soma dos números é:", soma)
}
