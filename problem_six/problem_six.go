

package main

import (
  "strings"
  "fmt"
  "log"
  "bufio"
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
  res := move(scanner, 14)
  fmt.Printf("Result for part one: %d \n", res)
}


func move(scanner *bufio.Scanner, amount int) int {
  ans := 0
  for scanner.Scan() {
    text := scanner.Text()
    ans = firstMarker(text, amount)
  }
  return ans
}

func firstMarker(input string, amount int) int {  
  left, right := 0, 0
  seen := "" 
  log.Printf("length: %d", len(input)) 
  for right < len(input) && len(seen) < amount {
    if strings.Contains(seen, string(input[right])) {
      for input[left] != input[right] {
        seen = strings.Replace(seen, string(input[left]), "", 1)
        left += 1
      }
      seen = strings.Replace(seen, string(input[left]), "", 1)
      left +=1
      
    }     
    seen += string(input[right])
    right +=1
  }
  log.Printf("seen: %s", seen)
  return right
}

