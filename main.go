package main

import "fmt"

func main() {
	text1 := "Get ready"

	score := 0
	
	fmt.Println(text1)
	fmt.Println("Ваш счёт:", score)

	fmt.Println("Вы пролетели через трубу")
	score = score + 1
	fmt.Println("Ваш счёт:", score)

	fmt.Println("Вы пролетели через трубу")
	score = score + 1
	fmt.Println("Ваш счёт:", score)

	text2 := "вы не смогли пролететь через трубу"
	score = 0
	fmt.Println(text2)
	fmt.Println("Ваш счёт:", score)

}
