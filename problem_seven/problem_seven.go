package main

import (
  "strings"
  "fmt"
  "log"
  "strconv"
  "bufio"
  "math"
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
  res := move(scanner, false)
  fmt.Printf("Result for part two: %d \n", res)
}


func get_values_bigger_than_n(dirs map[string]int) int {
  ans := 0
  for _, val := range dirs {
    if val <  100000 {
      ans += val
    }
  }
  return ans
}
func get_values_min_possible_value(dirs map[string]int) int {
  ans := 100000000000000000.0
  for _, val := range dirs {
    if val >= (dirs["/]/"] - 30000000) {
      ans = math.Min(ans, float64(val))
    }
  }
  return int(ans)
}

func move(scanner *bufio.Scanner, partOne bool) int {
  dirs := map[string]int {}
  curDir := [] string{}
  for scanner.Scan() {
    text := scanner.Text()
    // log.Printf("text: %s", text)
    dirs, curDir = handleCommand(dirs, curDir,  text) 
  }
  ans := 0
  if partOne {
     ans  = get_values_bigger_than_n(dirs)
  } else {
    ans = get_values_min_possible_value(dirs)    
  }

  return ans
}

func handleCommand(dirs map[string]int, curDir []string, text string) (map[string]int, []string) {  
  words := strings.Split(text, " ")
  if words[0] == "$" {
    if words[1] != "ls" {
      curDir = handleCD(words, curDir)

    }
  } else {
    dirs = handleLs(words, curDir, dirs)
  }
  return dirs, curDir
}

func handleCD(words []string, curDir []string)  []string {
  if words[2] == ".." {
    curDir = curDir[:len(curDir) -1]
  } else {
    curDir = append(curDir, words[2])
  }
  return curDir
}

func handleLs(words []string, curDirs []string, dirs map[string]int) map[string]int{
  amount, err := strconv.Atoi(words[0])

  if err != nil {
    // this is a dir command
    // curDirs = append(curDirs, words[1])
  } else {
    dirStr := ""
    for _, dir := range curDirs {
      dirStr +="/" + dir
      if  _, ok := dirs[dirStr]; ok {
        dirs[dirStr]  += amount

      } else {
        dirs[dirStr] = amount
      }
    }
  
  }
  return dirs
  

}

