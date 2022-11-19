package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func repetition(st string) map[string]int {
	// using strings.Field Function
	input := strings.Fields(st)
	wc := make(map[string]int)
	for _, word := range input {
		_, matched := wc[word]
		if matched {
			wc[word] += 1
		} else {
			wc[word] = 1
		}
	}
	fmt.Println(wc)
	return wc
}

func countRepWords(str map[string]int) int {
	count := 0
	for _, word := range str {
		if word > 1 {
			count += 1
		}
	}
	return count
}

func main() {
	content, err := os.ReadFile(`repitations`)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//fmt.Println(string(content))

	re, err := regexp.Compile(`.*(\.|(:|\.*))`)
	res := re.FindAllString(string(content), -1)
	fmt.Println(len(res))

	a := []int{}

	for _, sentenses := range res {
		s := repetition(sentenses)
		count := countRepWords(s)
		if count > 0 {
			a = append(a, count)
		}
	}
	fmt.Println(a)

	min := a[0]
	max := a[0]

	for i := 0; i < len(a); {
		if a[i] < min {
			min = a[i]
		}
		if a[i] > max {
			max = a[i]
		}
		i++
	}
	fmt.Println(`min elementof array: `, min)
	fmt.Println(`max elementof array: `, max)
	result := (max + min) / 2
	fmt.Println(result)
}
