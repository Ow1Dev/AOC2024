package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
  f, err := os.Open("./day_1/part1/input.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer f.Close()

  scanner := bufio.NewScanner(f)

  list1, list2 := make([]int, 0), make([]int, 0)

  for scanner.Scan() {
    numbers := strings.Split(scanner.Text(), " ")
    number1, _ := strconv.Atoi(strings.TrimSpace(numbers[0]))
    number2, _ := strconv.Atoi(strings.TrimSpace(numbers[len(numbers)-1]))

    list1 = append(list1, number1)
    list2 = append(list2, number2)
  }
  f.Close()

  slices.Sort(list1) 
  slices.Sort(list2) 

  var result int 

  for index := range list1{
    var s int

    if list2[index] > list1[index] {
      s = list2[index] - list1[index] 
    } else {
      s = list1[index] - list2[index] 
    }

    fmt.Println(s)
    result = result + s
  }

  fmt.Println("---")
  fmt.Println(result)
}
