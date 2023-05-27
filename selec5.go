package main

import "fmt"

func main() {
    var num int
    fmt.Println("Digite um número inteiro:")
    fmt.Scanln(&num)
    if num%3 == 0 && num%5 == 0 {
        fmt.Println("O número é múltiplo de 3 e 5 ao mesmo tempo.")
    } else {
        fmt.Println("O número não é múltiplo de 3 e 5 ao mesmo tempo.")
    }
}
