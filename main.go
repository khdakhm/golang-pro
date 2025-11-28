package main

import (
	"fmt"
	"time" // Добавили эту строку!
)

func main() {
	fmt.Println("Моя первая программа на Go!")
	fmt.Println("Сейчас время:", time.Now()) // Теперь time не подсвечивается красным
}
