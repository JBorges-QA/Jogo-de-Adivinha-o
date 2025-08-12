package main

import (
	"fmt"
	"math/rand/v2"
	"os"
)

func Adivinhacao() {
	alvo := rand.IntN(100) + 1
	tentativas := 0

	for {
		var palpite int
		fmt.Print("Digite um número (1-100): ")
		fmt.Scan(&palpite)

		if palpite == alvo {
			fmt.Println("Você acertou! o número secreto era: ", alvo)
			fmt.Println("Número de tentativas:", tentativas+1)
			os.Exit(-1)
		} else {
			if palpite < alvo {
				fmt.Println("O número é maior!")
			} else if palpite > alvo {
				fmt.Println("O número é menor!")
			}
		}
		tentativas++
	}
}

func main() {
	Adivinhacao()
}
