package main

import "fmt"

func main() {
    var pesoQuilos float64
    fmt.Println("Digite o peso em quilos:")
    fmt.Scanln(&pesoQuilos)
    pesoLibras := pesoQuilos * 2.20462
    fmt.Printf("O peso em libras é: %.2f\n", pesoLibras)
}
