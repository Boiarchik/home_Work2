package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Введите строку: ")
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	input = strings.TrimRight(input, "\r\n")
	input = strings.ToLower(input)

	count := 0
	for _, r := range input {
		switch r {
		case 'а', 'е', 'ё', 'и', 'о', 'у', 'ы', 'э', 'ю', 'я':
			count++
		}
	}
	fmt.Println(count)
}
