package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	interval := 5 * time.Second

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		num := rand.Intn(100)
		fmt.Println("Checking comic: ", num)

		if checkComic(num) {
			fmt.Println("Comic already exists")
		} else {
			fmt.Println("Comic missing, requesting download")
			requestDownload(num)

			//loops and stalls until comic finishes downloading
			for !checkComic(num) {
				fmt.Println("Waiting for comic to download")
				time.Sleep(2 * time.Second)
			}
			fmt.Println("Comic ", num, " downloaded")
		}
	}
}

func checkComic(id int) bool {
	url := fmt.Sprintf("http://localhost:8080/comic/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error checking comic: ", err)
		return false
	}
	defer resp.Body.Close()

	var status struct {
		Downloaded    bool `json:"downloaded"`
		IsDownloading bool `json:"isdownloading"`
	}

	err = json.NewDecoder(resp.Body).Decode(&status)
	if err != nil {
		fmt.Println("Error decoding response: ", err)
		return false
	}
	return status.Downloaded
}

func requestDownload(id int) {
	url := fmt.Sprintf("http://localhost:8080/comic/%d", id)
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		fmt.Println("Error requesting download: ", err)
		return
	}

	defer resp.Body.Close()
	fmt.Println("Requested download for comic: ", id)

}
