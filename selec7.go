package main

import "fmt"

func main() {
    var salario float64
    fmt.Println("Digite o salário do funcionário:")
    fmt.Scanln(&salario)
    if salario < 1000 {
        novoSalario := salario * 1.1
        fmt.Printf("O novo salário com aumento de 10%% é: %.2f\n", novoSalario)
    } else {
        novoSalario := salario * 1.05
        fmt.Printf("O novo salário com aumento de 5%% é: %.2f\n", novoSalario)
    }
}
