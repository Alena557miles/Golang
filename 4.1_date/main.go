package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	// 1 read
	content, err := os.ReadFile("dates")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// fmt.Print(string(content))
	re, _ := regexp.Compile(`[A-Za-z]+ [0-9]{1,2}(th)?(\, [0-9]{4})?`)
	res := re.FindAllString(string(content), -1)
	fmt.Println(res)
}
