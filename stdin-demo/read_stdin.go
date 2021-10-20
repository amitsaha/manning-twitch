package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func capitalize(r io.Reader) (string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	inputS := scanner.Text()
	return strings.ToUpper(inputS), nil
}

func main() {
	result, err := capitalize(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
