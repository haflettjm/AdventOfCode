package main

import (
  "bufio"
  "regexp"
  "fmt"
  "strconv"
  "log"
  "os"
)

func getFile()[]string{
  array := []string{}
  file, err := os.Open("../input/input.txt")
  if err != nil{
    log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan(){
    array = append(array, scanner.Text())
  }

  if err := scanner.Err(); err != nil{
    log.Fatal(err)
  }
  return array
}

func convert(array []string) string{
  newStr := array[0] + array[1]
  return newStr
}

func getSum(array []string) int{
  sum := 0

  for _, v := range array{
    c, _ := strconv.Atoi(v)
    sum += c
  }

  return sum
}

func main(){
  input := getFile()
  var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
  numbers := []string{}
  for _, w := range input{
    num := []string{}
    for _, c := range w{
      letter := IsLetter(string(c)) 
      if !letter{
        if len(num) == 0{
          num = append(num, string(c))
        }else{
          if len(num) < 2{
            num = append(num, string(c))
          }else{
            num[1] = string(c) 
          }
        }
      }
    }
    if len(num) < 2{
      num = append(num, num[0])
    }
    newNum := convert(num)
    numbers = append(numbers, newNum)
  }

  total := getSum(numbers)
  fmt.Printf("Done! The total is: \n %d \n", total)
}
