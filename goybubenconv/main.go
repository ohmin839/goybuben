package main

import (
	"bufio"
	"fmt"
	"os"

	goybuben "github.com/ohmin839/goybuben/goybubenapi"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(goybuben.ToAybuben(input.Text()))
	}
}
