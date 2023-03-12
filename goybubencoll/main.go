package main

import (
	"bufio"
	"fmt"
	"os"

	goybuben "github.com/ohmin839/goybuben/goybubenapi"
)

func main() {
	var wordSet []string
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		words := goybuben.ToHayerenWords(input.Text())
	OUTER:
		for _, w1 := range words {
			for _, w2 := range wordSet {
				if w1 == w2 {
					continue OUTER
				}
			}
			wordSet = append(wordSet, w1)
		}
	}
	for _, w := range wordSet {
		fmt.Println(w)
	}
}
