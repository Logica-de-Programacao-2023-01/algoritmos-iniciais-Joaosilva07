package main

import "fmt"

func main() {
    var num int
    fmt.Println("Digite um número inteiro positivo:")
    fmt.Scanln(&num)

    fmt.Printf("Divisores de %d: ", num)
    for i := 1; i <= num; i++ {
        if num%i == 0 {
            fmt.Printf("%d ", i)
        }
    }
    fmt.Println()
}
