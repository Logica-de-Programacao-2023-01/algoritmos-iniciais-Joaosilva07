package main

import "fmt"

func main() {
    var nota1, nota2 float64
    fmt.Println("Digite a primeira nota:")
    fmt.Scanln(&nota1)
    fmt.Println("Digite a segunda nota:")
    fmt.Scanln(&nota2)
    media := (nota1 + nota2) / 2
    if media >= 6 {
        fmt.Println("Aluno aprovado.")
    } else {
        fmt.Println("Aluno reprovado.")
    }
}
