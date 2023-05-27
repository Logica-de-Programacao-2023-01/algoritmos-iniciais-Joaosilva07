package main

import "fmt"

func main() {
    var salario float64
    fmt.Println("Digite o salário do funcionário:")
    fmt.Scanln(&salario)
    novoSalario := salario * 1.15
    fmt.Printf("O novo salário com aumento de 15%% é: %.2f\n", novoSalario)
}
