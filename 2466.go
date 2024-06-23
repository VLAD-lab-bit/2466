package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// isPrime проверяет, является ли число простым
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Введите целое число: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		number, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Неверный ввод. Пожалуйста, введите целое число.")
			continue
		}

		if isPrime(number) {
			fmt.Printf("Число %d является простым\n", number)
		} else {
			fmt.Printf("Число %d не является простым\n", number)
		}
		break
	}
}
