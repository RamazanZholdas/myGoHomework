package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func eightEx() {
	file, err := os.Open("votes.txt")
	CheckErr(err)
	scan := bufio.NewScanner(file)
	slice := []string{}
	for scan.Scan() {
		slice = append(slice, scan.Text())
	}
	slice2 := []string{}
	for i := 0; i < len(slice); i++ {
		str := swapFirstAndLast(strings.Trim(slice[i], " "))
		slice2 = append(slice2, str)
	}
	sort.Strings(slice2)
	for i, v := range slice2 {
		fmt.Println(i, v)
	}
}

func swapFirstAndLast(str string) string {
	slice := strings.Split(str, ",")
	return slice[1] + " " + slice[0]
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
