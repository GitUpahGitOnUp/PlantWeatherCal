package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	BASE_URL := os.Getenv("BASE_URL")
	API_KEY := os.Getenv("API_KEY")

	fmt.Println("Where do you want to check the weather? :")
	var city string
	fmt.Scanln(&city)

	SEARCH_URL := fmt.Sprintf("%v?appid=%s&q=%s", BASE_URL, API_KEY, city)

	response, err := http.Get(SEARCH_URL)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close() // prevents data leaks, executes when main function completes

	if response.StatusCode == http.StatusOK {
		weatherBytes, _ := io.ReadAll(response.Body)
		fmt.Println(string(weatherBytes))
	}
}
