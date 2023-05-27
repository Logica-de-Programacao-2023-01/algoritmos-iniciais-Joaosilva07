package main

import "fmt"

func main() {
    var num, count int
    soma := 0
    fmt.Println("Digite vários números inteiros (digite 0 para parar):")
    for {
        fmt.Scanln(&num)
        if num == 0 {
            break
        }
        soma += num
        count++
    }
    if count > 0 {
        media := float64(soma) / float64(count)
        fmt.Printf("Média aritmética: %.2f\n", media)
    } else {
        fmt.Println("Nenhum número foi inserido.")
    }
}
