package main

import "fmt"

func FichDich() {
	letters := []string{"b", "a", "c", "a", "b", "a",
		"a", "b", "c"}
	m := make(map[string]int)
	for _, v := range letters {
		if _, ok := m[v]; !ok {
			m[v] = 1
		} else {
			m[v] += 1
		}
	}
	for i, v := range m {
		fmt.Printf("Key: %s, Value: %d\n", i, v)
	}
}
