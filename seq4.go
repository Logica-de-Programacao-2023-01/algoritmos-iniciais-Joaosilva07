package main

import "fmt"

func main() {
    var num1, num2, num3 float64
    fmt.Println("Digite três números reais:")
    fmt.Scanln(&num1)
    fmt.Scanln(&num2)
    fmt.Scanln(&num3)
    mediaPonderada := (num1*2 + num2*3 + num3*5) / 10
    fmt.Printf("A média ponderada é: %.2f\n", mediaPonderada)
}
