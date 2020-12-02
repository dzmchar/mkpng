package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func processFile(filename *string) {
	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal("file not found")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		width, height := processLine(scanner.Text())
		flow(&width, &height)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func processLine(text string) (widthValue int, heightValue int) {
	var width string
	var height string
	var hasSwitched bool
	for _, char := range text {
		if hasSwitched == false && char != 'x' {
			width += string(char)
			continue
		}
		if char == 'x' {
			hasSwitched = true
			continue
		}
		if hasSwitched {
			height += string(char)
		}
	}
	widthValue, err := strconv.Atoi(width)
	if err != nil {
		log.Fatal("error while converting width")
	}
	heightValue, err = strconv.Atoi(height)
	if err != nil {
		log.Fatal("error while converting height")
	}
	return
}
