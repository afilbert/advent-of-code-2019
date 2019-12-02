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

		totalFuel += doNotBeADumbElf(float64(modulePayloadMass), 0.0)
	}

	fmt.Printf("Total fuel required: %d\n", int64(totalFuel))
}

// Recurse yourself.
func doNotBeADumbElf(mass float64, fuelAccum float64) float64 {
	// Hey dog. I heard your fuel likes fuel, so I'm adding some fuel to your fuel so your fuel can adequately fuel while you fuel.
	massFuelRequirement := math.Floor(mass / 3.0) - 2
	if massFuelRequirement > 0 {
		fuelAccum += massFuelRequirement
		return doNotBeADumbElf(massFuelRequirement, fuelAccum)
	} else {
		return fuelAccum
	}
}