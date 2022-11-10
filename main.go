package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	target := rand.Intn(100) + 1

	fmt.Println("Я загадал случайное число от 1 до 100")
	fmt.Println("Угадаешь его?")

	reader := bufio.NewReader(os.Stdin)
	success := false
	for guesses := 10; guesses > 0; guesses-- {
		fmt.Println("У тебя осталось", guesses, "попыток")

		fmt.Println("Введи новое число: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)

		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess < target {
			fmt.Println("Ой. Загаданное число больше введенного")
		} else if guess > target {
			fmt.Println("Ой. Загаданное число меньше введенного")
		} else {
			fmt.Println("Ура! Загаданное число угадано!!!")
			success = true
			break
		}
	}

	if !success {
		fmt.Println("Ты не смог угадать число... :-(")
	}
}
