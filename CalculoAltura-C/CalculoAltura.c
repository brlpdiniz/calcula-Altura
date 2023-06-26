#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(){

    char NomePai[256], NomeMae[256];
    float AlturaPai, AlturaMae;
    float Menino, Menina;

    printf("Digite o NOME do pai: ");
    scanf("%s", &NomePai);
    printf("Digite a ALTURA do pai: ");
    scanf("%f", &AlturaPai);
    printf("Digite o NOME da mae: ");
    scanf("%s", &NomeMae);
    printf("Digite a ALTURA da mae: ");
    scanf("%f", &AlturaMae);

    Menino = (AlturaPai + AlturaMae + 0.13) / 2;
    Menina = ((AlturaPai - 0.13) + AlturaMae) / 2;

    printf("O filho do %s e da %s terah %.2f de altura\n", NomePai, NomeMae, Menino);
    printf("A filha do %s e da %s terah %.2f de altura\n", NomePai, NomeMae, Menina);

    return 0;
}