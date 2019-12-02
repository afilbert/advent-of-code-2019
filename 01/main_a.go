package main

import (
	"bufio"
	"fmt"
	"github.com/afilbert/advent-of-code-2019/helpers"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(strings.NewReader(string(helpers.GetDataStringByDay())))
	totalFuel := float64(0)

	log.Println("Initiating data read")

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return
	}

	log.Println("Calculating mass and fuel requirements")

	for scanner.Scan() {
		modulePayloadMass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
			return
		}

		moduleFuelRequirement := math.Floor(float64(modulePayloadMass) / 3.0) - 2
		totalFuel += moduleFuelRequirement
	}

	fmt.Printf("Total fuel required: %d\n", int64(totalFuel))
}
