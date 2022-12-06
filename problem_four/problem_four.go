package main

import (
  "strings"
  "fmt"
  "log"
  "bufio"
  "strconv"
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
  // res := PartOne(scanner)
  // fmt.Printf("Result for part one: %d \n", int(res))
  part2res := PartTwo(scanner)
  fmt.Printf("Result for part two: %v \n", part2res)
}

func parseRanges(group string) []int {
  groupRange := strings.Split(group, "-") 
  intRange := []int {0, 0}
  intRange[0], _ = strconv.Atoi(groupRange[0])
  intRange[1], _ = strconv.Atoi(groupRange[1])
  return intRange
}

func PartialOverlaps(groupOne string, groupTwo string) bool {
  oneVal := parseRanges(groupOne)
  twoVal := parseRanges(groupTwo) 
  // 2-8 3-7
  if (oneVal[0] <= twoVal[0] && oneVal[1] >= twoVal[0]) || (twoVal[0] <= oneVal[0] && twoVal[1] >= oneVal[0]) {
    return true
  }
  return false

}
func FullyOverlaps(groupOne string, groupTwo string) bool {
  oneVal := parseRanges(groupOne)
  twoVal := parseRanges(groupTwo) 
  // 2-8 3-7
  if (oneVal[0] <= twoVal[0] && twoVal[1] <= oneVal[1]) || (twoVal[0] <= oneVal[0] && oneVal[1] <= twoVal[1]) {
    return true
  }
  return false

}
func PartTwo(scanner *bufio.Scanner) int {
  score := 0
  for scanner.Scan() {
    text := scanner.Text()
    groups := strings.Split(text, ",")
    if PartialOverlaps(groups[0], groups[1]) {
      score +=1
    }
  }
  return score

}
func PartOne(scanner *bufio.Scanner) int {
  score := 0
  for scanner.Scan() {
    text := scanner.Text()
    groups := strings.Split(text, ",")
    if FullyOverlaps(groups[0], groups[1]) {
      score +=1
    }
  }
  return score

}
