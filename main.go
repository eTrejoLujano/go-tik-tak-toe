package main

import "fmt"

func main() {
	var pick string
	var tally int
	leftTop := "LT"
	middleTop := "MT"
	rightTop := "RT"
	leftMiddle := "LC"
	middleMiddle := "MC"
	rightMiddle := "RC"
	leftBottom := "LB"
	middleBottom := "MB"
	rightBottom := "RB"
	fmt.Printf(" %s | %s | %s \n", leftTop, middleTop, rightTop)
	fmt.Println("____|____|____")
	fmt.Printf(" %s | %s | %s \n", leftMiddle, middleMiddle, rightMiddle)
	fmt.Println("____|____|____")
	fmt.Println("    |    |    ")
	fmt.Printf(" %s | %s | %s \n", leftBottom, middleBottom, rightBottom)
	fmt.Println("Please specify which square you'd like to take next.")
	fmt.Scanln(&pick)
	tally += 1
	fmt.Println(tally)
}
