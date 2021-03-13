package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type gameData struct {
	difficulty    int
	corrects      int
	totalAnswered int
}

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
func playQuiz() {
	filepath := "QuizGame/problems.csv"

	fmt.Println("Welcome to a simple math quiz game where you will \n" +
					"have 30 seconds to answer as many questions as possible!")
	fmt.Println("\nLet's start off by you giving a difficulty level to the quiz.")
	fmt.Println("Please start by typing a number between 1-100")

	correctInput := false
	questions := 50

	// Failsafe
	for correctInput == false {
		var difficulty int
		fmt.Scanln(&difficulty)

		if difficulty <= 0 || difficulty > 100 {
			fmt.Println("Please read the instructions!")
		} else {
			createQuiz(filepath, questions, difficulty)
			correctInput = true
		}
	}

	// Start line
	quiz := readQuiz(filepath)
	fmt.Println("Press enter to start!")
	fmt.Scanln()
	var timeOut bool

	// Timer
	timer := time.AfterFunc(30*time.Second, func() {
		fmt.Println("30 seconds has passed!")
		fmt.Println("You may answer the last question.")
		timeOut = true
	})
	defer timer.Stop()

	var correct int
	var answered int
	var input int

	// Game loop
	for i := 0; i < questions && timeOut == false; i++ {
		fmt.Println(quiz[i][0])
		fmt.Scanln(&input)

		if strconv.Itoa(input) == quiz[i][1] {
			correct++
			answered++
		} else {
			answered++
		}
	}

	// Results
	fmt.Println("Correct: ", correct, "/", answered)
}

// MAIN
func main() {
	playQuiz()
}