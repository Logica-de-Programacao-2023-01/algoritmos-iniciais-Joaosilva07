package main

import "fmt"

func main() {
    var num1, num2 int
    fmt.Println("Digite dois números inteiros:")
    fmt.Scanln(&num1)
    fmt.Scanln(&num2)
    if num1 >= 0 && num2 >= 0 {
        multiplicacao := num1 * num2
        fmt.Println("Resultado da multiplicação:", multiplicacao)
    } else {
        soma := num1 + num2
        fmt.Println("Resultado da soma:", soma)
    }
}
