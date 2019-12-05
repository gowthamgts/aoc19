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

		subFuelNeeded := 0
		fuelNeeded := (mass / 3) - 2
		for fuelNeeded > 0 {
			subFuelNeeded += fuelNeeded
			fuelNeeded = (fuelNeeded / 3) - 2
		}
		totalFuelNeeded += subFuelNeeded
	}

	fmt.Println("Total Fuel Needed: ", totalFuelNeeded)
}
