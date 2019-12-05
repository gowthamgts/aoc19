package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	totalFuelNeeded := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("Unable to convert string to int.")
		}

		totalFuelNeeded += (mass / 3) - 2
	}

	fmt.Println("Total Fuel Needed: ", totalFuelNeeded)
}
