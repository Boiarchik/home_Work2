package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main2() {
	fmt.Print("Введите строку: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println(len([]rune(strings.TrimSpace(input))))
}
