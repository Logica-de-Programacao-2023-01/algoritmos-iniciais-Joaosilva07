package main

import "fmt"

func main() {
    var altura float64
    var sexo string
    fmt.Println("Digite a altura (em metros):")
    fmt.Scanln(&altura)
    fmt.Println("Digite o sexo (M para masculino ou F para feminino):")
    fmt.Scanln(&sexo)

    pesoIdeal := 0.0
    if sexo == "M" {
        pesoIdeal = 72.7*altura - 58
    } else if sexo == "F" {
        pesoIdeal = 62.1*altura - 44.7
    }

    var situacao string
    if pesoIdeal < 0 {
        situacao = "abaixo do peso ideal"
    } else if pesoIdeal > 0 {
        situacao = "acima do peso ideal"
    } else {
        situacao = "no peso ideal"
    }

    fmt.Printf("A pessoa está %s.\n", situacao)
}
