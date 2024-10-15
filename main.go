package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	totalPoints       int = 100
	pointsPerQuestion int = 20
)

func main() {
	fmt.Println("Вітаємо у грі!")
	time.Sleep(3 * time.Second)

	for i := 5; i >= 1; i-- {
		fmt.Printf("Гра почнеться через: %v\n", i)
		time.Sleep(1 * time.Second)
	}

	myPoints := 0
	for myPoints < totalPoints {
		x, y := rand.Intn(100), rand.Intn(100)
		fmt.Printf("%v + %v = ", x, y)
		ans := ""
		fmt.Scan(&ans)

		ansInt, err := strconv.Atoi(ans)
		if err != nil {
			fmt.Println("Не правильна відповідь! Тільки числові значення!")
		} else {

			if ansInt == x+y {
				myPoints = myPoints + pointsPerQuestion
				fmt.Println("Правильно! Набрано балів:", myPoints)
				time.Sleep(2 * time.Second)

			} else {
				fmt.Println("Не правильно! Набрано балів:", myPoints)
			}

		}
		//fmt.Println(" Все розв'язано правильно! Кінець гри!")
		//time.Sleep(3 * time.Second)
	}

}
