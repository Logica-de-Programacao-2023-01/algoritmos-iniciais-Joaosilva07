package main

import "fmt"

func main() {
    var idadeAnos int
    fmt.Println("Digite a idade em anos:")
    fmt.Scanln(&idadeAnos)
    idadeDias := idadeAnos * 365
    fmt.Println("A idade em dias é:", idadeDias)
}
