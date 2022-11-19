package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func createFileandWriteData(filename, data string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	_, err = file.WriteString(data)
}

func openAndPrintFileData(s string) {
	contentLinks, err := os.ReadFile(s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(contentLinks))
}

func main() {
	resultFile := "links"
	arrayOfLinks := []string{}

	files, err := os.ReadDir("files")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		// fmt.Println(file.Name())
		matched, err := regexp.MatchString(`_`, file.Name())
		if err != nil {
			fmt.Println(err)
		}
		if matched {
			os.Chdir("files")
			fileContent, err := os.ReadFile(file.Name())
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			// fmt.Println(string(fileContent))
			re, err := regexp.Compile(`(https://)?\w+\.\w+((/\w+)|(\.\w+))*\s`)
			res := re.FindAllString(string(fileContent), -1)
			// fmt.Println(res)
			if len(res) >= 3 {
				data := strings.Join(res, "")
				arrayOfLinks = append(arrayOfLinks, data)
				// fmt.Println(data)
			}
		}
	}
	data := strings.Join(arrayOfLinks, "")
	createFileandWriteData(resultFile, data)
	openAndPrintFileData(resultFile)
}
