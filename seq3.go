package main

import "fmt"

func main() {
    var peso, altura float64
    fmt.Println("Digite o peso (em kg):")
    fmt.Scanln(&peso)
    fmt.Println("Digite a altura (em metros):")
    fmt.Scanln(&altura)
    imc := peso / (altura * altura)
    fmt.Printf("O IMC é: %.2f\n", imc)
}
