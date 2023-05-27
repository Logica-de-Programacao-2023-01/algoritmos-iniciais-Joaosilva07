package main

import "fmt"

func main() {
    var num, maior int
    fmt.Println("Digite vários números inteiros (digite 0 para parar):")
    for {
        fmt.Scanln(&num)
        if num == 0 {
            break
        }
        if num > maior {
            maior = num
        }
    }
    if maior != 0 {
        fmt.Println("Maior número:", maior)
    } else {
        fmt.Println("Nenhum número foi inserido.")
    }
}
