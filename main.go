package main

import (
	"bufio"
	"fmt"
	"os"

	"./bencoding"
)

func main() {
	// Create a bufio.Reader from os.Stdin
	reader := bufio.NewReader(os.Stdin)

	// Call the Decode function from the bencoding package
	result, err := bencoding.Decode(reader)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the result
	fmt.Println("Result:", result)
}
