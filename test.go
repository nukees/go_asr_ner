package main

import (
	"fmt"
	"time"
)

func main() {
	// Работа таймера
	t1 := time.Now()
	time.Sleep(time.Second * 2)
	t2 := time.Now()
	t := t2.Sub(t1)

	// Форматирование даты и времени для вывода
	// В шаблоне формата обязателько указывается время Mon 02-01-2006 (2 января 2016) 15:04:05
	// Приколы го?
	fmt.Println("Time now: ", t1.Format("02-01-2006"))

	// Преобразовнаие времени в строковую переменную
	fmt.Println("Time ", t.String())

	// Задание времени - обязательны все поля
	t3 := time.Date(2019, time.June, 15, 0, 0, 0, 0, time.UTC)
	t4 := time.Date(2019, time.July, 15, 0, 0, 0, 0, time.UTC)
	t = t4.Sub(t3)
	fmt.Println(t.String())

	var s string
	s = t1.Format("02-Jan-2006")
	fmt.Printf("Type - %T\n", s)

	fmt.Println("--- Work complete ---")
}
