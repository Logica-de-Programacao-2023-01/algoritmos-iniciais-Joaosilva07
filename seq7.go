package main

import "fmt"

func main() {
    var num int
    fmt.Println("Digite um número inteiro:")
    fmt.Scanln(&num)
    antecessor := num - 1
    sucessor := num + 1
    fmt.Println("Antecessor:", antecessor)
    fmt.Println("Sucessor:", sucessor)
}
