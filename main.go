package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/drpaij0se/cligpt/cli"
)

type Data struct {
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func main() {
	var err error
	var config map[string]string
	if err = cli.CreateConfigDirectory(); err != nil {
		log.Fatal(err)
	}

	if config, err = cli.ReadYml(); err != nil {
		log.Fatal(err)
	}

	if len(config["auth"]) < 51 {
		log.Fatal("Ensure to insert a valid token in cligpt.yml file.")
	}
	client := &http.Client{}
	var data = strings.NewReader(`{
		  "model": "` + config["model"] + `",
		  "messages": [{"role": "user", "content": "` + os.Args[1] + `"}],
		  "temperature": 0.7
		}`)
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", `Bearer `+config["auth"]+``)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("The token is valid?", err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var response Data
	json.Unmarshal(bodyText, &response)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(response.Choices[0].Message.Content)
}
