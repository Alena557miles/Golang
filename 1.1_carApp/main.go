package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var brand string
var model string
var volume float32

func main() {
	fmt.Print("------------------------- CAR APP -------------------------\n \n")
	brand := askCharacters("please, tell us about car brand:")
	model := askCharacters("please, tell us about car model:")
	volume := askCharacters("please, tell us about engine volume:")
	fmt.Println("---------------------------------------------------------------")
	fmt.Printf("\nCar brand : %s, car model : %s , engine volume : %s \n", brand, model, volume)
	fmt.Print("      You a have a nice car! ;) \n \n")
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
