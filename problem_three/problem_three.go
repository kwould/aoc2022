package main

import (
  "strings"
  "fmt"
  "log"
  "bufio"
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

func getPartOneAnswer(text string) string {
  wholeValue := strings.Split(text, "")   
  firstCompartment := wholeValue[:len(wholeValue)/2]
  secondCompartment := wholeValue[len(wholeValue)/2:] 
  firstMap := make(map[string]bool)
  for _, item := range firstCompartment {
    firstMap[item] = true
  }
  for _, item := range secondCompartment {
    if _, ok := firstMap[item]; ok {
      return item
    }
  }
  return ""
}
func getUniqueKeys(keys string) map[string]bool {
  var set = map[string]bool {}
  for _, key := range keys {
    set[string(key)] = true 
  }
  return set
}
func Intersect(setOne map[string]bool, setTwo map[string]bool) map[string]bool {
  result := map[string]bool {}
  for key, _ := range setOne {
    if _, ok := setTwo[key]; ok {
      result[key] = true
    }
  }
  return result
}

func getPartTwoAnswer(groups []string) string  {
  seenSet := getUniqueKeys(groups[0])
  for _, groupItem := range groups[1:] {
    keysForGroup := getUniqueKeys(groupItem)
    seenSet = Intersect(keysForGroup, seenSet)
  }
  for key, _ := range seenSet {
    return key 
  }
  return ""
}

func getPriority(itemForPriority string) int {
  val := int(rune([]byte(itemForPriority)[0]))
  if val >= 97 {
    val = val - 96
  } else {
    // values start at 65
    val = val -38
  }

  return val 
}

func partTwo(scanner *bufio.Scanner) int {
  score:= 0
  var group []string = nil
  for scanner.Scan() {
    text := scanner.Text()
    group = append(group, text)
    if len(group) == 3 {
      commonValue := getPartTwoAnswer(group)
      score += getPriority(commonValue)
      group = group[:0]
    }

  }
  return score
}

func partOne(scanner *bufio.Scanner) int {
  score:= 0
  for scanner.Scan() {
    text := scanner.Text()
    commonValue := getPartOneAnswer(text)
    score += getPriority(commonValue)
  }
  return score
}
