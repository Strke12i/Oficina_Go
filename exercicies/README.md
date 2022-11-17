## Exercicios da oficina de Golang

##### 1 - Implementação de um programa que le uma string e retorna valor de cada caracter na tabela ascii:
Nesse exercicío, você deve criar um programa que lê uma string como argumento de linha 
de comando e retorna uma tabela com os valores de cada caracter em ascii presente na string

| Caracter | Valor |
|----------|-------|
| G        | 71    |
| o        | 111   |

    go run exercicio1.go teste

##### 2 - Passando arrays para funções e alterando-os
Faça uma função que recebe um array de 10 inteiros nela você precisa fazer a soma de cada valor multiplicado pelos demais elementos, ao final você precisa alterar o elemento inicial para o ser o resultado da soma.

    test := [10]int{0,1,2,3,4,5,6,7,8,9}

##### 3 - Ler um arquivo json e escrevendo em um txt
Leia um arquivo json e armazena seus valores em um map. Depois percorra esse map escrevendo em cada linha de
arquivo txt o seguinte trecho : "Me chama [nome], tenho [idade] e curso [curso]."


