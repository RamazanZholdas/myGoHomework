package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func firstEx() {
	reader := bufio.NewReader(os.Stdin)
	var res float64

	for {
		fmt.Print("Enter a number: ")
		input, err := reader.ReadString('\n')
		checkErr(err)
		input = strings.TrimSpace(input)
		num, err := strconv.ParseFloat(input, 64)
		checkErr(err)
		fmt.Println("Wanna add some more?[y/n]")
		inputStr, err := reader.ReadString('\n')
		checkErr(err)
		inputStr = strings.TrimSpace(inputStr)
		res += num
		if inputStr == "n" {
			break
		}
	}

	fmt.Println(res)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	} else {
		return
	}
}
