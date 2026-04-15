package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func capitalizeWords(s string) string {
	runes := []rune(s)
	inWord := false
	for i, r := range runes {
		if unicode.IsSpace(r) {
			inWord = false
			continue
		}
		if !inWord {
			runes[i] = unicode.ToUpper(r)
			inWord = true
		} else {
			runes[i] = unicode.ToLower(r)
		}
	}
	return string(runes)
}

func main() {
	fmt.Print("Введите строку: ")
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	input = strings.TrimRight(input, "\r\n")
	fmt.Println(capitalizeWords(input))
}
