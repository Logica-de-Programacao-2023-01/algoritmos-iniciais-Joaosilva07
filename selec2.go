package main

import "fmt"

func main() {
    var num1, num2, num3 int
    fmt.Println("Digite três números inteiros:")
    fmt.Scanln(&num1)
    fmt.Scanln(&num2)
    fmt.Scanln(&num3)
    menor := num1
    if num2 < menor {
        menor = num2
    }
    if num3 < menor {
        menor = num3
    }
    fmt.Println("O menor número é:", menor)
}
