package helpers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
)

const AOCYear = "2019"
const AOCBaseURL = "https://adventofcode.com/%s/day/%s/input"

func GetDataStringByDay() string {
	req, err := http.NewRequest("GET", getAOCUrlFromCallerDirectory(), nil)
	if err != nil {
		log.Fatalln(err)
		return ""
	}

	log.Println("Logging in with token:", os.Getenv("AOC_SESSION"))

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

func getAOCUrlFromCallerDirectory() string {
	_, path, _, ok := runtime.Caller(2)

	if !ok {
		log.Fatalln("Couldn't derive directory number.")
		return ""
	}

	log.Println("Caller directory:", path)

	daySlice := strings.Split(path, "/")
	day := daySlice[len(daySlice) - 2]
	dayInt, err := strconv.Atoi(day)

	if err != nil {
		log.Fatalln(err)
		return ""
	}

	log.Println("Derived day:", strconv.Itoa(dayInt))

	return fmt.Sprintf(AOCBaseURL, AOCYear, strconv.Itoa(dayInt))
}
