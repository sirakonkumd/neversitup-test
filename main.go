package main

import (
	"fmt"
	"neversitup-test/services/oddService"
	"neversitup-test/services/permutationService"
	"neversitup-test/services/smileService"
)

func main() {
	input := []int{7}
	input2 := []int{0}
	input3 := []int{1, 1, 2}
	input4 := []int{0, 1, 0, 1, 0}
	input5 := []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}
	input6 := []int{}

	fmt.Println("==> result FindOdd : ", oddService.FindOdd(input))
	fmt.Println("==> result FindOdd : ", oddService.FindOdd(input2))
	fmt.Println("==> result FindOdd : ", oddService.FindOdd(input3))
	fmt.Println("==> result FindOdd : ", oddService.FindOdd(input4))
	fmt.Println("==> result FindOdd : ", oddService.FindOdd(input5))
	fmt.Println("==> result FindOdd : ", oddService.FindOdd(input6))

	inputSm := []string{":)", ";(", ";}", ":D", ":-()", "(:-)", ";~)"}
	fmt.Println("==> result find smile : ", smileService.FindSmile(inputSm))
	// fmt.Println("==> result count smile : ", service.CountSmileys(inputSm))

	inputPer := "a"
	inputPer1 := "ab"
	inputPer2 := "abc"
	inputPer3 := "abcc"
	fmt.Println("==> result permute : ", (permutationService.Permutations(inputPer)))
	fmt.Println("==> result permute : ", (permutationService.Permutations(inputPer1)))
	fmt.Println("==> result permute : ", (permutationService.Permutations(inputPer2)))
	fmt.Println("==> result permute : ", (permutationService.Permutations(inputPer3)))

}
