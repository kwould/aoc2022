package main

import (
  "strings"
  "fmt"
  "log"
  "bufio"
  // "math"
  // "strconv"
  "os"
)

func main() {
  inputFile := os.Args[1]
  log.Printf("Opening file: %s", inputFile)
  file, err := os.Open(inputFile)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  scanner := bufio.NewScanner(file)
  // res := partOne(scanner)
  // fmt.Printf("Result for part one: %d \n", int(res))
  part2res := partTwo(scanner)
  fmt.Printf("Result for part two: %v \n", part2res)
}


func getRoundDecisionPoints(oponnentChoice, myChoice string) int {
  if oponnentChoice == myChoice {
    return 3
  } else if myChoice == getWinningValue(oponnentChoice) {
    return 6
  } else {
    return 0
  }
}

func getLosingValue(move string) string {
  if move == "ROCK" {
    return "SCISSORS"
  } else if move == "SCISSORS" {
    return "PAPER"
  } else {
    return "ROCK"
  }
}

func getWinningValue(move string) string {
  if move == "ROCK" {
    return "PAPER"
  } else if move == "SCISSORS" {
    return "ROCK"
  } else {
    return "SCISSORS"
  }
}
func getRoundDecisionPointsTwo(oponnentChoice, myChoice string) int {
  score := 0
  if myChoice == "X" {
    score+= getMoveValue(getLosingValue(oponnentChoice))
  } else if myChoice == "Y" {
    score+= 3 + getMoveValue(oponnentChoice)
  } else {
    score+= 6 + getMoveValue(getWinningValue(oponnentChoice))
  }
  return score

}

func getMoveValue(m string) int {
  if m == "ROCK" {
    return 1
  } else if m == "PAPER" {
    return 2
  } else {
    return 3
  }
}

func getRPSFromLetter(choice string) string {
  if choice == "A" || choice == "X" {
    return "ROCK"
  } else if choice == "B" || choice == "Y" {
    return "PAPER"
  } else {
    return "SCISSORS"
  }
}

func getScorePartTwo(oponnentChoice, myChoice string) int {
  oponnentRPS := getRPSFromLetter(oponnentChoice)
  return getRoundDecisionPointsTwo(oponnentRPS, myChoice)
}

func getScorePartOne(oponnentChoice, myChoice string) int {
  roundScore := 0
  oponnentRPS := getRPSFromLetter(oponnentChoice)
  myRPS := getRPSFromLetter(myChoice)
  roundScore += getMoveValue(myRPS)
  roundScore += getRoundDecisionPoints(oponnentRPS, myRPS)
  return roundScore
}
func partOne(scanner *bufio.Scanner) int {
  score:= 0
  for scanner.Scan() {
    round := scanner.Text()
    choices := strings.Split(round, " ")
    score += getScorePartOne(choices[0], choices[1])  
  }
  return score
}
func partTwo(scanner *bufio.Scanner) int {
  score:= 0
  for scanner.Scan() {
    round := scanner.Text()
    choices := strings.Split(round, " ")
    score += getScorePartTwo(choices[0], choices[1])  
  }
  return score
}
