package main

import (
	"bufio"
	"fmt"
	"log"
	"github.com/afilbert/advent-of-code-2019/helpers"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(strings.NewReader(string(helpers.GetDataStringByDay())))

	log.Println("Initiating data read")

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return
	}

	log.Println("Calculating mass")

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
