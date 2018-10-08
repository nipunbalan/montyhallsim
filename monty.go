package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	fmt.Println("first_choice,reveal,second_choice,no_switch,switch")
	var numIterations int64
	var err error
	if len(os.Args) > 1 {
		numIterations, err = strconv.ParseInt(os.Args[1], 10, 64)
		if numIterations > 1000000 {
			fmt.Println("Enter a number less than 1000000")
			return
		}
		failOnError(err, "Invalid Iteration number")
	} else {
		numIterations = 1000
	}

	noSwitchStrategyCount := int64(0)
	switchStrategyCount := int64(0)

	for i := int64(0); i < numIterations; i++ {
		problem := generateProblem()
		//	fmt.Println(problem)
		firstChoice := firstChoice()
		//	fmt.Printf("First Choice: %d\n", firstChoice)
		reveal := revealOne(problem, firstChoice)
		//	fmt.Printf("Reveal: %d\n", reveal)

		secondChoice := executeStrategySwitch(firstChoice, reveal)
		//	fmt.Printf("Second Choice: %d\n", secondChoice)
		evaluationF := evaluateChoice(problem, firstChoice)

		evaluationS := evaluateChoice(problem, secondChoice)
		//	fmt.Printf("Evaluation: %d\n", evaluation)

		if evaluationF == 1 {
			noSwitchStrategyCount++
		} else if evaluationS == 1 {
			switchStrategyCount++
		}
		fmt.Printf("%d,%d,%d,%d,%d\n", firstChoice, reveal, secondChoice, evaluationF, evaluationS)
	}
	fmt.Printf(",,,=SUM(D2:D%d),=SUM(E2:E%d)\n", numIterations+1, numIterations+1)
	fmt.Println("---------------------------------------------------")
	fmt.Printf("Number of Trials: %d\n", numIterations)
	fmt.Printf("Chances for no switch strategy to win: %d%%\n", noSwitchStrategyCount*100/numIterations)
	fmt.Printf("Chances for switch strategy to win: %d%%\n", switchStrategyCount*100/numIterations)
}

/* generateProblem() generates an array with 3 integers one of which
   will be set to 1 randomly */
func generateProblem() [3]int {
	var problem [3]int
	rand.Seed(time.Now().UnixNano() * 56)
	winIndex := rand.Intn(3)
	problem[winIndex] = 1
	return problem
}

//firstChoice generates a random number between 0 and 2
func firstChoice() int {
	rand.Seed(time.Now().UnixNano() * 4)
	choiceIndex := rand.Intn(3)
	return choiceIndex
}

//revealOne accepts an integer array of 3 and the first choice made and return a number
//other than the winIndex and firstChoice
func revealOne(problem [3]int, firstChoice int) int {
	var reveal int
	rand.Seed(time.Now().UnixNano() * 65)
	reveal = rand.Intn(3)
	for problem[reveal] == 1 || reveal == firstChoice {
		reveal = rand.Intn(3)
	}
	return reveal
}

//executeStrategyNoSwitch returns the same number
func executeStrategyNoSwitch(firstChoice int) int {
	return firstChoice
}

//executeStrategySwitch returns a number other than the first close and the reveal
func executeStrategySwitch(firstChoice int, reveal int) int {
	for i := 0; i <= 2; i++ {
		if i != firstChoice && i != reveal {
			return i
		}
	}
	return -1
}

//returns 1 when the value in the choosen index is 1
func evaluateChoice(problem [3]int, choice int) int {
	return problem[choice]
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
