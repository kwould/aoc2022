
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
  // res := move(scanner, true)
  // fmt.Printf("Result for part one: %s \n", res)
  part2res := move(scanner, false)
  fmt.Printf("Result for part two: %v \n", part2res)
}

// {T} {V}                     {W}    
// {V} {C} {P} {D}             {B}    
// {J} {P} {R} {N} {B}         {Z}    
// {W} {Q} {D} {M} {T}     {L} {T}    
// {N} {J} {H} {B} {P} {T} {P} {L}    
// {R} {D} {F} {P} {R} {P} {R} {S} {G}
// {M} {W} {J} {R} {V} {B} {J} {C} {S}
// {S} {B} {B} {F} {H} {C} {B} {N} {L}
//  1   2   3   4   5   6   7   8   9 

func move(scanner *bufio.Scanner, partOne bool) string {
  grid := map[int][]string { 1: {"S", "M", "R", "N", "W", "J", "V", "T"},
                            2: {"B", "W", "D", "J", "Q", "P", "C", "V"}, 
                            3: {"B", "J", "F", "H", "D", "R", "P"}, 
                            4: {"F", "R", "P", "B", "M", "N", "D"}, 
                            5: {"H", "V", "R", "P", "T", "B"}, 
                            6: {"C", "B", "P", "T"}, 
                            7: {"B", "J", "R", "P", "L"}, 
                            8: {"N", "C", "S", "L", "T", "Z", "B", "W"}, 
                            9: {"L", "S", "G"}, }
  for scanner.Scan() {
    text := scanner.Text()
    handleMove(text, grid, partOne)  
  }
  ans := ""
  for i:=1; i<= 9; i++ {
    if len(grid[i]) > 1{
      ans += grid[i][len(grid[i]) - 1]

    }
  }
  return ans
}

func movePartTwo(amount int, source int, dest int, grid map[int][]string) {
    crates := grid[source][len(grid[source]) - amount:]
    grid[source] = grid[source][:len(grid[source]) - amount]
    grid[dest] = append(grid[dest], crates...)
}
func movePartOne(amount int, source int, dest int, grid map[int][]string) {
  for i := 1; i <= amount; i++ {
    crate := grid[source][len(grid[source]) -1 ]
    grid[source] = grid[source][:len(grid[source])-1]
    grid[dest] = append(grid[dest], crate)
  }
}
func handleMove(move string, grid map[int][]string, partOne bool) {
  words := strings.Split(move, " ")
  amount, _ := strconv.Atoi(words[1])
  source, _ := strconv.Atoi(words[3])
  dest, _ := strconv.Atoi(words[5])
  if partOne {
    movePartOne(amount, source, dest, grid)
  } else {
    movePartTwo(amount, source, dest, grid)
  }
}
