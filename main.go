package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/adityjoshi/GoCo/bencoding"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	result, err := bencoding.Decode(reader)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result:", result)
}
