package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func main() {
	file := openFile("input")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var total int

	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Error("An error occured while converting the line of the file from string to int")
			log.Error(err)
		}
		total += calculateFuel(mass)
	}

	fmt.Println("Total: ", total)
}

func openFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("An error occured when tried to open the file")
	}
	return file
}

func calculateFuel(mass int) int {
	return (mass / 3) - 2
}
