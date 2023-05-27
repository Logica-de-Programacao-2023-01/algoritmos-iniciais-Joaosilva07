package main

import "fmt"

func main() {
    var diasTrabalhados int
    var valorDiaria float64
    fmt.Println("Digite o número de dias trabalhados:")
    fmt.Scanln(&diasTrabalhados)
    fmt.Println("Digite o valor da diária:")
    fmt.Scanln(&valorDiaria)
    salario := float64(diasTrabalhados) * valorDiaria
    fmt.Printf("O salário é: %.2f\n", salario)
}
