package helpers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

const AOCYear = "2019"

func GetDataStringByDay() string {
	req, err := http.NewRequest("GET", GetAOCUrl(), nil)
	if err != nil {
		log.Fatalln(err)
		return ""
	}

	log.Println("Logging in with token: \n", os.Getenv("AOC_SESSION"))

	req.AddCookie(&http.Cookie{Name: "session", Value: os.Getenv("AOC_SESSION")})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return ""
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return ""
	}
	return string(body)
}

func GetAOCUrl() string {
	// Get path folder
	ex, err := os.Executable()
	if err != nil {
		log.Fatalln(err)
		return ""
	}

	exPath := filepath.Dir(ex)

	log.Println("Exec folder: ", exPath)

	intVal, err := strconv.Atoi(exPath)
	if err != nil {
		log.Fatalln(err)
		return ""
	}

	return fmt.Sprintf("https://adventofcode.com/%s/day/%s/input", AOCYear, strconv.Itoa(intVal))
}
