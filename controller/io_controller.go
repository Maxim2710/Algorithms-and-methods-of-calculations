package controller

import "fmt"

func InputOutputData() {
	var numLabarotary, numTask int
	fmt.Println()
	fmt.Print("Please, enter the number of the laboratory work: ")
	fmt.Scan(&numLabarotary)

	fmt.Print("Laboratory work number: ")
	fmt.Print(numLabarotary)
	fmt.Print("; Select the task number: ")
	fmt.Scan(&numTask)
	fmt.Println("Laboratory work:", numLabarotary, "| Task:", numTask)
	fmt.Println()

	IODistribution(numLabarotary, numTask)
}
