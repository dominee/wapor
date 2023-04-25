package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"args", "https://www.citadelo.com", "https://sme.sk", "https://github.com"}
	// Test main function
	main()

}
