package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/slack-go/slack"
)

func main() {
	LoadConfig()

	http.HandleFunc("/", start)
	http.HandleFunc("/commands", commands)
	http.HandleFunc("/headers", headers)

	fmt.Println("Starting server on port", serverPort, "...")
	http.ListenAndServe(":"+strconv.Itoa(serverPort), nil)
}

func commands(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Command fired")

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	// Work / inspect body. You may even modify it!

	// And now set a new body, which will simulate the same data we read:
	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	w.WriteHeader(http.StatusAccepted)
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func start(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	aCommand, err := slack.SlashCommandParse(req)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch aCommand.Command {
	case "/firebot":
		params := &slack.Msg{Text: aCommand.Text}
		response := fmt.Sprintf("You asked to fire %v", params.Text)
		w.Write([]byte(response))

	default:
		fmt.Println("Default command")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
