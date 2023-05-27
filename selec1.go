package main

import "fmt"

func main() {
    var num1, num2 int
    fmt.Println("Digite dois números inteiros:")
    fmt.Scanln(&num1)
    fmt.Scanln(&num2)
    maior := num1
    if num2 > maior {
        maior = num2
    }
    fmt.Println("O maior número é:", maior)
}
