package main

import (
  "container/heap"
  "fmt"
  "log"
  "bufio"
  "math"
  "strconv"
  "os"
)


type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
  inputFile := os.Args[1]
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

func partOne(scanner *bufio.Scanner) int {
  res := 0.0
  curMax := 0
  for scanner.Scan() {
    lineVal := scanner.Text()
    if lineVal == "\n" || lineVal == ""{
      res = math.Max(float64(curMax), res)
      curMax = 0
    } else {
      calories, err:= strconv.Atoi(lineVal)
      if err != nil {
        log.Fatal(err)
      }
      curMax += calories
    }
 
  }
  return int(res)

}

func partTwo(scanner *bufio.Scanner) int{
  res := &IntHeap{}
  heap.Init(res)
  elfsCals := 0
  for scanner.Scan() {
    lineVal := scanner.Text()
    if lineVal == "\n" || lineVal == ""{
      if res.Len() < 3{
        res.Push(elfsCals)
      } else if elfsCals > (*res)[0] {
        heap.Push(res, elfsCals)
        heap.Pop(res)
      }
      elfsCals = 0
    } else {
      calories, err:= strconv.Atoi(lineVal)
      if err != nil {
        log.Fatal(err)
      }
      elfsCals += calories
    }
 
  }
  totalCals := 0
  for res.Len() > 0{
    cals, _ := heap.Pop(res).(int)
    totalCals += cals 
  }
  return totalCals
}
