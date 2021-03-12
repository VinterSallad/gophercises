package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

/**
Takes file to clear and then create a simple math quiz based on user inputted, amount of questions and difficulty.
 */
func createQuiz(filepath string, questions int, difficulty int) {
	// Setting seed for RNG
	rand.Seed(time.Now().UnixNano())

	file, err := os.OpenFile(filepath, os.O_RDWR, 0)
	if err != nil {
		fmt.Println("File not found", err)
	}

	// Emptying file
	file.Truncate(0)

	for index := 0; index < questions; index++ {
		operand := getOperator()
		i := rand.Intn(difficulty)
		j := rand.Intn(difficulty)
		iString := strconv.Itoa(i)
		jString := strconv.Itoa(j)

		question := iString + operand + jString
		answer := calculateAnswer(i, j, operand)

		file.WriteString(question + "," + answer + "\n")
	}

	// Save file
	file.Sync()
	file.Close()
}

/**
Returns a string representation of an operand
 */
func getOperator() string {
	n := rand.Intn(3)

	switch n {
	case 0:
		return "+"
	case 1:
		return "-"
	case 2:
		return "*"
	}

	return ""
}

/**
Calculates two integers with specified string operand and returns answer as string.
 */
func calculateAnswer(i int, j int, operand string) string {
	switch operand {
	case "+":
		return strconv.Itoa(i+j)
	case "-":
		return strconv.Itoa(i-j)
	case "*":
		return strconv.Itoa(i*j)
	}

	return ""
}

/**
Reads the specified csv file and returns an array representation of csv file.
 */
func readQuiz(filepath string) [][]string {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("File not found", err)
	}

	reader := csv.NewReader(file)
	quiz, _ := reader.ReadAll()

	return quiz
}

/**
Initiates game loop from string array representation of games csv file.
 */
func playQuiz(quiz [][]string) {

}

func main() {
	filepath := "QuizGame/problems.csv"
	createQuiz(filepath, 10, 10)
	fmt.Println(readQuiz(filepath))
}