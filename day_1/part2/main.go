package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main()  {
  list1, list2 := ReadInput()

  var result int
  for _, value := range list1 {
    var times int
    for _, v := range list2 {
      if v == value {
        times = times + 1
      }
    }

    result = result + (value * times)
  }

  fmt.Println(result)

}

func ReadInput() ([]int, []int) {
  f, err := os.Open("./day_1/part2/input.txt")
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

  return list1, list2
}
