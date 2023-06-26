package main

import (
	"fmt"      // stdio.h
)

func main(){

    var NomePai, NomeMae string
    var AlturaPai, AlturaMae float32

    fmt.Print("Digite o NOME do pai: ");
    fmt.Scanln(&NomePai);
    fmt.Print("Digite a ALTURA do pai: ");
    fmt.Scanln(&AlturaPai);
    fmt.Print("Digite o NOME da mãe: ");
    fmt.Scanln(&NomeMae);
    fmt.Print("Digite a ALTURA da mãe: ");
    fmt.Scanln(&AlturaMae);

    Menino := (AlturaPai + AlturaMae + 0.13) / 2;
    Menina := ((AlturaPai - 0.13) + AlturaMae) / 2;

    fmt.Printf("O filho do %s e da %s terá %.2f de altura\n", NomePai, NomeMae, Menino);
    fmt.Printf("A filha do %s e da %s terá %.2f de altura\n", NomePai, NomeMae, Menina);

}
