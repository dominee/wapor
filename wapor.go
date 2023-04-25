package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	wappalyzer "github.com/projectdiscovery/wappalyzergo"
)

func main() {
	var urls []string

	// Check if URLs were provided as arguments
	if len(os.Args) > 1 {
		urls = os.Args[1:]
	} else {
		// Read URLs from stdin
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			urls = append(urls, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "error reading input: %s\n", err)
			os.Exit(1)
		}
	}

	// Create a new Wappalyzer instance
	wappalyzerClient, err := wappalyzer.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating Wappalyzer instance: %s\n", err)
		os.Exit(1)
	}

	// Loop through each URL and print the analysis output
	for _, url := range urls {
		// Analyze the URL with Wappalyzer
		resp, err := http.DefaultClient.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		data, _ := io.ReadAll(resp.Body) // Ignoring error for example

		fingerprints := wappalyzerClient.Fingerprint(resp.Header, data)
		// iterate over the fingerprints and print the name of the technology
		var fingerprintStr string
		//fmt.Println(fingerprints)
		for fingerprint := range fingerprints {
			fingerprintStr = fingerprintStr + fingerprint + "; "
		}
		// remove the last space and semicolon
		fingerprintStr = fingerprintStr[:len(fingerprintStr)-2]
		fmt.Printf("%s : %v\n", url, fingerprintStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error analyzing URL %s: %s\n", url, err)
			continue
		}
	}

}
