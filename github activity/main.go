package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

type Event struct {
	Type      string    `json:"type"`
	Repo      Repo      `json:"repo"`
	CreatedAt time.Time `json:"created_at"`
}

type Repo struct {
	Name string `json:"name"`
}

func main() {
	fmt.Println("Welcome To GitHub Activity Fetcher.")
	fmt.Println("This is simple CLI that uses GitHub APIs to fetch Users Recent Activity")
	fmt.Println("Type Your Username to start:")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">")
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("wrong input")
		}
		input = strings.TrimSpace(input)
		url := fmt.Sprintf("https://api.github.com/users/%s/events", input)
		resp, err := http.Get(url)

		if err != nil {
			fmt.Println("error in fetching username")
			break
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Println("github responded : ", resp.Status)
		}

		var events []Event

		if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
			fmt.Println("error in decoding json")
		}

		for _, event := range events {
			fmt.Printf("[%s] %s on %s\n", event.CreatedAt.Format("2006-01-02 15:04:05"), event.Type, event.Repo.Name)
		}
	}
}
