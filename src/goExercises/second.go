package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func secondEx() {
	reader := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(100)
	fmt.Println(randNum)

	for i := 0; i < 10; i++ {
		fmt.Print("Enter a number: ")
		input, err := reader.ReadString('\n')
		checkErr(err)
		input = strings.TrimSpace(input)
		num, err := strconv.ParseFloat(input, 64)
		checkErr(err)
		if num == float64(randNum) {
			fmt.Println("U won!!!")
			return
		} else if num > float64(randNum) {
			fmt.Println("Too high")
		} else if num < float64(randNum) {
			fmt.Println("Too low")
		}
	}

	fmt.Println("U wrong number was =", randNum)
}
