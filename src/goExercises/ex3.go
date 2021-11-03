package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type mySlice []float64

func (a *mySlice) Add(n float64) {
	*a = append(*a, n)
}

func thirdEx() {
	fmt.Println("How many students do u need?: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	sliceOfNumbers := mySlice{}
	var sliceOfNames []string
	sliceOfAvg := mySlice{}
	iterations, _ := strconv.ParseFloat(input, 64)
	for i := 0; i < int(iterations); i++ {
		fmt.Println("Pls write the name of the student")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)
		sliceOfNames = append(sliceOfNames, name)
		fmt.Println("His first grade:")
		num1 := ReadTrimStrconv()
		sliceOfNumbers.Add(num1)
		fmt.Println("His second grade:")
		num2 := ReadTrimStrconv()
		sliceOfNumbers.Add(num2)
		fmt.Println("His third grade")
		num3 := ReadTrimStrconv()
		sliceOfNumbers.Add(num3)
		sliceOfAvg.Add((num1 + num2 + num3) / 3)
	}
	fmt.Printf("Names | Score1 | Score2 | Score3 | Average\n")
	someNum := 0
	for i := 0; i < int(iterations); i++ {
		fmt.Printf("%10s %8.2f %8.2f %8.2f %8.2f \n",
			sliceOfNames[i], sliceOfNumbers[someNum],
			sliceOfNumbers[someNum+1],
			sliceOfNumbers[someNum+2],
			sliceOfAvg[i])
		someNum += 3
	}
}

func ReadTrimStrconv() float64 {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	num, _ := strconv.ParseFloat(input, 64)
	return num
}
