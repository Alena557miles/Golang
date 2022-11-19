package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var temperature float32
var pain string
var health string

func main() {
	fmt.Print("------------------------- iHealth -------------------------\n \n")
	pain := askCharacters("tell me, where your pain is concentrated? Your answer:")
	temperature := askCharacters("What is your body temperature? Your answer:")
	health := askCharacters("Do you want to became a healthy men/women? ")

	fmt.Println("---------------------------------------------------------------")
	fmt.Printf("\nYour feel pain in : %s, your temperature is : %s , and want you think about healthy life : %s \n", pain, temperature, health)
	fmt.Print("      Keep calm and take care ;) \n \n")
	fmt.Println("---------------------------------------------------------------")
}

func askCharacters(messege string) any {
	fmt.Println(messege)
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	line := scanner.Text()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	return line
}
