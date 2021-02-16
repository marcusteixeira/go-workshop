package main

import "fmt"

func main() {
	words := map[string]int{
		"Gonna": 3,
		"You":   3,
		"Give":  2,
		"Never": 1,
		"Up":    4,
	}
	topWord := ""
	topCount := 0

	for k, v := range words {
		if v > topCount {
			topCount = v
			topWord = k
		}
		fmt.Println(k, "=", v)
	}

	fmt.Println("The most popular word: ", topWord)
	fmt.Println("Count Number: ", topCount)
}
