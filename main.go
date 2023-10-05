package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type InspirationalQuote struct {
	Quote  string `json:"q"`
	Author string `json:"a"`
}

func main() {
	resp, err := http.Get("https://zenquotes.io/api/random")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	b, err := io.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	var quotes []InspirationalQuote

	json.Unmarshal(b, &quotes)

	for i := 0; i < len(quotes); i++ {
		fmt.Printf("\n\"%s\" \n- %s\n\n", quotes[i].Quote, quotes[i].Author)
	}
}
