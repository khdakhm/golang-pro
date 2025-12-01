package main

import (
	"fmt"
	"time" // Добавили эту строку!
)

func main() {
	fmt.Println("проверка!")
	fmt.Println("Сейчас время:", time.Now()) // Теперь time не подсвечивается красным
}
