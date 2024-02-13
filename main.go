package main

import (
	"fmt"
	"os"
)

func exit() {
	fmt.Println("")
	fmt.Println("Program execution complete")
	os.Exit(0)
}

func main() {
	fmt.Println(":: Running BubbleTutorial")
	RunBubbleTutorial()

	exit()
	fmt.Println(":: Running SchollzBar")
	RunSchollz()

	fmt.Println(":: Running BubbleTimer")
	RunBubbleTimer()
	fmt.Printf("\n\n\n")

	fmt.Println(":: Running BubblePureProgress")
	RunBubblePureProgress()
	fmt.Printf("\n\n\n")

	fmt.Println(":: Running BubbleDynamic")
	RunBubbleDynamic()
}
