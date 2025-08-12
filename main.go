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
		fmt.Print("Digite um número (1-100):")
		fmt.Scan(&palpite)

		if palpite == alvo {
			fmt.Printf("Você acertou! o número secreto era: %v", alvo)
			os.Exit(-1)
		} else {
			if palpite < alvo {
				fmt.Printf("O número é maior!")
			} else if palpite > alvo {
				fmt.Printf("O número é menor!")
			}
		}
		tentativas++
	}
}

func main() {

}
